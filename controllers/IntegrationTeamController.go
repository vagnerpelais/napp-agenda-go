package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
)

type CreateIntegrationTeamInput struct {
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	LimitPerHour uint   `json:"limit_per_hour" binding:"required"`
}

type UpdateIntegrationTeamInput struct {
	Name         string `json:"name"`
	Color        string `json:"color"`
	LimitPerHour uint   `json:"limit_per_hour"`
}

func GetIntegrationTeams(c *gin.Context) {
	db := database.GetDatabase()

	var integrationTeam []models.IntegrationTeam
	err := db.Find(&integrationTeam).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all teams: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(integrationTeam)), "data": integrationTeam})
}

func GetIntegrationTeamByID(c *gin.Context) {
	var integrationTeam models.IntegrationTeam

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&integrationTeam).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": integrationTeam})
}

func CreateIntegrationTeam(c *gin.Context) {
	var input CreateIntegrationTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	integrationTeam := models.IntegrationTeam{
		Name:         input.Name,
		Color:        input.Color,
		LimitPerHour: input.LimitPerHour,
	}

	db := database.GetDatabase()

	db.Create(&integrationTeam)

	c.JSON(http.StatusOK, gin.H{"data": integrationTeam})
}

func UpdateIntegrationTeam(c *gin.Context) {
	var integrationTeam models.IntegrationTeam
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&integrationTeam).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateIntegrationTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateStore := models.IntegrationTeam{
		Name:         input.Name,
		Color:        input.Color,
		LimitPerHour: input.LimitPerHour,
	}

	db.Model(&integrationTeam).Updates(updateStore)

	c.JSON(http.StatusOK, gin.H{"data": integrationTeam})
}
