package main

import (
	"log"
	"monolithic-go-app/internal/handlers"
	"monolithic-go-app/internal/middleware"
	"monolithic-go-app/pkg/config"
	"monolithic-go-app/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)

	}
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	userHandler := handlers.NewUserHandler(db)
	r.HandleFunc("/api/users", userHandler.GetAllUsers).Methods("Get")
	r.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("Get")
	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
