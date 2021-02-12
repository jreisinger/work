package main

import (
	"fmt"
	"log"

	"github.com/jreisinger/work"
)

type factory struct{}

func (f *factory) Generate(line string) work.Task {
	t := &task{filename: line}
	return t
}

type task struct {
	filename string
}

func (t *task) Process() {
	err := compress(t.filename)
	if err != nil {
		log.Printf("compressing %s: %v\n", t.filename, err)
	}
}

func (t *task) Print() {
	fmt.Printf("created %s\n", t.filename+".gz")
}

func main() {
	f := &factory{}
	work.Run(f, 100)
}
