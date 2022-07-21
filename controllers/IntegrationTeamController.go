package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewIntegrationTeamRepository(db)
	integrationTeam, err := repository.GetIntegrationTeams(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all teams: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(integrationTeam)), "data": integrationTeam})
}

func GetIntegrationTeamByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewIntegrationTeamRepository(db)
	integrationTeam, err := repository.GetIntegrationTeamByID(id)
	if err != nil {
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

	repository := repositories.NewIntegrationTeamRepository(db)
	newIntegrationTeam := repository.CreateIntegrationTeam(integrationTeam)

	c.JSON(http.StatusOK, gin.H{"data": newIntegrationTeam})
}

func UpdateIntegrationTeam(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	var input UpdateIntegrationTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateIntegrationTeam := models.IntegrationTeam{
		Name:         input.Name,
		Color:        input.Color,
		LimitPerHour: input.LimitPerHour,
	}

	repository := repositories.NewIntegrationTeamRepository(db)
	newIntegrationTeam, err := repository.UpdateIntegrationTeam(id, updateIntegrationTeam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newIntegrationTeam})
}

func DeleteIntegrationTeam(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewIntegrationTeamRepository(db)
	integrationTeam, err := repository.DeleteIntegrationTeam(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": integrationTeam})
}
