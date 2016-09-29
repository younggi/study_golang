package dataaccess

import (
	"errors"
	"fmt"
)

// ID is data type to identify a task
type ID string

// DataAccess is an interface to access tasks
type DataAccess interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}

// MemoryDataAccess is a simple in-memory database.
type MemoryDataAccess struct {
	tasks  map[ID]Task
	nextID int64
}

// NewMemoryDataAccess returns a new MemoryDataAccess.
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]Task{},
		nextID: int64(1),
	}
}

// ErrTaskNotExist occurs when the task with the ID was not found.
var ErrTaskNotExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *MemoryDataAccess) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put updates a task with a given ID with t
func (m *MemoryDataAccess) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new
func (m *MemoryDataAccess) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID.
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}
