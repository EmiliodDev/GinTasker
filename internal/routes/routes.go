package routes

import (
	"github.com/EmiliodDev/GinTasker/internal/controllers"
	"github.com/EmiliodDev/GinTasker/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    taskService := services.NewTaskService(db)
    taskController := controllers.NewTaskController(taskService)

    api := r.Group("/api/v1")
    {
        api.POST("/create", taskController.CreateNewTask)
        api.DELETE("/delete/:id", taskController.EliminateTask)
        api.PUT("/update/:id", taskController.UpdateTask)
    }

    return r
}
