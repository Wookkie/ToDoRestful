package service

import (
	"errors"

	"github.com/Wookkie/ToDoRestful/internal/model"
)

type FakeTaskRepo struct {
	tasks map[string]model.Task
}

type FakeUserRepo struct {
	users map[string]model.User
}

func NewFakeTaskRepo() *FakeTaskRepo {
	return &FakeTaskRepo{tasks: make(map[string]model.Task)}
}

func NewFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{users: make(map[string]model.User)}
}

//Tasks

func (r *FakeTaskRepo) GetAllTasks() []model.Task {
	var res []model.Task
	for _, t := range r.tasks {
		res = append(res, t)
	}
	return res
}

func (r *FakeTaskRepo) GetTaskByID(id string) (*model.Task, error) {
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &t, nil
}

func (r *FakeTaskRepo) CreateTask(task model.Task) model.Task {
	r.tasks[task.ID] = task
	return task
}

func (r *FakeTaskRepo) UpdateTask(id string, task model.Task) (*model.Task, error) {
	if _, ok := r.tasks[id]; !ok {
		return nil, errors.New("not found")
	}
	r.tasks[id] = task
	return &task, nil
}

func (r *FakeTaskRepo) DeleteTask(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return errors.New("not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *FakeTaskRepo) GetTasksByUserID(userID string) []model.Task {
	var res []model.Task
	for _, t := range r.tasks {
		if t.UserID == userID {
			res = append(res, t)
		}
	}
	return res
}

//Users

func (r *FakeUserRepo) GetAllUsers() []model.User {
	var res []model.User
	for _, u := range r.users {
		res = append(res, u)
	}
	return res
}

func (r *FakeUserRepo) GetUserByID(id string) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &u, nil
}

func (r *FakeUserRepo) CreateUser(user model.User) model.User {
	r.users[user.ID] = user
	return user
}

func (r *FakeUserRepo) UpdateUser(id string, user model.User) (*model.User, error) {
	if _, ok := r.users[id]; !ok {
		return nil, errors.New("not found")
	}
	r.users[id] = user
	return &user, nil
}

func (r *FakeUserRepo) DeleteUser(id string) error {
	if _, ok := r.users[id]; !ok {
		return errors.New("not found")
	}
	delete(r.users, id)
	return nil
}
