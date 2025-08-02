package inmemory

import (
	"errors"
	"strconv"
	"sync"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/repository"
)

type TaskMemoryRepo struct {
	mu     sync.RWMutex
	tasks  map[string]model.Task
	nextID int
}

func NewTaskMemoryRepo() repository.TaskRepository {
	return &TaskMemoryRepo{
		tasks:  make(map[string]model.Task),
		nextID: 1,
	}
}

func (r *TaskMemoryRepo) CreateTask(task model.Task) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	idStr := strconv.Itoa(r.nextID)
	r.nextID++
	r.tasks[idStr] = task

	return task
}

func (r *TaskMemoryRepo) GetTaskByID(id string) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (r *TaskMemoryRepo) GetAllTasks() []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *TaskMemoryRepo) UpdateTask(id string, task model.Task) (*model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return nil, errors.New("task not found")
	}

	task.ID = id
	r.tasks[id] = task
	return &task, nil
}

func (r *TaskMemoryRepo) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *TaskMemoryRepo) GetTasksByUserID(userID string) []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tasks []model.Task
	for _, task := range r.tasks {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
