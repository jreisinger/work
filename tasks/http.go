// Package tasks holds some concrete tasks that can be run in factory.
package tasks

import (
	"fmt"
	"net/http"

	"github.com/jreisinger/work/factory"
)

// HTTPTask represents a URL and whether it's OK.
type HTTPTask struct {
	URL string
	OK  bool
}

// Process tries to get a URL and sets OK to true if it returns 200.
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

// Output prints URL and whether it's OK.
func (h *HTTPTask) Output() {
	fmt.Printf("%s %t\n", h.URL, h.OK)
}

// HTTPBoss represents HTTP tasks generator.
type HTTPBoss struct{}

// Create creates HTTP tasks.
func (b *HTTPBoss) Create(line string) factory.Task {
	h := &HTTPTask{}
	h.URL = line
	return h
}
