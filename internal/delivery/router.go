package delivery

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(Mux *chi.Mux, H *Handler) {
	Mux.Use(middleware.Logger)

	Mux.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("health"))
	})

	Mux.Post("/api/todo-list/tasks", H.CreateTask)
	Mux.Put("/api/todo-list/tasks/{ID}", H.UpdateTaskByID)
	Mux.Delete("/api/todo-list/tasks/{ID}", H.DeleteTaskByID)
	Mux.Put("/api/todo-list/tasks/{ID}/done", H.MarkTaskStatusByID)
	Mux.Get("/api/todo-list/tasks", H.ListTasks)

}
