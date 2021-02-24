// Package work is useful for building CLI tools that need to run many tasks
// quickly. It concurrently generates and processes tasks. The tasks are
// generated from lines read from file(s) or STDIN and load balanced among
// workers for processing. After each task is processed a result is printed on
// STDOUT.
//
// To use it you just need to implement Factory and Task interfaces. See
// examples folder for sample implementations. Adapted from John
// Graham-Cumming's talk:
// https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014
package work

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// Factory generates tasks from lines read from file(s) or STDIN.
type Factory interface {
	Generate(line string) Task
}

// Task is anything that can be processed. The result of the processing is then
// printed on STDOUT.
type Task interface {
	Process()
	Print()
}

// Run concurrently runs Factory and nWorkers. Factory generates tasks from
// lines read from filenames or from STDIN if filenames is empty. The tasks are
// load balanced among workers that process them. When all tasks are processed
// the results are printed on STDOUT.
func Run(f Factory, nWorkers int, filenames []string) {
	var wg sync.WaitGroup
	in := make(chan Task)
	out := make(chan Task)

	// Generate tasks that will be processed.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		if len(filenames) == 0 {
			if err := genTasks(os.Stdin, f, in); err != nil {
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
				if err := genTasks(file, f, in); err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}()

	// Create workers to process the tasks.
	for i := 0; i < nWorkers; i++ {
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

// getTasks generate Tasks from file and send the to in channel. It uses factory
// to generate them.
func genTasks(file *os.File, factory Factory, in chan Task) error {
	s := bufio.NewScanner(file)
	for s.Scan() {
		in <- factory.Generate(s.Text())
	}
	return s.Err()
}
