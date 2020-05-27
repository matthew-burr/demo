package parallel

import (
	"context"
	"demo/click"
	"fmt"
	"sync"
	"time"
)

// A Worker is a function that iterates over a channel of jobs, does some work on them, and writes
// results to another channel.
func Worker(id int, jobs <-chan int, results chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var counter = click.Counter(0)

	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "cancelled.")
			return
		case <-time.After(time.Second):
			click.DoClicks(&counter, job)
			fmt.Println("Worker", id, "finished job", job)
			results <- counter.Total()
		}
	}
}
