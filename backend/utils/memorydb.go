package utils

import (
	"sync"
)

// InMemoryDB simulates an in-memory database for tasks and users.
type InMemoryDB struct {
	tasks map[string]*Task
	users map[string]*User
	mu    sync.RWMutex
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		tasks: make(map[string]*Task),
		users: make(map[string]*User),
	}
}

// AddTask adds a new task to the in-memory database.
func (db *InMemoryDB) AddTask(task *Task) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.tasks[task.ID] = task
}

// GetTask retrieves a task by its ID.
func (db *InMemoryDB) GetTask(id string) (*Task, bool) {
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
func (db *InMemoryDB) AddUser(user *User) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.Username] = user
}

// GetUser retrieves a user by their username.
func (db *InMemoryDB) GetUser(username string) (*User, bool) {
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