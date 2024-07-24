package services

import (
	"errors"

	"github.com/EmiliodDev/GinTasker/internal/models"
	"gorm.io/gorm"
)

type TaskService struct {
    db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
    return &TaskService{db: db}
}

func (s *TaskService) CreateTask(task *models.Task) error {
    return s.db.Create(task).Error
}

func (s *TaskService) DeleteTask(taskID uint) error {
    r := s.db.Delete(&models.Task{}, taskID)
    
    if r.Error != nil {
        return r.Error
    }

    if r.RowsAffected == 0 {
        return errors.New("task not found")
    }

    return nil
}

func (s *TaskService) UpdateTask(task *models.Task, taskID uint) error {
    return s.db.Save(&models.Task{ID: taskID, Name: task.Name, Description: task.Description, Completed: task.Completed}).Error
}
