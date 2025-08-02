package handlers

import (
	"net/http"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

// func getUID(ctx *gin.Context) string {
// 	uid, _ := ctx.Get("uid")
// 	return uid.(string)
// }

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasks(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	tasks := h.service.GetTasksByUserID(uid)
	ctx.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := ctx.GetString("uid")

	task, err := h.service.GetTaskByID(id, uid)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var task model.Task
	uid := ctx.GetString("uid")

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать задачу"})
		return
	}

	task.UserID = uid
	created := h.service.CreateTask(task)
	ctx.JSON(http.StatusCreated, created)
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task model.Task
	uid := ctx.GetString("uid")

	updated, err := h.service.UpdateTask(id, task, uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.JSON(http.StatusOK, updated)

}

func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := ctx.GetString("uid")

	err := h.service.DeleteTask(id, uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
