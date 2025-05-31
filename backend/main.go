package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "todo-app/backend/config"
    "todo-app/backend/handlers"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize router
    r := mux.NewRouter()

    // Set up routes
    r.HandleFunc("/api/auth/signup", handlers.SignUpHandler).Methods("POST")
    r.HandleFunc("/api/auth/signin", handlers.SignInHandler).Methods("POST")
    r.HandleFunc("/api/tasks", handlers.GetTasksHandler).Methods("GET")
    r.HandleFunc("/api/tasks", handlers.CreateTaskHandler).Methods("POST")
    r.HandleFunc("/api/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
    r.HandleFunc("/api/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
    r.HandleFunc("/api/pomodoro/start", handlers.StartPomodoroHandler).Methods("POST")
    r.HandleFunc("/api/pomodoro/stop", handlers.StopPomodoroHandler).Methods("POST")

    // Start server
    log.Printf("Starting server on port %s...", cfg.ServerPort)
    if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}