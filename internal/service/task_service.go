package service

import (
	"errors"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/google/uuid"
)

type TaskService struct {
	tasks []model.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make([]model.Task, 0),
	}
}

func (s *TaskService) GetAllTasks() []model.Task {
	return s.tasks
}

func (s *TaskService) GetTaskByID(id string) (*model.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("задача не найдена")
}

func (s *TaskService) CreateTask(task model.Task) model.Task { //почему здесь без указателей
	task.ID = uuid.New().String() //преобразовывает НОВЫЙ UUID в строку
	s.tasks = append(s.tasks, task)
	return task
}

// что в данной функции происходит и как
func (s *TaskService) UpdateTask(id string, updated model.Task) (*model.Task, error) { //почему что-то с указателем, а что-то без него
	for i, task := range s.tasks {
		if task.ID == id {
			updated.ID = id
			s.tasks[i] = updated
			return &updated, nil
		}
	}
	return nil, errors.New("задача не найдена")
}

func (s *TaskService) DeleteTask(id string) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("задача не найдена")
}
