package migrations

import (
	"gorm.io/gorm"

	"github.com/vagnerpelais/napp-agenda/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.IntegrationTeam{})
	db.AutoMigrate(models.Store{})
	db.AutoMigrate(models.Task{})
	db.AutoMigrate(models.TaskType{})
	db.AutoMigrate(models.Time{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Event{})
}
