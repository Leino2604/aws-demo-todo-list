package memory

import (
	"errors"
	"sync"
	"time"
)

type Todo struct {
	ID          string
	Description string
	Estimated   time.Duration
	StartedAt   time.Time
	Completed   bool
}

type TodoRepository struct {
	mu    sync.Mutex
	todos map[string]*Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make(map[string]*Todo),
	}
}

func (r *TodoRepository) Create(todo *Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[todo.ID]; exists {
		return errors.New("todo already exists")
	}

	r.todos[todo.ID] = todo
	return nil
}

func (r *TodoRepository) GetAll() []*Todo {
	r.mu.Lock()
	defer r.mu.Unlock()

	var todos []*Todo
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (r *TodoRepository) Update(todo *Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	r.todos[todo.ID] = todo
	return nil
}

func (r *TodoRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(r.todos, id)
	return nil
}