// The parallel package demonstrates parallelization using
// goroutines and channels. I've based it on an example at
// https://gobyexample.com/worker-pools.
//
// What's Included
//
// The parallel package includes examples of:
//
// - Goroutines
//
// - Channels
//
// - Buffered Channels
//
// - Directed Channels
//
// - Channel Multiplexing
//
// - Deferred Calls
//
// - Contexts
//
// - Synchronization WaitGroups
//
// - Timers
package parallel

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
)

// The Command structure defines the CLI for a parallel demo command.
type Command struct {
	JobCount    int `arg optional name:"jobs" help:"Specifies the number of jobs to execute." default:"5"`
	WorkerCount int `short:"w" name:"workers" help:"Specifies the number of workers to create." default:"3"`
	jobs        chan int
	results     chan int
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	total       int
}

// Run executes the parallel demo command.
func (cmd *Command) Run() error {
	cmd.listenForCancel()
	cmd.createChannels()
	cmd.startWorkers()
	cmd.distributeJobs()
	cmd.gatherResults()
	cmd.waitForWorkerShutdown()
	cmd.printResult()
	return nil
}

func (cmd *Command) createChannels() {
	// We create two channels: jobs is where we send jobs to workers, and results
	// are where the workers send their results back to us.
	cmd.jobs, cmd.results = make(chan int, cmd.JobCount), make(chan int, cmd.JobCount)
}

func (cmd *Command) listenForCancel() {
	// We start up a process to listen for the user to enter the "q" or "Q" keys
	// as a way to cancel early.
	cmd.ctx, cmd.cancel = context.WithCancel(context.Background())

	go func() {
		defer cmd.cancel()
		var in = bufio.NewReader(os.Stdin)
		for {
			select {
			case <-cmd.ctx.Done():
				return
			default:
				switch c, _, err := in.ReadRune(); {
				case err != nil:
					fmt.Println("error listening for cancel:", err)
					return
				case c == 'q' || c == 'Q':
					return
				}
			}
		}
	}()
}

func (cmd *Command) startWorkers() {
	// Here, we create a number of "workers". These are goroutines, which we
	// can think of as threads. To help us keep track of when all the threads
	// are done, we us a synchronization wait group.
	for id := 1; id <= cmd.WorkerCount; id++ {
		cmd.wg.Add(1)
		go Worker(id, cmd.jobs, cmd.results, cmd.ctx, &cmd.wg)
	}
}

func (cmd *Command) distributeJobs() {
	// We create a number of jobs, which we submit to our "jobs" channel.
	for job := 1; job <= cmd.JobCount; job++ {
		cmd.jobs <- job
	}
	// We close the channel as a signal that we're done submitting jobs.
	close(cmd.jobs)
}

func (cmd *Command) gatherResults() {
	// We'll read all the results, and sum them up for output.
	defer close(cmd.results)
	for result := 1; result <= cmd.JobCount; result++ {
		// Here, we use channel multi-plexing to listen to multiple channels
		// at once and take action based on the first one to respond.
		select {
		case r := <-cmd.results:
			cmd.total += r
		case <-cmd.ctx.Done():
			return
		}
	}
}

func (cmd *Command) waitForWorkerShutdown() {
	// We wait for all of the workers to signal that they're done before
	// we report the result. This ensures that if we're stopping due to cancellation,
	// all the cancellations have been processed.
	cmd.wg.Wait()
}

func (cmd *Command) printResult() {
	// We print out the total of the workers' work.
	fmt.Println("Total of all results:", cmd.total)
}
