package prob4

import (
	"fmt"
	"sort"
	"time"
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d *Deadline) OverDue() bool {
	return d != nil && d.Before(time.Now())
}

type status int

const (
	UNKNOWN status = iota // 0
	TODO                  // 1
	DONE                  // 2
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

type DeadlineTask []Task

func (t DeadlineTask) Len() int {
	return len(t)
}

func (t DeadlineTask) Less(i, j int) bool {
	return t[i].Deadline.Unix() < t[j].Deadline.Unix()
}

func (t DeadlineTask) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func ExampleDeadlineTaskSort() {
	tasks := DeadlineTask{
		{"task1", TODO, NewDeadline(time.Now().Add(time.Hour * 3)), 2, nil},
		{"task2", TODO, NewDeadline(time.Now().Add(time.Hour * 2)), 1, nil},
		{"task3", TODO, NewDeadline(time.Now().Add(time.Hour * 1)), 3, nil},
	}
	sort.Sort(tasks)
	for _, t := range tasks {
		fmt.Println(t)
	}
	// Output:
}

type PriorityTask []Task

func (t PriorityTask) Len() int {
	return len(t)
}

func (t PriorityTask) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority
}

func (t PriorityTask) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func ExamplePriorityTaskSort() {
	tasks := PriorityTask{
		{"task1", TODO, NewDeadline(time.Now().Add(time.Hour * 3)), 2, nil},
		{"task2", TODO, NewDeadline(time.Now().Add(time.Hour * 2)), 1, nil},
		{"task3", TODO, NewDeadline(time.Now().Add(time.Hour * 1)), 3, nil},
	}
	sort.Sort(tasks)
	for _, t := range tasks {
		fmt.Println(t)
	}
	// Output:
}
