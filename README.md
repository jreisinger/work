Work is a scalable work system that can generate and work on many tasks
*concurrently*. To use it you need to implement `Generator` and `Task`
interfaces, for example:

```go
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jreisinger/work"
)

type HTTPGenerator struct{}

func (b *HTTPGenerator) Generate(line string) work.Task {
	h := &HTTPTask{}
	h.URL = line
	return h
}

type HTTPTask struct {
	URL string
	OK  bool
}

func (h *HTTPTask) Process() {
	resp, err := http.Get(h.URL)
	if err != nil {
		h.OK = false
		return
	}
	if resp.StatusCode == http.StatusOK {
		h.OK = true
		return
	}
}

func (h *HTTPTask) Print() {
	fmt.Printf("%s %t\n", h.URL, h.OK)
}

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	g := &HTTPGenerator{}
	work.Do(g, *n)
}
```

```bash
go run main.go < urls.txt
```

Adapted from John Graham-Cumming's [talk](https://github.com/jgrahamc/dotgo).