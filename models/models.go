package models

import (
	"time"

	"gorm.io/gorm"
)

type IntegrationTeam struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Color        string         `json:"color"`
	LimitPerHour uint           `json:"limit_per_hour"`
	CreatedAt    time.Time      `json:"created"`
	UpdatedAt    time.Time      `json:"updated"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Email      string         `json:"email"`
	IsActive   bool           `json:"is_active" gorm:"default:true"`
	IsAdmin    bool           `json:"is_admin" gorm:"default:false"`
	IsTechLead bool           `json:"is_techlead" gorm:"default:false"`
	CreatedAt  time.Time      `json:"created"`
	UpdatedAt  time.Time      `json:"updated"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Store struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Cnpj      string         `json:"cnpj"`
	LegalName string         `json:"legal_name"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Time struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Hour      string         `json:"hour"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type TaskType struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Task struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Name              string         `json:"name"`
	Erp               string         `json:"erp"`
	Date              time.Time      `json:"date"`
	Observations      string         `json:"observations"`
	StoreID           uint           `json:"store"`
	UserID            uint           `json:"user"`
	TaskTypeID        uint           `json:"task_type"`
	TimeID            uint           `json:"time"`
	IntegrationTeamID uint           `json:"integration_team"`
	CreatedAt         time.Time      `json:"created"`
	UpdatedAt         time.Time      `json:"updated"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted"`
}
