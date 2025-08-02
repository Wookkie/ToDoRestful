package inmemory

import (
	"errors"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/google/uuid"
)

type UserMemoryRepo struct {
	users []model.User
}

func NewUserMemoryRepo() *UserMemoryRepo {
	return &UserMemoryRepo{users: []model.User{}}
}

func (r *UserMemoryRepo) GetAllUsers() []model.User {
	return r.users
}

func (r *UserMemoryRepo) GetUserByID(id string) (*model.User, error) {
	for i := range r.users {
		if r.users[i].ID == id {
			return &r.users[i], nil
		}
	}
	return nil, errors.New("not found")
}

func (r *UserMemoryRepo) CreateUser(user model.User) model.User {
	user.ID = uuid.New().String()
	r.users = append(r.users, user)
	return user
}

func (r *UserMemoryRepo) UpdateUser(id string, user model.User) (*model.User, error) {
	for i, u := range r.users {
		if u.ID == id {
			user.ID = id
			r.users[i] = user
			return &user, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *UserMemoryRepo) DeleteUser(id string) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
