package router

import (
	"github.com/Wookkie/ToDoRestful/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
	tasks := r.Group("/tasks")
	{
		tasks.GET("/task", taskHandler.GetAllTasks)
		tasks.GET("/task/:id", taskHandler.GetTaskByID)
		tasks.POST("/task", taskHandler.CreateTask)
		tasks.PUT("/task/:id", taskHandler.UpdateTask)
		tasks.DELETE("/task/:id", taskHandler.DeleteTask)
	}
}
