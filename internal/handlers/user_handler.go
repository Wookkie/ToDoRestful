package handlers

import (
	"net/http"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.service.GetAllUsers())
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	ctx.JSON(http.StatusFound, user)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать пользователя"})
		return
	}
	create := h.service.CreateUser(user)
	ctx.JSON(http.StatusOK, create)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User

	updated, err := h.service.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}
	ctx.JSON(http.StatusOK, updated)

}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}
	ctx.Status(http.StatusNoContent)

}
