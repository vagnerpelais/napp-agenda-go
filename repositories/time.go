package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type Time struct {
	db *gorm.DB
}

func NewTimeRepository(db *gorm.DB) *Time {
	return &Time{db}
}

func (repository Time) GetTimes(perPage int, page int) ([]models.Time, error) {
	var time []models.Time

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Find(&time).Error

	if err != nil {
		return []models.Time{}, err
	}

	return time, nil
}

func (repository Time) GetTimeByID(id int) (models.Time, error) {
	var time models.Time

	if err := repository.db.Where("id = ?", id).First(&time).Error; err != nil {
		return models.Time{}, nil
	}

	return time, nil
}

func (repository Time) CreateTime(time models.Time) models.Time {
	repository.db.Create(&time)

	return time
}

func (repository Time) UpdateTime(id int, newTime models.Time) (models.Time, error) {
	var time models.Time

	if err := repository.db.Where("id = ?", id).First(&time).Error; err != nil {
		return models.Time{}, err
	}

	repository.db.Model(&time).Updates(newTime)

	return newTime, nil
}

func (repository Time) DeleteTime(id int) (bool, error) {
	var time models.Time

	if err := repository.db.Where("id = ?", id).First(&time).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&time)

	return true, nil
}
