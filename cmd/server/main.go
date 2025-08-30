package main

import (
	"log"
	"monolithic-go-app/pkg/config"
	"monolithic-go-app/pkg/database"
	"monolithic-go-app/pkg/server"
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

	srv := server.NewServer(db)
	if err := srv.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	//
	//r := mux.NewRouter()
	//
	//r.Use(middleware.LoggingMiddleware)
	//
	//userHandler := controllers.NewUserHandler(db)
	//r.HandleFunc("/users", userHandler.GetAllUsers).Methods("Get")
	//r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("Get")
	//r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	//r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	//
	//log.Println("Server starting on :8080")
	//log.Fatal(http.ListenAndServe(":8080", r))
}
