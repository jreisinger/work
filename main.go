package main

import (
	"flag"

	"github.com/jreisinger/work/factory"
)

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	b := &factory.HTTPBoss{}
	factory.Run(b, *n)
}
