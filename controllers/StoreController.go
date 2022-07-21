package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
)

type CreateStoreInput struct {
	Name      string `json:"name" binding:"required"`
	Cnpj      string `json:"cnpj" binding:"required"`
	LegalName string `json:"legal_name" binding:"required"`
}

type UpdateStoreInput struct {
	Name      string `json:"name"`
	Cnpj      string `json:"cnpj"`
	LegalName string `json:"legal_name"`
}

func GetStores(c *gin.Context) {
	db := database.GetDatabase()

	var stores []models.Store

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	err := db.Limit(perPage).Offset((page - 1) * perPage).Find(&stores).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all stores: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(stores)), "data": stores})
}

func GetStoreByID(c *gin.Context) {
	var store models.Store

	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&store).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

func CreateStore(c *gin.Context) {
	var input CreateStoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := models.Store{
		Name:      input.Name,
		Cnpj:      input.Cnpj,
		LegalName: input.LegalName,
	}

	db := database.GetDatabase()

	db.Create(&store)

	c.JSON(http.StatusOK, gin.H{"data": store})
}

func UpdateStore(c *gin.Context) {
	var store models.Store
	db := database.GetDatabase()

	if err := db.Where("id = ?", c.Param("id")).First(&store).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateStoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateStore := models.Store{
		Name:      input.Name,
		Cnpj:      input.Cnpj,
		LegalName: input.LegalName,
	}

	db.Model(&store).Updates(updateStore)

	c.JSON(http.StatusOK, gin.H{"data": store})
}
