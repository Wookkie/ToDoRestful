package main

import (
	"github.com/Wookkie/ToDoRestful/internal/handlers"
	"github.com/Wookkie/ToDoRestful/internal/router"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	taskService := service.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()
	router.TaskRoutes(r, taskHandler)

	r.Run(":8080")
}
