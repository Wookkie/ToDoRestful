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
	UserID      string `json:"user_id"`
}
