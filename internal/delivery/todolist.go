package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/internal/models"
	"todolist/internal/utils"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.TodoList
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = h.usecase.TodoList.Create(context.TODO(), req)
	if err != nil {
		switch err {
		case utils.ErrLenOfTitle:
			http.Error(w, err.Error(), 404)
		case utils.ErrActiveAt:
			http.Error(w, err.Error(), 404)
		default:
			fmt.Printf("internal error: %v\n", err.Error())
		}
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	var req models.TodoList
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = h.usecase.TodoList.UpdateByID(context.TODO(), req)
	if err != nil {
		switch err {
		case utils.ErrLenOfTitle:
			http.Error(w, err.Error(), 404)
		case utils.ErrActiveAt:
			http.Error(w, err.Error(), 404)
		default:
			fmt.Printf("internal error: %v\n", err.Error())
		}
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID") // ID is string here

	// Convert the provided ID string to an ObjectId
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		http.Error(w, utils.ErrIDIsNotNum.Error(), 404)
		return
	}

	td := models.TodoList{
		ID: objID,
	}

	// check if there is task with that ID
	_, err = h.usecase.TodoList.Read(context.TODO(), td)
	if err != nil {
		http.Error(w, utils.ErrNotFound.Error(), 400)
		return
	}

	// deleting task
	err = h.usecase.TodoList.Delete(context.TODO(), td)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Printf("internal error: %v\n", err.Error())
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) MarkTaskStatusByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID") // ID is string here

	// Convert the provided ID string to an ObjectId
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		http.Error(w, utils.ErrIDIsNotNum.Error(), 404)
		return
	}

	td := models.TodoList{
		ID: objID,
	}

	_, err = h.usecase.TodoList.Read(context.TODO(), td)
	if err != nil {
		http.Error(w, utils.ErrNotFound.Error(), 400)
		return
	}

	err = h.usecase.TodoList.MarkStatusByID(context.TODO(), td)
	if err != nil {
		http.Error(w, utils.ErrInternalError.Error(), 500)
		fmt.Printf("internal error: %v\n", err.Error())
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.usecase.TodoList.List(context.TODO())
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Printf("internal error: %v\n", err.Error())
		return
	}

	fmt.Println("List: ", tasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Printf("internal error: %v\n", err.Error())
		return
	}
}
