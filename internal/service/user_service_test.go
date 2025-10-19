package service

import (
	"errors"
	"testing"

	"github.com/Wookkie/ToDoRestful/internal/model"
	"github.com/stretchr/testify/assert"
)

type fakeUserRepo struct {
	users map[string]model.User
}

func newUserRepo() *fakeUserRepo {
	return &fakeUserRepo{users: make(map[string]model.User)}
}

func (r *fakeUserRepo) GetAllUsers() []model.User {
	var res []model.User
	for _, u := range r.users {
		res = append(res, u)
	}
	return res
}
func (r *fakeUserRepo) GetUserByID(id string) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &u, nil
}
func (r *fakeUserRepo) CreateUser(user model.User) model.User {
	r.users[user.ID] = user
	return user
}
func (r *fakeUserRepo) UpdateUser(id string, user model.User) (*model.User, error) {
	if _, ok := r.users[id]; !ok {
		return nil, errors.New("not found")
	}
	r.users[id] = user
	return &user, nil
}
func (r *fakeUserRepo) DeleteUser(id string) error {
	if _, ok := r.users[id]; !ok {
		return errors.New("not found")
	}
	delete(r.users, id)
	return nil
}
func TestUserService_CreateUser(t *testing.T) {
	repo := newUserRepo()
	svc := NewUserService(repo)

	user := model.User{Name: "Alice", Email: "alice@example.com"}
	created := svc.CreateUser(user)

	assert.NotEmpty(t, created.ID, "ID должен генерироваться")
	assert.Equal(t, "Alice", created.Name)
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	repo := newUserRepo()
	svc := NewUserService(repo)

	user := model.User{ID: "1", Name: "Bob"}
	repo.users["1"] = user

	got, err := svc.GetUserByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "Bob", got.Name)
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	repo := newUserRepo()
	svc := NewUserService(repo)

	_, err := svc.GetUserByID("42")
	assert.Error(t, err, "ожидаем ошибку если пользователь не найден")
}

func TestUserService_UpdateUser(t *testing.T) {
	repo := newUserRepo()
	svc := NewUserService(repo)

	repo.users["1"] = model.User{ID: "1", Name: "Old"}

	updated, err := svc.UpdateUser("1", model.User{Name: "New"})
	assert.NoError(t, err)
	assert.Equal(t, "New", updated.Name)
}

func TestUserService_DeleteUser(t *testing.T) {
	repo := newUserRepo()
	svc := NewUserService(repo)

	repo.users["1"] = model.User{ID: "1", Name: "DeleteMe"}

	err := svc.DeleteUser("1")
	assert.NoError(t, err)
	assert.Empty(t, repo.users)
}
