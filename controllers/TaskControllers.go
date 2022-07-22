package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
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

func GetTasks(c *gin.Context) {
	db := database.GetDatabase()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewTaskRepository(db)
	tasks, err := repository.GetTasks(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all task: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(tasks)), "data": tasks})
}

func GetTaskByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTaskRepository(db)
	task, err := repository.GetTaskByID(id)
	if err != nil {
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
	repository := repositories.NewTaskRepository(db)

	counter := repository.GetTasksCount(int64(input.IntegrationTeamID.ID), int64(input.TimeID.ID), dateString)
	limit := repository.GetIntegrationTeamLimit(0, int64(input.IntegrationTeamID.ID))

	if counter >= limit {
		c.JSON(http.StatusForbidden, gin.H{"error": "Limit of schedules per hour exceeded!"})
		return
	}

	newTask := repository.CreateTask(task)

	c.JSON(http.StatusOK, gin.H{"data": newTask})
}

func UpdateTask(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
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

	repository := repositories.NewTaskRepository(db)

	counter := repository.GetTasksCountUpdate(id, int(input.IntegrationTeamID), int(input.TimeID), dateString)
	limit := repository.GetIntegrationTeamLimit(int64(id), int64(input.IntegrationTeamID))

	if counter >= limit {
		c.JSON(http.StatusForbidden, gin.H{"error": "Limit of schedules per hour exceeded!"})
		return
	}

	newTask, err := repository.UpdateTask(id, updateTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newTask})
}

func DeleteTask(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTaskRepository(db)
	task, err := repository.TaskDelete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
