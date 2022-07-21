package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/services"
	"gorm.io/gorm"
)

type UpdateTaskInput struct {
	Name              string `json:"name"`
	Erp               string `json:"erp"`
	Date              string `json:"date,omitempty"`
	Observations      string `json:"observations"`
	StoreID           uint   `json:"store"`
	UserID            uint   `json:"user"`
	TaskTypeID        uint   `json:"task_type"`
	TimeID            uint   `json:"time"`
	IntegrationTeamID uint   `json:"integration_team"`
}

type ResponseIntegrationTeam struct {
	gorm.Model
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	LimitPerHour uint   `json:"limit_per_hour"`
}

type ResponseStore struct {
	gorm.Model
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Cnpj      string `json:"cnpj"`
	LegalName string `json:"legal_name"`
}

type ResponseTaskType struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ResponseTime struct {
	gorm.Model
	ID   uint   `json:"id"`
	Hour string `json:"hour"`
}

type ResponseUser struct {
	gorm.Model
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
	IsAdmin    bool   `json:"is_admin"`
	IsTechLead bool   `json:"is_tech_lead"`
}

type CreateTaskInput struct {
	Name              string                  `json:"name" binding:"required"`
	Erp               string                  `json:"erp" binding:"required"`
	Date              string                  `json:"date" binding:"required"`
	Observations      string                  `json:"observations" binding:"required"`
	StoreID           ResponseStore           `json:"store" binding:"required"`
	UserID            ResponseUser            `json:"user" binding:"required"`
	TaskTypeID        ResponseTaskType        `json:"task_type" binding:"required"`
	TimeID            ResponseTime            `json:"time" binding:"required"`
	IntegrationTeamID ResponseIntegrationTeam `json:"integration_team" binding:"required"`
}

func getTasksCount(teamid int64, timeid int64, date string) uint {
	var count uint

	db := database.GetDatabase()

	db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", teamid, timeid, date).Scan(&count)

	return count
}

func getIntegrationTeamLimit(id int64) uint {
	var limit uint

	db := database.GetDatabase()

	db.Raw("SELECT limit_per_hour from integration_team where id = ?", id).Scan(&limit)

	return limit
}

func GetTasks(c *gin.Context) {
	db := database.GetDatabase()

	var task []models.Task

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	err := db.Limit(perPage).Offset((page - 1) * perPage).Preload("Store").Preload("User").
		Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all task: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(task)), "data": task})
}

func GetTaskByID(c *gin.Context) {
	var task models.Task
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).Preload("Store").Preload("User").
		Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateString, dateFormat, err := services.FormatTime(input.Date)
	if err != nil {
		log.Fatalf("the date inputed is invalid: %s  error: %s", input.Date, err)
	}

	weekend := services.IsWeekend(dateFormat)
	if weekend {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot schedule in weekends!"})
		return
	}

	holiday := services.CheckHoliday(dateFormat)
	if holiday {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot schedule in holidays!"})
		return
	}

	task := models.Task{
		Name:              input.Name,
		Erp:               input.Erp,
		Date:              dateString,
		Observations:      input.Observations,
		StoreID:           input.StoreID.ID,
		UserID:            input.UserID.ID,
		TaskTypeID:        input.TaskTypeID.ID,
		TimeID:            input.TimeID.ID,
		IntegrationTeamID: input.IntegrationTeamID.ID,
	}

	db := database.GetDatabase()

	// var event []models.Events
	// var count int64
	// db.Find(&event, "ilike date = ?", dateFormat).Count(&count)

	// log.Printf("%d", count)

	counter := getTasksCount(int64(input.IntegrationTeamID.ID), int64(input.TimeID.ID), dateString)
	limit := getIntegrationTeamLimit(int64(input.IntegrationTeamID.ID))

	if counter >= limit {
		c.JSON(http.StatusForbidden, gin.H{"error": "Limit of schedules per hour exceeded!"})
		return
	}

	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateString, dateFormat, err := services.FormatTime(input.Date)
	if err != nil {
		log.Fatalf("the date inputed is invalid: %s  error: %s", input.Date, err)
	}

	if (time.Time{}) != dateFormat {
		weekend := services.IsWeekend(dateFormat)
		if weekend {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot schedule in weekends!"})
			return
		}

		holiday := services.CheckHoliday(dateFormat)
		if holiday {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot schedule in holidays!"})
			return
		}
	}

	updateTask := models.Task{
		Name:              input.Name,
		Erp:               input.Erp,
		Date:              dateString,
		Observations:      input.Observations,
		StoreID:           input.StoreID,
		UserID:            input.UserID,
		TaskTypeID:        input.TaskTypeID,
		TimeID:            input.TimeID,
		IntegrationTeamID: input.IntegrationTeamID,
	}

	db.Model(&task).Updates(updateTask)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}