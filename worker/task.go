package worker

import "sync"

// Task is the generic task model
type Task struct {
	Error error
	f     func() error
}

// NewTask initializes a new task based on a given work function.
func NewTask(f func() error) *Task {
	return &Task{f: f}
}

// Run runs a task and marks wait group as done
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Error = t.f()
	wg.Done()
}
