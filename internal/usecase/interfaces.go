package usecase

import (
	"context"
	"todolist/internal/models"
)

type (
	CRUD interface {
		Create(context.Context, models.TodoList) error
		Read(context.Context, models.TodoList) (models.TodoList, error)
		UpdateByID(context.Context, models.TodoList) error
		MarkStatusByID(context.Context, models.TodoList) error
		List(context.Context) ([]models.TodoList, error)
		Delete(context.Context, models.TodoList) error
	}
)
