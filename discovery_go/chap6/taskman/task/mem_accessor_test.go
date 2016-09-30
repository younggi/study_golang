package task

import (
	"testing"
	"time"
)

// Test Empty
func TestInMemoryAccessor_Get(t *testing.T) {
	ma := NewInMemoryAccessor()
	task, err := ma.Get("1")
	if task.Title != "" && err != ErrTaskNotExist {
		t.Errorf("Result must be Error")
	}
}

// Test POST, GET
func TestInMemoryAccessor_Post(t *testing.T) {
	ma := NewInMemoryAccessor()
	task := Task{"Test Task", TODO, NewDeadline(time.Now()), 0, nil}
	id, err := ma.Post(task)
	if err != nil {
		t.Error("Post Task Error")
	}
	task1, err := ma.Get(id)
	if err != nil {
		t.Error("Get Task Error")
	}
	if !isSameTask(task, task1) {
		t.Error("Get Task is not same")
	}
}

func isSameTask(s Task, t Task) bool {
	return s.Title == t.Title &&
		s.Status == t.Status &&
		s.Priority == t.Priority &&
		s.Deadline.Unix() == t.Deadline.Unix()
}

// Test Put, Get
func TestInMemoryAccessor_Put(t *testing.T) {
	ma := NewInMemoryAccessor()
	id, err := ma.Post(Task{})
	if err != nil {
		t.Error("Post Task Error")
	}
	task := Task{"Test Task", TODO, NewDeadline(time.Now()), 0, nil}
	err = ma.Put(id, task)
	if err != nil {
		t.Error("Put Task Error")
	}
	task1, err := ma.Get(id)
	if err != nil {
		t.Error("Get Task Error")
	}
	if !isSameTask(task, task1) {
		t.Error("Get Task is not same")
	}
}

// Test Delete
func TestInMemoryAccessor_Delete(t *testing.T) {
	ma := NewInMemoryAccessor()
	id, err := ma.Post(Task{})
	if err != nil {
		t.Error("Post Task Error")
	}
	err = ma.Delete(id)
	if err != nil {
		t.Error("Delete Task Error")
	}
	_, err = ma.Get(id)
	if err != ErrTaskNotExist {
		t.Error("Delete Task Error")
	}
}

// Test GetAll
func TestInMemoryAccessor_GetAll(t *testing.T) {
	ma := NewInMemoryAccessor()
	for i := 0; i < 10; i++ {
		_, err := ma.Post(Task{
			"Test Task",
			TODO,
			NewDeadline(time.Now()),
			i,
			nil,
		})
		if err != nil {
			t.Error("Post Task Error")
		}
	}
	tasks, err := ma.GetAll()
	if err != nil {
		t.Error("GetAll Task Error")
	}
	if len(tasks) != 10 {
		t.Error("GetAll Task Count Error")
	}
}
