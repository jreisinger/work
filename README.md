package work // import "github.com/jreisinger/work"

Package work concurrently generates and processes tasks. The tasks are
generated from lines supplied on STDIN. The results of tasks processing are
then printed on STDOUT. To use it you just need to implement Factory and
Task interfaces.

func Run(f Factory, n int)
type Factory interface{ ... }
type Task interface{ ... }

For example:

```go
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jreisinger/work"
)

type HTTPFactory struct{}

func (f *HTTPFactory) Generate(line string) work.Task {
	h := &HTTPTask{}
	h.URL = line
	return h
}

type HTTPTask struct {
	URL string
	OK  bool
}

func (t *HTTPTask) Process() {
	resp, err := http.Get(t.URL)
	if err != nil {
		t.OK = false
		return
	}
	if resp.StatusCode == http.StatusOK {
		t.OK = true
	}
}

func (t *HTTPTask) Print() {
	status := map[bool]string{
		true:  "OK",
		false: "FAIL",
	}
	fmt.Printf("%-60s %s\n", t.URL, status[t.OK])
}

func main() {
	n := flag.Int("n", 100, "number of concurrent workers")
	flag.Parse()

	f := &HTTPFactory{}
	work.Run(f, *n)
}
```

```bash
go get -u github.com/jreisinger/work
go run main.go < urls.txt
```

Adapted from John Graham-Cumming's [talk](https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014).