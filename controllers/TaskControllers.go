package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
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

func formatTime(timestr string) (time.Time, error) {
	if timestr != "" {
		estLocation, err := time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			return time.Time{}, nil
		}

		layout := "2006-01-02"

		t, erro := time.ParseInLocation(layout, timestr, estLocation)
		if erro != nil {
			return time.Time{}, erro
		}

		return t, nil
	}
	return time.Time{}, nil
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

	dateFormat, err := formatTime(input.Date)

	if err != nil {
		log.Fatalf("the date inputed is invalid: %s  error: %s", input.Date, err)
	}

	task := models.Task{
		Name:              input.Name,
		Erp:               input.Erp,
		Date:              dateFormat,
		Observations:      input.Observations,
		StoreID:           input.StoreID.ID,
		UserID:            input.UserID.ID,
		TaskTypeID:        input.TaskTypeID.ID,
		TimeID:            input.TimeID.ID,
		IntegrationTeamID: input.IntegrationTeamID.ID,
	}

	db := database.GetDatabase()

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

	dateFormat, err := formatTime(input.Date)
	if err != nil {
		log.Fatalf("the date inputed is invalid: %s  error: %s", input.Date, err)
	}

	updateTask := models.Task{
		Name:              input.Name,
		Erp:               input.Erp,
		Date:              dateFormat,
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
