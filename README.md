Package work generates tasks from lines of STDIN, processes them concurrently
and prints to STDOUT. To use it you just need to implement Generator and Task
interfaces.

For example:

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
	}
}

func (h *HTTPTask) Print() {
	status := map[bool]string{
		true:  "OK",
		false: "FAIL",
	}
	fmt.Printf("%-60s %s\n", h.URL, status[h.OK])
}

func main() {
	n := flag.Int("n", 100, "number of workers")
	flag.Parse()

	g := &HTTPGenerator{}
	work.Do(g, *n)
}
```

```bash
go get -u github.com/jreisinger/work
go run main.go < urls.txt
```

Adapted from John Graham-Cumming's [talk](https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014).