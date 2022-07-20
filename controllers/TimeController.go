package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
)

type CreateTimeInput struct {
	Hour string `json:"hour" binding:"required"`
}

type UpdateTimeInput struct {
	Hour string `json:"hour"`
}

func GetTimes(c *gin.Context) {
	db := database.GetDatabase()

	var times []models.Time
	err := db.Find(&times).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all times: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(times)), "data": times})
}

func GetTimeByID(c *gin.Context) {
	var time models.Time

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&time).Error; err != nil {
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

	db.Create(&time)

	c.JSON(http.StatusOK, gin.H{"data": time})
}

func UpdateTime(c *gin.Context) {
	var time models.Time
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTime := models.Time{
		Hour: input.Hour,
	}

	db.Model(&time).Updates(updateTime)

	c.JSON(http.StatusOK, gin.H{"data": time})
}
