package inmemory

import (
	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/repository"
)

type MemoryRepo struct {
	userRepo repository.UserRepository
	taskRepo repository.TaskRepository
}

func New() *MemoryRepo {
	return &MemoryRepo{
		userRepo: NewUserMemoryRepo(),
		taskRepo: NewTaskMemoryRepo(),
	}
}

func (r *MemoryRepo) GetAllUsers() []model.User {
	return r.userRepo.GetAllUsers()
}

func (r *MemoryRepo) GetUserByID(id string) (*model.User, error) {
	return r.userRepo.GetUserByID(id)
}

func (r *MemoryRepo) CreateUser(user model.User) model.User {
	return r.userRepo.CreateUser(user)
}

func (r *MemoryRepo) UpdateUser(id string, user model.User) (*model.User, error) {
	return r.userRepo.UpdateUser(id, user)
}

func (r *MemoryRepo) DeleteUser(id string) error {
	return r.userRepo.DeleteUser(id)
}

func (r *MemoryRepo) CreateTask(task model.Task) model.Task {
	return r.taskRepo.CreateTask(task)
}

func (r *MemoryRepo) GetTaskByID(id string) (*model.Task, error) {
	return r.taskRepo.GetTaskByID(id)
}

func (r *MemoryRepo) GetAllTasks() []model.Task {
	return r.taskRepo.GetAllTasks()
}

func (r *MemoryRepo) UpdateTask(id string, task model.Task) (*model.Task, error) {
	return r.taskRepo.UpdateTask(id, task)
}

func (r *MemoryRepo) DeleteTask(id string) error {
	return r.taskRepo.DeleteTask(id)
}

func (m *MemoryRepo) Close() error {
	return nil
}

func (r *MemoryRepo) GetTasksByUserID(userID string) []model.Task {
	return r.taskRepo.GetTasksByUserID(userID)
}
