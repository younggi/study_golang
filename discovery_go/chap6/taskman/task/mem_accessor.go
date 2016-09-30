package task

import (
	"errors"
	"fmt"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	tasks  map[ID]Task
	nextID int64
}

// NewInMemoryAccessor returns a new InMemoryAccessor.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		tasks:  map[ID]Task{},
		nextID: int64(1),
	}
}

// ErrTaskNotExist occurs when the task with the ID was not found.
var ErrTaskNotExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *InMemoryAccessor) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put updates a task with a given ID with t
func (m *InMemoryAccessor) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new
func (m *InMemoryAccessor) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID.
func (m *InMemoryAccessor) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

// GetAll gets all of the tasks
func (m *InMemoryAccessor) GetAll() ([]Task, error) {
	v := make([]Task, len(m.tasks))

	for _, value := range m.tasks {
		v = append(v, value)
	}
	return v, nil
}
