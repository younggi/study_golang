package task

// ID is data type to identify a task
type ID string

// Accessor is an interface to access tasks
type Accessor interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}
