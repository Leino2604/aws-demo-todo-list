package utils

import (
	"sync"
	"todo-app/backend/models"
)

// InMemoryDB simulates an in-memory database for tasks and users.
type InMemoryDB struct {
	tasks map[string]*models.Task
	users map[string]*models.User
	mu    sync.RWMutex
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		tasks: make(map[string]*models.Task),
		users: make(map[string]*models.User),
	}
}

// AddTask adds a new task to the in-memory database.
func (db *InMemoryDB) AddTask(task *models.Task) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.tasks[task.ID] = task
}

// GetTask retrieves a task by its ID.
func (db *InMemoryDB) GetTask(id string) (*models.Task, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	task, exists := db.tasks[id]
	return task, exists
}

// DeleteTask removes a task from the in-memory database.
func (db *InMemoryDB) DeleteTask(id string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.tasks, id)
}

// AddUser adds a new user to the in-memory database.
func (db *InMemoryDB) AddUser(user *models.User) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.Username] = user
}

// GetUser retrieves a user by their username.
func (db *InMemoryDB) GetUser(username string) (*models.User, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	user, exists := db.users[username]
	return user, exists
}

// DeleteUser removes a user from the in-memory database.
func (db *InMemoryDB) DeleteUser(username string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.users, username)
}
