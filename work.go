// Package work is a scalable work system. It can generate and work
// on many tasks concurrently. To Run a factory you need a Boss and a Task.
package work

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// Generator generates tasks from lines supplied on STDIN.
type Generator interface {
	Generate(line string) Task
}

// Task needs to be processed and print something to STDOUT.
type Task interface {
	Process()
	Output()
}

// Do spawns generator and workers. Generator generates tasks that are load
// balanced among workers.
func Do(b Generator, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task)
	out := make(chan Task)

	// Create tasks that will be processed.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- b.Generate(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %s", s.Err())
		}
	}()

	// Create workers to process the tasks.
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {
				t.Process()
				out <- t
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Output()
	}
}
