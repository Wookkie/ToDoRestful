package dbstorage

import (
	"context"
	"time"

	"github.com/Wookkie/ToDoRestful/internal/model"
)

func (db *DBStorage) GetAllUsers() []model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			continue
		}
		users = append(users, user)
	}

	return users
}

func (db *DBStorage) GetUserByID(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := db.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DBStorage) CreateUser(user model.User) model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.db.Exec(ctx, "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return model.User{}
	}
	return user
}

func (db *DBStorage) UpdateUser(id string, user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.db.Exec(ctx, "UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DBStorage) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
