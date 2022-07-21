package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type IntegrationTeam struct {
	db *gorm.DB
}

func NewIntegrationTeamRepository(db *gorm.DB) *IntegrationTeam {
	return &IntegrationTeam{db}
}

func (repository IntegrationTeam) GetIntegrationTeams(perPage int, page int) ([]models.IntegrationTeam, error) {
	var integrationTeam []models.IntegrationTeam

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Find(&integrationTeam).Error

	if err != nil {
		return []models.IntegrationTeam{}, err
	}

	return integrationTeam, nil
}

func (repository IntegrationTeam) GetIntegrationTeamByID(id int) (models.IntegrationTeam, error) {
	var integrationTeam models.IntegrationTeam

	if err := repository.db.Where("id = ?", id).First(&integrationTeam).Error; err != nil {
		return models.IntegrationTeam{}, err
	}

	return integrationTeam, nil
}

func (repository IntegrationTeam) CreateIntegrationTeam(integrationTeam models.IntegrationTeam) models.IntegrationTeam {
	repository.db.Create(&integrationTeam)

	return integrationTeam
}

func (repository IntegrationTeam) UpdateIntegrationTeam(id int, updateIntegrationTeam models.IntegrationTeam) (models.IntegrationTeam, error) {
	var integrationTeam models.IntegrationTeam

	if err := repository.db.Where("id = ?", id).First(&integrationTeam).Error; err != nil {
		return models.IntegrationTeam{}, err
	}

	repository.db.Model(&integrationTeam).Updates(updateIntegrationTeam)

	return updateIntegrationTeam, nil
}

func (repository IntegrationTeam) DeleteIntegrationTeam(id int) (bool, error) {
	var integrationTeam models.IntegrationTeam

	if err := repository.db.Where("id = ?", id).First(&integrationTeam).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&integrationTeam)

	return true, nil
}
