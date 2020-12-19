Work is a scalable work system using Go goroutines, channels and interfaces. It
can generate and work on many tasks concurrently. To use it you need to
implement Generator and Task. See `example`.

Adapted from John Graham-Cumming's [talk](https://github.com/jgrahamc/dotgo).

Usage:

```go
package main

import (
	"flag"

	"github.com/jreisinger/work"
	"github.com/jreisinger/work/example"
)

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	g := &example.HTTTGenerator{}
	work.Do(g, *n)
}
```

```bash
go run main.go < urls.txt
```