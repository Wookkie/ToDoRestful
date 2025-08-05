package service

import (
	"errors"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/repository"
	"github.com/google/uuid"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() []model.Task {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id string, userID string) (*model.Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, errors.New("не получилось получить доступ к задаче")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task model.Task) model.Task {
	task.ID = uuid.New().String()
	return s.repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(id string, updated model.Task, userID string) (*model.Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, errors.New("не получилось получить доступ к задаче")
	}

	updated.ID = id
	updated.UserID = userID
	return s.repo.UpdateTask(id, updated)
}

func (s *TaskService) DeleteTask(id string, userID string) error {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("не получилось получить доступ к задаче")
	}
	return s.repo.DeleteTask(id)
}

func (s *TaskService) GetTasksByUserID(userID string) []model.Task {
	return s.repo.GetTasksByUserID(userID)
}
