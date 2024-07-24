package controllers

import (
	"net/http"
	"strconv"

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

func (ctrl *TaskController) EliminateTask(c *gin.Context) {
    taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    err = ctrl.service.DeleteTask(uint(taskID))
    if err != nil {
        if err.Error() == "task not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "task successfully deleted"})
}

func (ctrl *TaskController) UpdateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    err = ctrl.service.UpdateTask(&task, uint(taskID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "task successfully updated"})
}
