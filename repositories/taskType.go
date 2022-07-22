package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type TaskType struct {
	db *gorm.DB
}

func NewTaskTypeRepository(db *gorm.DB) *TaskType {
	return &TaskType{db}
}

func (repository TaskType) GetTaskTypes(perPage int, page int) ([]models.TaskType, error) {
	var taskType []models.TaskType

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Find(&taskType).Error

	if err != nil {
		return []models.TaskType{}, err
	}

	return taskType, nil
}

func (repository TaskType) GetTaskTypeByID(id int) (models.TaskType, error) {
	var taskType models.TaskType

	if err := repository.db.Where("id = ?", id).First(&taskType).Error; err != nil {
		return models.TaskType{}, nil
	}

	return taskType, nil
}

func (repository TaskType) CreateTaskType(taskType models.TaskType) models.TaskType {
	repository.db.Create(&taskType)

	return taskType
}

func (repository TaskType) UpdateTaskType(id int, newTaskType models.TaskType) (models.TaskType, error) {
	var taskType models.TaskType

	if err := repository.db.Where("id = ?", id).First(&taskType).Error; err != nil {
		return models.TaskType{}, err
	}

	repository.db.Model(&taskType).Updates(newTaskType)

	return newTaskType, nil
}

func (repository TaskType) DeleteTaskType(id int) (bool, error) {
	var taskType models.TaskType

	if err := repository.db.Where("id = ?", id).First(&taskType).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&taskType)

	return true, nil
}
