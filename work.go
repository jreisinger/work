// Package work generates tasks from lines of STDIN, processes them concurrently
// and prints to STDOUT. To use it you just need to implement Generator and Task
// interfaces.
package work

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// Factory generates tasks from lines supplied on STDIN.
type Factory interface {
	Generate(line string) Task
}

// Task is anything that can be processed and print the result to STDOUT.
type Task interface {
	Process()
	Print()
}

// Run spawns factory and workers. Factory generates tasks that are load
// balanced among workers who process them. When task is processed its output is
// printed.
func Run(g Factory, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task)
	out := make(chan Task)

	// Generate tasks that will be processed.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- g.Generate(s.Text())
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
		t.Print()
	}
}
