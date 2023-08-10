package usecase

import (
	"context"
	"todolist/internal/models"
)

type TodoList struct {
	repo CRUD
}

func NewTodoList(repo CRUD) *TodoList {
	return &TodoList{repo}
}

func (t *TodoList) Create(ctx context.Context, td models.TodoList) error {

	return t.repo.Create(ctx, td)
}

func (t *TodoList) Read(ctx context.Context, td models.TodoList) (models.TodoList, error) {

	return t.repo.Read(ctx, td)
}

func (t *TodoList) UpdateByID(ctx context.Context, td models.TodoList) error {

	return t.repo.UpdateByID(ctx, td)
}

func (t *TodoList) MarkStatusByID(ctx context.Context, td models.TodoList) error {

	return t.repo.MarkStatusByID(ctx, td)
}

func (t *TodoList) List(ctx context.Context) ([]models.TodoList, error) {

	return t.repo.List(ctx)
}

func (t *TodoList) Delete(ctx context.Context, td models.TodoList) error {

	return t.repo.Delete(ctx, td)
}
