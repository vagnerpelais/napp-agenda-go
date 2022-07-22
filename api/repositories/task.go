package repositories

import (
	"log"

	"github.com/vagnerpelais/napp-agenda/models"
	"gorm.io/gorm"
)

type Task struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *Task {
	return &Task{db}
}

func (repository Task) GetTasksCountUpdate(id int, teamid int, timeid int, date string) uint {
	if teamid != 0 && timeid != 0 && date != "" {
		var count uint

		repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", teamid, timeid, date).Scan(&count)

		return count
	}

	if teamid != 0 && timeid == 0 && date == "" {
		var count uint
		var task models.Task

		if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
			Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
			log.Fatal("Error", err)
		}

		repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", teamid, task.TimeID, task.Date).Scan(&count)

		return count
	}

	if teamid == 0 && timeid != 0 && date == "" {
		var count uint
		var task models.Task

		if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
			Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
			log.Fatal("Error", err)
		}

		repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", task.IntegrationTeamID, timeid, task.Date).Scan(&count)

		return count
	}

	if teamid == 0 && timeid == 0 && date != "" {
		var count uint
		var task models.Task

		if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
			Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
			log.Fatal("Error", err)
		}

		repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", task.IntegrationTeamID, task.TimeID, date).Scan(&count)

		return count
	}

	if teamid == 0 && timeid == 0 && date == "" {
		var count uint
		var task models.Task

		if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
			Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
			log.Fatal("Error", err)
		}

		repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", task.IntegrationTeamID, task.TimeID, task.Date).Scan(&count)

		return count
	}

	return 0
}

func (repository Task) GetTasksCount(teamid int64, timeid int64, date string) uint {
	var count uint

	repository.db.Raw("SELECT COUNT(*) FROM task WHERE integration_team_id = ? and time_id = ? and date = ?", teamid, timeid, date).Scan(&count)

	return count
}

func (repository Task) GetIntegrationTeamLimit(id int64, teamid int64) uint {
	if teamid == 0 {
		var limit uint
		var task models.Task

		if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
			Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
			log.Fatal("Error", err)
		}

		repository.db.Raw("SELECT limit_per_hour from integration_team where id = ?", task.IntegrationTeamID).Scan(&limit)

		return limit
	}

	var limit uint

	repository.db.Raw("SELECT limit_per_hour from integration_team where id = ?", teamid).Scan(&limit)

	return limit
}

func (repository Task) GetTasks(perPage int, page int) ([]models.Task, error) {
	var task []models.Task

	err := repository.db.Limit(perPage).Offset((page - 1) * perPage).Preload("Store").Preload("User").
		Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error
	if err != nil {
		return []models.Task{}, err
	}

	return task, nil
}

func (repository Task) GetTaskByID(id int) (models.Task, error) {
	var task models.Task

	if err := repository.db.Where("id = ?", id).Preload("Store").Preload("User").
		Preload("TaskType").Preload("Time").Preload("IntegrationTeam").Find(&task).Error; err != nil {
		return models.Task{}, nil
	}

	return task, nil
}

func (repository Task) CreateTask(task models.Task) models.Task {
	repository.db.Create(&task)

	return task
}

func (repository Task) UpdateTask(id int, newTask models.Task) (models.Task, error) {
	var task models.Task

	if err := repository.db.Where("id = ?", id).First(&task).Error; err != nil {
		return models.Task{}, err
	}

	repository.db.Model(&task).Updates(newTask)

	return newTask, nil
}

func (repository Task) TaskDelete(id int) (bool, error) {
	var task models.Task

	if err := repository.db.Where("id = ?", id).First(&task).Error; err != nil {
		return false, err
	}

	repository.db.Delete(&task)

	return true, nil
}
