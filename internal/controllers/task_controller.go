package controllers

import (
	"net/http"

	"github.com/EmiliodDev/GinTasker/internal/models"
	"github.com/EmiliodDev/GinTasker/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
    service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
    return &TaskController{service: service}
}

func (ctrl *TaskController) CreateNewTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.service.CreateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating task"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "task successfully created"})
}
