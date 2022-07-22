package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewUserRepository(db)
	users, err := repository.GetUsers(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all users: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(users)), "data": users})
}

func GetUserByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.GetUserByID(id)
	if err != nil {
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

	repository := repositories.NewUserRepository(db)
	newUser := repository.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
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

	repository := repositories.NewUserRepository(db)
	newUser, err := repository.UpdateUser(id, updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func DeleteUser(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewUserRepository(db)
	user, err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
