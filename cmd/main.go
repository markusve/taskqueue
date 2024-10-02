package main

import (
	"fmt"
	"time"

	"github.com/markusve/taskqueue/taskqueue"
)

func main() {
	tq := taskqueue.NewTaskQueue()
	tq.StartScheduler()

	// Add tasks
	tq.AddTask("Task 1", time.Now().Add(10*time.Second))
	tq.AddTask("Task 2", time.Now().Add(20*time.Second))

	fmt.Println("Tasks scheduled. Waiting for execution...")

	// Delete a task
	if tq.DeleteTask("Task 1") {
		fmt.Println("Task 1 deleted successfully.")
	} else {
		fmt.Println("Task 1 not found.")
	}

	// Keep the main function running
	select {}
}
