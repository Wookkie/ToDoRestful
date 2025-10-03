package service_test

import (
	"errors"
	"testing"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/service"
	"github.com/stretchr/testify/assert"
)

type fakeTaskRepo struct {
	tasks map[string]model.Task
}

func newTaskRepo() *fakeTaskRepo {
	return &fakeTaskRepo{tasks: make(map[string]model.Task)}
}

func (r *fakeTaskRepo) GetAllTasks() []model.Task {
	var res []model.Task
	for _, t := range r.tasks {
		res = append(res, t)
	}
	return res
}
func (r *fakeTaskRepo) GetTaskByID(id string) (*model.Task, error) {
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &t, nil
}
func (r *fakeTaskRepo) CreateTask(task model.Task) model.Task {
	r.tasks[task.ID] = task
	return task
}
func (r *fakeTaskRepo) UpdateTask(id string, task model.Task) (*model.Task, error) {
	if _, ok := r.tasks[id]; !ok {
		return nil, errors.New("not found")
	}
	r.tasks[id] = task
	return &task, nil
}
func (r *fakeTaskRepo) DeleteTask(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return errors.New("not found")
	}
	delete(r.tasks, id)
	return nil
}
func (r *fakeTaskRepo) GetTasksByUserID(userID string) []model.Task {
	var res []model.Task
	for _, t := range r.tasks {
		if t.UserID == userID {
			res = append(res, t)
		}
	}
	return res
}

func TestTaskService_CreateTask(t *testing.T) {
	repo := newTaskRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{Title: "Test", UserID: "u1"}
	created := svc.CreateTask(task)

	assert.NotEmpty(t, created.ID, "ID должен генерироваться")
	assert.Equal(t, "u1", created.UserID)
}

func TestTaskService_GetTaskByID_Success(t *testing.T) {
	repo := newTaskRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{ID: "1", Title: "Test", UserID: "u1"}
	repo.tasks["1"] = task

	got, err := svc.GetTaskByID("1", "u1")
	assert.NoError(t, err)
	assert.Equal(t, "Test", got.Title)
}

func TestTaskService_GetTaskByID_WrongUser(t *testing.T) {
	repo := newTaskRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{ID: "1", Title: "Test", UserID: "u1"}
	repo.tasks["1"] = task

	_, err := svc.GetTaskByID("1", "u2")
	assert.Error(t, err, "чужую задачу получить нельзя")
}

func TestTaskService_DeleteTask_Success(t *testing.T) {
	repo := newTaskRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{ID: "1", UserID: "u1"}
	repo.tasks["1"] = task

	err := svc.DeleteTask("1", "u1")
	assert.NoError(t, err)
	assert.Empty(t, repo.tasks)
}

func TestTaskService_DeleteTask_WrongUser(t *testing.T) {
	repo := newTaskRepo()
	svc := service.NewTaskService(repo)

	task := model.Task{ID: "1", UserID: "u1"}
	repo.tasks["1"] = task

	err := svc.DeleteTask("1", "u2")
	assert.Error(t, err, "чужую задачу удалить нельзя")
}
