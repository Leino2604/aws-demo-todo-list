package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"todo-app/backend/models"
)

type TasksHandler struct {
	Tasks map[string]models.Task
}

func NewTasksHandler() *TasksHandler {
	return &TasksHandler{
		Tasks: make(map[string]models.Task),
	}
}

func (h *TasksHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.Tasks[task.ID] = task
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TasksHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = id
	h.Tasks[id] = task
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TasksHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	delete(h.Tasks, id)
	w.WriteHeader(http.StatusNoContent)
}

func (h *TasksHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]models.Task, 0, len(h.Tasks))
	for _, task := range h.Tasks {
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}