package delivery

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(Mux *chi.Mux, Handler *Handler) {
	Mux.Use(middleware.Logger)

	Mux.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("health"))
	})

	Mux.Post("/api/todo-list/tasks", nil)
	Mux.Put("/api/todo-list/tasks/{ID}", nil)
	Mux.Delete("/api/todo-list/tasks/{ID}", nil)
	Mux.Put("/api/todo-list/tasks/{ID}/done", nil)
	Mux.Get("/api/todo-list/tasks", nil)

}
