package worker

import (
	"sync"
)

// Pool is the base model for tasks work group
type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// NewPool creates a worker pool given tasks and concurrency
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// all workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
