// Package work is useful for building CLI tools that need to run many tasks
// concurrently. To use it implement Factory and Task interfaces. Then Run the
// Factory.
//
// Taken from https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014.
package work

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// Run concurrently runs Factory and workers. Factory generates tasks from lines
// read from filenames or from STDIN if filenames is empty. The tasks are load
// balanced among workers that process them. When a task is processed its result
// is printed on STDOUT.
func Run(f Factory, workers int, filenames []string) {
	var wg sync.WaitGroup
	in := make(chan Task)
	out := make(chan Task)

	// Generate tasks that will be processed.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		if len(filenames) == 0 {
			if err := genTask(os.Stdin, f, in); err != nil {
				log.Println(err)
			}
		} else {
			for _, arg := range filenames {
				file, err := os.Open(arg)
				if err != nil {
					log.Println(err)
					continue
				}
				defer file.Close()
				if err := genTask(file, f, in); err != nil {
					log.Println(err)
					continue
				}
			}
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

// getTask generates Task from file and send it to in channel. It uses factory
// to generate the Task.
func genTask(file *os.File, factory Factory, in chan Task) error {
	s := bufio.NewScanner(file)
	for s.Scan() {
		in <- factory.Generate(s.Text())
	}
	return s.Err()
}

// Factory generates task from line read from file or from STDIN.
type Factory interface {
	Generate(line string) Task
}

// Task is anything that can be processed and its result printed on STDOUT.
type Task interface {
	Process()
	Print()
}
