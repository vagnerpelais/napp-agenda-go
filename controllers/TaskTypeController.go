package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
)

type CreateTaskTypeInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskTypeInput struct {
	Name string `json:"name"`
}

func GetTaskTypes(c *gin.Context) {
	db := database.GetDatabase()

	var taskTypes []models.TaskType

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	err := db.Limit(perPage).Offset((page - 1) * perPage).Find(&taskTypes).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all task types: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(taskTypes)), "data": taskTypes})
}

func GetTaskTypeByID(c *gin.Context) {
	var taskType models.TaskType

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&taskType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": taskType})
}

func CreateTaskType(c *gin.Context) {
	var input CreateTaskTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskType := models.TaskType{
		Name: input.Name,
	}

	db := database.GetDatabase()

	db.Create(&taskType)

	c.JSON(http.StatusOK, gin.H{"data": taskType})
}

func UpdateTaskType(c *gin.Context) {
	var taskType models.TaskType
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&taskType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTaskTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTaskType := models.TaskType{
		Name: input.Name,
	}

	db.Model(&taskType).Updates(updateTaskType)

	c.JSON(http.StatusOK, gin.H{"data": taskType})
}

func DeleteTaskType(c *gin.Context) {
	var taskType models.TaskType

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&taskType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&taskType)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
