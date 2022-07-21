package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
)

type CreateUserInput struct {
	Email      string `json:"email" binding:"required"`
	IsActive   bool   `json:"is_active"`
	IsAdmin    bool   `json:"is_admin"`
	IsTechLead bool   `json:"is_tech_lead"`
}

type UpdateUserInput struct {
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
	IsAdmin    bool   `json:"is_admin"`
	IsTechLead bool   `json:"is_tech_lead"`
}

func GetUsers(c *gin.Context) {
	db := database.GetDatabase()

	var users []models.User

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	err := db.Limit(perPage).Offset((page - 1) * perPage).Find(&users).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all users: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(users)), "data": users})
}

func GetUserByID(c *gin.Context) {
	var user models.User

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:      input.Email,
		IsActive:   input.IsActive,
		IsAdmin:    input.IsAdmin,
		IsTechLead: input.IsTechLead,
	}

	db := database.GetDatabase()

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser := models.User{
		Email:      input.Email,
		IsActive:   input.IsActive,
		IsAdmin:    input.IsAdmin,
		IsTechLead: input.IsTechLead,
	}

	db.Model(&user).Updates(updateUser)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
