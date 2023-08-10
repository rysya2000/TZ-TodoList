package delivery

import "todolist/internal/usecase"

type Handler struct {
	usecase usecase.UseCase
}

func New(u usecase.UseCase) *Handler {
	return &Handler{
		usecase: u,
	}
}
