package main

import (
	"flag"

	"github.com/jreisinger/work/factory"
	"github.com/jreisinger/work/jobs"
)

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	b := &jobs.HTTPBoss{}
	factory.Run(b, *n)
}
