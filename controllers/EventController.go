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
)

type CreateEventInput struct {
	Name              string `json:"name" binding:"required"`
	Date              string `json:"date" binding:"required"`
	Weigth            uint   `json:"weigth" binding:"required"`
	TimeID            uint   `json:"time_id" binding:"required"`
	IntegrationTeamID uint   `json:"integration_team_id" binding:"required"`
}

type UpdateEventInput struct {
	Name              string `json:"name"`
	Date              string `json:"date"`
	Weigth            uint   `json:"weigth"`
	TimeID            uint   `json:"time_id"`
	IntegrationTeamID uint   `json:"integration_team_id"`
}

func GetEvents(c *gin.Context) {
	db := database.GetDatabase()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewEventsRepository(db)
	events, err := repository.GetEvents(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all events: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(events)), "data": events})
}

func GetEventByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewEventsRepository(db)
	event, err := repository.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

func CreateEvent(c *gin.Context) {
	var input CreateEventInput

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

	event := models.Event{
		Name:              input.Name,
		Date:              dateString,
		Weigth:            input.Weigth,
		TimeID:            input.TimeID,
		IntegrationTeamID: input.IntegrationTeamID,
	}

	db := database.GetDatabase()

	repository := repositories.NewEventsRepository(db)
	newEvent := repository.CreateEvent(event)

	c.JSON(http.StatusOK, gin.H{"data": newEvent})
}

func UpdateEvent(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	var input UpdateEventInput
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

	updateEvent := models.Event{
		Name:              input.Name,
		Date:              dateString,
		Weigth:            input.Weigth,
		TimeID:            input.TimeID,
		IntegrationTeamID: input.IntegrationTeamID,
	}

	repository := repositories.NewEventsRepository(db)
	newEvent, err := repository.UpdateEvent(id, updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newEvent})
}

func DeleteEvent(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewEventsRepository(db)
	event, err := repository.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}
