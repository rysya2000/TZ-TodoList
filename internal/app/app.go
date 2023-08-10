package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todolist/config"
	"todolist/internal/delivery"
	"todolist/internal/usecase"
	"todolist/internal/usecase/repo"
	"todolist/pkg/mongodb"

	"github.com/go-chi/chi"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg)

	// MongoDB
	client, err := mongodb.New(cfg.MongoDB.URL)
	if err != nil {
		log.Fatalf("MongoDB error: %v\n", err)
	}
	defer client.Disconnect(context.TODO())

	repo := repo.NewRepo(client.Client)
	usecase := usecase.New(repo)
	handler := delivery.New(*usecase)

	// HTTP Server
	Mux := chi.NewRouter()
	delivery.NewRouter(Mux, handler)

	log.Printf("Starting up on http://localhost:%s", cfg.HTTP.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.HTTP.Port, Mux))
}
