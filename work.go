// Package work concurrently generates and processes tasks. The tasks are
// generated from lines supplied on STDIN. The results of tasks processing are
// then printed on STDOUT. To use it you just need to implement Factory and Task
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

// Task is anything that can be processed and print the result on STDOUT.
type Task interface {
	Process()
	Print()
}

// Run concurrently runs factory and n workers. Factory generates tasks that are
// load balanced among workers. Workers process the tasks. When all tasks are
// processed the results are printed.
func Run(f Factory, n int) {
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
			in <- f.Generate(s.Text())
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
		t.Print()
	}
}
