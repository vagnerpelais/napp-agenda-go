package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
)

type CreateTimeInput struct {
	Hour string `json:"hour" binding:"required"`
}

type UpdateTimeInput struct {
	Hour string `json:"hour"`
}

func GetTimes(c *gin.Context) {
	db := database.GetDatabase()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewTimeRepository(db)
	times, err := repository.GetTimes(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all times: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(times)), "data": times})
}

func GetTimeByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTimeRepository(db)
	time, err := repository.GetTimeByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}

func CreateTime(c *gin.Context) {
	var input CreateTimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	time := models.Time{
		Hour: input.Hour,
	}

	db := database.GetDatabase()

	repository := repositories.NewTimeRepository(db)
	newTime := repository.CreateTime(time)

	c.JSON(http.StatusOK, gin.H{"data": newTime})
}

func UpdateTime(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	var input UpdateTimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTime := models.Time{
		Hour: input.Hour,
	}

	repository := repositories.NewTimeRepository(db)
	newTime, err := repository.UpdateTime(id, updateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newTime})
}

func DeleteTime(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewTimeRepository(db)
	time, err := repository.DeleteTime(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}
