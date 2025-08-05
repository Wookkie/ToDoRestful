package service

import (
	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/Wookkie/ToDoRestful/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []model.User {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user model.User) model.User {
	user.ID = uuid.New().String()
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id string, updated model.User) (*model.User, error) {
	updated.ID = id
	return s.repo.UpdateUser(id, updated)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
