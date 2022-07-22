package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
)

type CreateTaskTypeInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskTypeInput struct {
	Name string `json:"name"`
}

func GetTaskTypes(c *gin.Context) {
	db := database.GetDatabase()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewTaskTypeRepository(db)
	taskTypes, err := repository.GetTaskTypes(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all task types: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(taskTypes)), "data": taskTypes})
}

func GetTaskTypeByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTaskTypeRepository(db)
	taskType, err := repository.GetTaskTypeByID(id)
	if err != nil {
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

	repository := repositories.NewTaskTypeRepository(db)
	newTaskType := repository.CreateTaskType(taskType)

	c.JSON(http.StatusOK, gin.H{"data": newTaskType})
}

func UpdateTaskType(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	var input UpdateTaskTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTaskType := models.TaskType{
		Name: input.Name,
	}

	repository := repositories.NewTaskTypeRepository(db)
	newTaskType, err := repository.UpdateTaskType(id, updateTaskType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newTaskType})
}

func DeleteTaskType(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTaskTypeRepository(db)
	taskType, err := repository.DeleteTaskType(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": taskType})
}
