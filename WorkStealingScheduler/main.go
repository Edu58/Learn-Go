package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task func()

type Deque struct {
	mu    sync.Mutex
	tasks []Task
}

type WorkerPool struct {
	workers []*Worker
}

type Worker struct {
	id    int
	deque *Deque
	pool  *WorkerPool
}

func (d *Deque) PushFront(task Task) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.tasks = append([]Task{task}, d.tasks...)
}

func (d *Deque) PopFront() (Task, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if len(d.tasks) <= 0 {
		return nil, false
	}

	task := d.tasks[0]
	d.tasks = d.tasks[1:]
	return task, true
}

func (d *Deque) PopBack() (Task, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if len(d.tasks) <= 0 {
		return nil, false
	}

	task := d.tasks[len(d.tasks)-1]
	d.tasks = d.tasks[:len(d.tasks)-1]
	return task, true
}

func (w *Worker) Start() {
	go func() {
		for {
			if task, ok := w.deque.PopFront(); ok {
				task()
			} else {
				if task, ok := w.pool.Steal(w.id); ok {
					task()
				} else {
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
	}()
}

func NewWorkerPool(numOfWorkers int) *WorkerPool {
	pool := &WorkerPool{workers: make([]*Worker, numOfWorkers)}
	for i := range numOfWorkers {
		pool.workers[i] = &Worker{id: i, deque: &Deque{}, pool: pool}
		pool.workers[i].Start()
	}
	return pool
}

func (wp *WorkerPool) Steal(thiefID int) (Task, bool) {
	fmt.Printf("Worker %d is stealing work\n", thiefID)

	for i := range len(wp.workers) {
		if i != thiefID {
			if task, ok := wp.workers[i].deque.PopBack(); ok {
				return task, true
			}
		}
	}

	return nil, false
}

func (p *WorkerPool) SubmitTask(task Task) {
	worker := p.workers[rand.Intn(len(p.workers))]
	worker.deque.PushFront(task)
}

func main() {
	pool := NewWorkerPool(2)

	for i := range 50 {
		taskId := i
		pool.SubmitTask(func() {
			fmt.Printf("Task %d is being executed\n", taskId)
		})
	}

	// Give time to execute the tasks.
	// Only for testing so as to view the results
	time.Sleep(2 * time.Second)
}
