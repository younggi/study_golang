package task

import (
	"errors"
	"fmt"
	"sync"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	tasks  map[ID]Task
	nextID int64
	L      *sync.RWMutex
}

// NewInMemoryAccessor returns a new InMemoryAccessor.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		tasks:  map[ID]Task{},
		nextID: int64(1),
		L:      &sync.RWMutex{},
	}
}

// ErrTaskNotExist occurs when the task with the ID was not found.
var ErrTaskNotExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *InMemoryAccessor) Get(id ID) (Task, error) {
	m.L.RLock()
	t, exists := m.tasks[id]
	m.L.RUnlock()
	if !exists {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put updates a task with a given ID with t
func (m *InMemoryAccessor) Put(id ID, t Task) error {
	m.L.Lock()
	defer m.L.Unlock()
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new
func (m *InMemoryAccessor) Post(t Task) (ID, error) {
	m.L.Lock()
	defer m.L.Unlock()
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID.
func (m *InMemoryAccessor) Delete(id ID) error {
	m.L.Lock()
	defer m.L.Unlock()
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

// GetAll gets all of the tasks
func (m *InMemoryAccessor) GetAll() ([]Task, error) {
	m.L.RLock()
	v := []Task{}
	for _, value := range m.tasks {
		v = append(v, value)
	}
	m.L.RUnlock()
	return v, nil
}
