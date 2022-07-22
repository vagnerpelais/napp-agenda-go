package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{db}
}

func (repository User) GetUsers(perPage int, page int) ([]models.User, error) {
	var user []models.User

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Find(&user).Error

	if err != nil {
		return []models.User{}, err
	}

	return user, nil
}

func (repository User) GetUserByID(id int) (models.User, error) {
	var user models.User

	if err := repository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, nil
	}

	return user, nil
}

func (repository User) CreateUser(user models.User) models.User {
	repository.db.Create(&user)

	return user
}

func (repository User) UpdateUser(id int, newUser models.User) (models.User, error) {
	var user models.User

	if err := repository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	repository.db.Model(&user).Updates(newUser)

	return newUser, nil
}

func (repository User) DeleteUser(id int) (bool, error) {
	var user models.User

	if err := repository.db.Where("id = ?", id).First(&user).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&user)

	return true, nil
}
