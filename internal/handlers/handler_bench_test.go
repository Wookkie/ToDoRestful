package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Wookkie/ToDoRestful/internal/handlers"
	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

func BenchmarkTaskHandler_GetAllTasks(b *testing.B) {
	repo := service.NewFakeTaskRepo()
	svc := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	for i := 0; i < 100; i++ {
		t := model.Task{ID: strconv.Itoa(i), Title: "Task", UserID: "u1"}
		repo.CreateTask(t)
	}

	gin.SetMode(gin.TestMode)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set("uid", "u1")
		handler.GetAllTasks(ctx)
	}
}

func BenchmarkTaskHandler_CreateTask(b *testing.B) {
	repo := service.NewFakeTaskRepo()
	svc := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	gin.SetMode(gin.TestMode)

	taskJSON := `{"Title":"BenchmarkTask"}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(taskJSON))
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Set("uid", "u1")

		handler.CreateTask(ctx)
	}
}

func BenchmarkUserHandler_GetAllUsers(b *testing.B) {
	repo := service.NewFakeUserRepo()
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc)

	for i := 0; i < 100; i++ {
		u := model.User{ID: strconv.Itoa(i), Name: "User"}
		repo.CreateUser(u)
	}

	gin.SetMode(gin.TestMode)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		handler.GetAllUsers(ctx)
	}
}

func BenchmarkUserHandler_CreateUser(b *testing.B) {
	repo := service.NewFakeUserRepo()
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc)

	gin.SetMode(gin.TestMode)

	userJSON := `{"Name":"BenchmarkUser","Email":"b@bench.com"}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		ctx.Request.Header.Set("Content-Type", "application/json")

		handler.CreateUser(ctx)
	}
}
