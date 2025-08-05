package repository

import "github.com/Wookkie/ToDoRestful/internal/model"

type UserRepository interface {
	GetAllUsers() []model.User
	GetUserByID(id string) (*model.User, error)
	CreateUser(user model.User) model.User
	UpdateUser(id string, user model.User) (*model.User, error)
	DeleteUser(id string) error
}

type TaskRepository interface {
	GetAllTasks() []model.Task
	GetTaskByID(id string) (*model.Task, error)
	CreateTask(task model.Task) model.Task
	UpdateTask(id string, task model.Task) (*model.Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID string) []model.Task
}
