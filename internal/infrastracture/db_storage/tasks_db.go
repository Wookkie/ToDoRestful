package dbstorage

import (
	"context"
	"errors"
	"time"

	"github.com/Wookkie/ToDoRestful/internal/model"
)

func (db *DBStorage) GetAllTasks() []model.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.db.Query(ctx, "SELECT id, title, description, status, user_id FROM tasks")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.UserID)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (db *DBStorage) GetTaskByID(id string) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task model.Task
	err := db.db.QueryRow(ctx, "SELECT id, title, description, status, user_id FROM tasks WHERE id=$1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.UserID)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func (db *DBStorage) CreateTask(task model.Task) model.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.db.Exec(ctx,
		"INSERT INTO tasks (id, title, description, status, user_id) VALUES ($1, $2, $3, $4, $5)",
		task.ID, task.Title, task.Description, task.Status, task.UserID)
	if err != nil {
		return model.Task{}
	}
	return task
}

func (db *DBStorage) UpdateTask(id string, task model.Task) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmdTag, err := db.db.Exec(ctx,
		"UPDATE tasks SET title=$1, description=$2, status=$3, user_id=$4 WHERE id=$5",
		task.Title, task.Description, task.Status, task.UserID, id)
	if err != nil {
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, errors.New("task not found")
	}
	task.ID = id
	return &task, nil
}

func (db *DBStorage) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmdTag, err := db.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (db *DBStorage) GetTasksByUserID(userID string) []model.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.db.Query(ctx, "SELECT id, title, description, status, user_id FROM tasks WHERE user_id=$1", userID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.UserID)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}
