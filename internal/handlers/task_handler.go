package handlers

import (
	"net/http"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct { //почему здесь есть указатель, а в TaskService в структуре его нет
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.service.GetAllTasks()) //почему в данном случае не нужен return
}

func (h *TaskHandler) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.JSON(http.StatusOK, task) //почему в данном случае не нужен return
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var task model.Task //почему model.Task без []
	created := h.service.CreateTask(task)
	ctx.JSON(http.StatusCreated, created)
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task model.Task

	updated, err := h.service.UpdateTask(id, task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.JSON(http.StatusOK, updated)

}

func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
