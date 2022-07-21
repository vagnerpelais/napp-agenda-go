package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/models"
	"github.com/vagnerpelais/napp-agenda/repositories"
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 2

	repository := repositories.NewStoreRepository(db)
	store, err := repository.GetStores(perPage, page)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find all stores: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"total": strconv.Itoa(len(store)), "data": store})
}

func GetStoreByID(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewStoreRepository(db)
	store, err := repository.GetStoreByID(id)
	if err != nil {
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

	repository := repositories.NewStoreRepository(db)
	newStore := repository.CreateStore(store)

	c.JSON(http.StatusOK, gin.H{"data": newStore})
}

func UpdateStore(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
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

	repository := repositories.NewStoreRepository(db)
	newStore, err := repository.UpdateStore(id, updateStore)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newStore})
}

func DeleteStore(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param malformated!"})
	}

	repository := repositories.NewStoreRepository(db)
	store, err := repository.DeleteStore(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}
