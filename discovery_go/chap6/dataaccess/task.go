package dataaccess

import (
	"errors"
	"fmt"
	"time"
)

// Deadline is a struct to hold the deadline time
type Deadline struct {
	time.Time
}

// NewDeadline returns a newly created Deadline with time t
func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

// Task is a struct to hold a single task
type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

func (s status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}
