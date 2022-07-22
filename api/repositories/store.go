package repositories

import (
	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *Store {
	return &Store{db}
}

func (repository Store) GetStores(perPage int, page int) ([]models.Store, error) {
	var store []models.Store

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Find(&store).Error

	if err != nil {
		return []models.Store{}, err
	}

	return store, nil
}

func (repository Store) GetStoreByID(id int) (models.Store, error) {
	var store models.Store

	if err := repository.db.Where("id = ?", id).First(&store).Error; err != nil {
		return models.Store{}, nil
	}

	return store, nil
}

func (repository Store) CreateStore(store models.Store) models.Store {
	repository.db.Create(&store)

	return store
}

func (repository Store) UpdateStore(id int, newStore models.Store) (models.Store, error) {
	var store models.Store

	if err := repository.db.Where("id = ?", id).First(&store).Error; err != nil {
		return models.Store{}, err
	}

	repository.db.Model(&store).Updates(newStore)

	return newStore, nil
}

func (repository Store) DeleteStore(id int) (bool, error) {
	var store models.Store

	if err := repository.db.Where("id = ?", id).First(&store).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&store)

	return true, nil
}
