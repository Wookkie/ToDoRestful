package model

type Status string

const (
	StatusNew        Status = "Новая"
	StatusInProgress Status = "В процессе"
	StatusDone       Status = "Завершена"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

// Далее: task_repository.go — slice хранения и методы: Create, Get, Update, Delete
// task_service.go — вызывает методы репозитория, бизнес-правила
// task_handler.go — парсит запросы, вызывает сервис и отдает ответы через Gin
// main.go — инициализация Gin, роутинг, запуск сервера
