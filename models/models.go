package models

import (
	"time"
)

type IntegrationTeam struct {
	ID           uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name         string    `json:"name,omitempty" gorm:"unique"`
	Color        string    `json:"color,omitempty"`
	LimitPerHour uint      `json:"limit_per_hour,omitempty" gorm:"default:0"`
	CreatedAt    time.Time `json:"created,omitempty"`
	UpdatedAt    time.Time `json:"updated,omitempty"`
}

type Event struct {
	ID                uint            `json:"id,omitempty" gorm:"primaryKey"`
	Name              string          `json:"name,omitempty"`
	Date              string          `json:"date,omitempty"`
	Weigth            uint            `json:"weight,omitempty"`
	TimeID            uint            `json:"time_id,omitempty"`
	Time              Time            `json:"time,omitempty" gorm:"foreignKey:TimeID;references:ID"`
	IntegrationTeamID uint            `json:"integration_team_id,omitempty"`
	IntegrationTeam   IntegrationTeam `json:"integration_team,omitempty" gorm:"foreignKey:IntegrationTeamID;references:ID"`
	CreatedAt         time.Time       `json:"created,omitempty"`
	UpdatedAt         time.Time       `json:"updated,omitempty"`
}

type User struct {
	ID         uint      `json:"id,omitempty" gorm:"primaryKey"`
	Email      string    `json:"email,omitempty" gorm:"unique"`
	IsActive   bool      `json:"is_active,omitempty" gorm:"default:true"`
	IsAdmin    bool      `json:"is_admin,omitempty" gorm:"default:false"`
	IsTechLead bool      `json:"is_techlead,omitempty" gorm:"default:false"`
	CreatedAt  time.Time `json:"created,omitempty"`
	UpdatedAt  time.Time `json:"updated,omitempty"`
}

type Store struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty"`
	Cnpj      string    `json:"cnpj,omitempty" gorm:"unique"`
	LegalName string    `json:"legal_name,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}

type Time struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Hour      string    `json:"hour,omitempty" gorm:"unique"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}

type TaskType struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty" gorm:"unique"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}

type Task struct {
	ID                uint            `json:"id,omitempty" gorm:"primaryKey"`
	Name              string          `json:"name,omitempty"`
	Erp               string          `json:"erp,omitempty"`
	Date              string          `json:"date,omitempty"`
	Observations      string          `json:"observations,omitempty"`
	StoreID           uint            `json:"store_id,omitempty"`
	Store             Store           `json:"store,omitempty" gorm:"foreignKey:StoreID;references:ID"`
	UserID            uint            `json:"user_id,omitempty"`
	User              User            `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	TaskTypeID        uint            `json:"task_type_id,omitempty"`
	TaskType          TaskType        `json:"task_type,omitempty" gorm:"foreignKey:TaskTypeID;references:ID"`
	TimeID            uint            `json:"time_id,omitempty"`
	Time              Time            `json:"time,omitempty" gorm:"foreignKey:TimeID;references:ID"`
	IntegrationTeamID uint            `json:"integration_team_id,omitempty"`
	IntegrationTeam   IntegrationTeam `json:"integration_team,omitempty" gorm:"foreignKey:IntegrationTeamID;references:ID"`
	CreatedAt         time.Time       `json:"created,omitempty"`
	UpdatedAt         time.Time       `json:"updated,omitempty"`
}
