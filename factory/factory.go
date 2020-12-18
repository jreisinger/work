// Package factory implements a scalable work system. It can generate and work
// on many tasks concurrently.
package factory

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// Task is anything that can be processed and generate some output.
type Task interface {
	Process()
	Output()
}

// Boss creates tasks from lines of text supplied on STDIN.
type Boss interface {
	Create(line string) Task
}

// Run spawns a Boss and n workers. Boss generates tasks and workers work on
// them.
func Run(b Boss, n int) {
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
			in <- b.Create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %s", s.Err())
		}
	}()

	// Create workers to process the tasks.
	for i := 0; i < n; i++ {
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
