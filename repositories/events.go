package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type Events struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) *Events {
	return &Events{db}
}

func (repository Events) GetEvents(perPage int, page int) ([]models.Event, error) {
	var events []models.Event

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Preload("Time").Preload("IntegrationTeam").Find(&events).Error
	if err != nil {
		return []models.Event{}, err
	}

	return events, nil
}

func (repository Events) GetEventByID(id int) (models.Event, error) {
	var event models.Event

	if err := repository.db.Where("id = ?", id).Preload("Time").Preload("IntegrationTeam").First(&event).Error; err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (repository Events) CreateEvent(event models.Event) models.Event {
	repository.db.Create(&event)

	return event
}

func (repository Events) UpdateEvent(id int, updateEvent models.Event) (models.Event, error) {
	var event models.Event

	if err := repository.db.Where("id = ?", id).First(&event).Error; err != nil {
		return models.Event{}, err
	}

	repository.db.Model(&event).Updates(updateEvent)

	return updateEvent, nil
}

func (repository Events) DeleteEvent(id int) (bool, error) {
	var event models.Event

	if err := repository.db.Where("id = ?", id).First(&event).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&event)

	return true, nil
}
