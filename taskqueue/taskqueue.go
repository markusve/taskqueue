package taskqueue

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Context string
	Time    time.Time
}

type TaskQueue struct {
	tasks []Task
	mu    sync.Mutex
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make([]Task, 0),
	}
}

func (tq *TaskQueue) AddTask(context string, triggerTime time.Time) {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	tq.tasks = append(tq.tasks, Task{Context: context, Time: triggerTime})
}

func (tq *TaskQueue) DeleteTask(context string) bool {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	for i, task := range tq.tasks {
		if task.Context == context {
			tq.tasks = append(tq.tasks[:i], tq.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (tq *TaskQueue) StartScheduler() {
	go func() {
		for {
			tq.mu.Lock()
			now := time.Now()
			for i := 0; i < len(tq.tasks); i++ {
				if tq.tasks[i].Time.Before(now) || tq.tasks[i].Time.Equal(now) {
					fmt.Println("Executing task:", tq.tasks[i].Context)
					tq.tasks = append(tq.tasks[:i], tq.tasks[i+1:]...)
					i-- // Adjust index after removal
				}
			}
			tq.mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()
}
