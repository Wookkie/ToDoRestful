package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Wookkie/ToDoRestful/internal"
	"github.com/Wookkie/ToDoRestful/internal/handlers"
	"github.com/Wookkie/ToDoRestful/internal/repository"
	"github.com/Wookkie/ToDoRestful/internal/router"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	repository.UserRepository
	repository.TaskRepository
	Close() error
}

type API struct {
	cfg         *internal.Config
	httpServe   *http.Server
	repo        Repository
	taskService *service.TaskService
	userService *service.UserService
}

func New(cfg *internal.Config, repo Repository) *API {
	log.Println("Configuring API server")

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	httpServe := http.Server{
		Addr: addr,
	}

	api := &API{
		cfg:         cfg,
		repo:        repo,
		httpServe:   &httpServe,
		taskService: service.NewTaskService(repo),
		userService: service.NewUserService(repo),
	}

	api.configureRoutes()

	return api
}

func (api *API) Run() error {
	log.Printf("API started on %s\n", api.httpServe.Addr)
	return api.httpServe.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	log.Println("Shutting down API server")
	return api.httpServe.Shutdown(ctx)
}

func (api *API) configureRoutes() {
	log.Println("Configuring routes")
	routerEngine := gin.Default()

	taskHandler := handlers.NewTaskHandler(api.taskService)
	userHandler := handlers.NewUserHandler(api.userService)

	router.TaskRoutes(routerEngine, taskHandler)
	router.UserRoutes(routerEngine, userHandler)

	api.httpServe.Handler = routerEngine
}
