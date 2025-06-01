package main

import (
    "log"
    "net/http"
    "todo-app/backend/internal/handlers"
)

func main() {
    http.HandleFunc("/api/login", handlers.LoginHandler)
    http.HandleFunc("/api/todos", handlers.TodoHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}