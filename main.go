package main

import (
	"flag"

	"github.com/jreisinger/work/factory"
	"github.com/jreisinger/work/tasks"
)

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	b := &tasks.HTTPBoss{}
	factory.Run(b, *n)
}
