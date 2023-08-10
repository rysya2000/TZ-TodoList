package models

import (
	"regexp"
	"todolist/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoList struct {
	ID         primitive.ObjectID
	Title      string `json:"title" example:"Купить книгу"`
	ActiveAt   string `json:"activeAt" example:"2023-08-04"`
	CreateAt   string
	StatusDone bool
}

func (t *TodoList) Validate() error {
	if !t.ValidateLenOfTitle() {
		return utils.ErrLenOfTitle
	}

	if !t.ValidateActiveAt() {
		return utils.ErrActiveAt
	}

	return nil
}

func (t *TodoList) ValidateLenOfTitle() bool {
	return len([]rune(t.Title)) > 0 && len([]rune(t.Title)) <= 200
}

func (t *TodoList) ValidateActiveAt() bool {
	re := regexp.MustCompile(`^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`)
	return re.MatchString(t.ActiveAt)
}
