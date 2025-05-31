package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"todo-app/backend/config"
	"todo-app/backend/handlers"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize router
	r := mux.NewRouter()

	// Create handler instances
	authHandler := handlers.NewAuthHandler()
	tasksHandler := handlers.NewTasksHandler()
	pomodoroHandler := handlers.NewPomodoroHandler(25*time.Minute, 5*time.Minute, 4)

	// Set up routes
	api := r.PathPrefix("/api").Subrouter()
	// Auth routes
	api.HandleFunc("/auth/signup", authHandler.SignUp).Methods("POST")
	api.HandleFunc("/auth/signin", authHandler.SignIn).Methods("POST")
	// Task routes
	api.HandleFunc("/tasks", tasksHandler.GetTasks).Methods("GET")
	api.HandleFunc("/tasks", tasksHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", tasksHandler.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", tasksHandler.DeleteTask).Methods("DELETE")
	// Pomodoro routes
	api.HandleFunc("/pomodoro/start", pomodoroHandler.StartPomodoro).Methods("POST")
	api.HandleFunc("/pomodoro/stop", pomodoroHandler.StopPomodoro).Methods("POST")

	// Start server
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}