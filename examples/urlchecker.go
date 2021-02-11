package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jreisinger/work"
)

type factory struct{}

func (f *factory) Generate(line string) work.Task {
	return &task{URL: line}
}

type task struct {
	URL    string
	Status bool
}

func (t *task) Process() {
	resp, err := http.Get(t.URL)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		t.Status = true
	}
}

func (t *task) Print() {
	status := map[bool]string{
		true:  "OK",
		false: "NOTOK",
	}
	fmt.Printf("%-5s %s\n", status[t.Status], t.URL)
}

func main() {
	w := flag.Int("w", 100, "number of concurrent workers")
	flag.Parse()

	f := &factory{}
	work.Run(f, *w)
}
