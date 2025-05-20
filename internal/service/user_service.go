package service

import (
	"errors"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/google/uuid"
)

type UserService struct {
	users []model.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]model.User, 0),
	}
}

func (s *UserService) GetAllUsers() []model.User {
	return s.users
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("пользователь не найден")
}

func (s *UserService) CreateUser(user model.User) model.User {
	user.ID = uuid.New().String()
	s.users = append(s.users, user)
	return user
}

func (s *UserService) UpdateUser(id string, updated model.User) (*model.User, error) {
	for i, user := range s.users {
		if user.ID == id {
			updated.ID = id
			s.users[i] = updated
			return &updated, nil
		}
	}
	return nil, errors.New("пользователь не найден")
}

func (s *UserService) DeleteUser(id string) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("пользователь не найден")
}
