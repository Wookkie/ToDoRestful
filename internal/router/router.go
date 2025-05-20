package router

import (
	"github.com/Wookkie/ToDoRestful/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
	tasks := r.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetAllTasks)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.POST("", taskHandler.CreateTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}
}

func UserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.POST("", userHandler.CreateUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
