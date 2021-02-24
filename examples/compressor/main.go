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
	ok       bool
}

func (t *task) Process() {
	err := compress(t.filename)
	if err != nil {
		log.Printf("compressing %s: %v\n", t.filename, err)
		return
	}
	t.ok = true
}

func (t *task) Print() {
	if t.ok {
		fmt.Printf("created %s\n", t.filename+".gz")
	}
}

func main() {
	f := &factory{}
	work.Run(f, 100, []string{})
}
