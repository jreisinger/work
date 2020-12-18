// Package factory contains HTTPTask and HTTPBoss. These are concrete
// implementions of Boss and Task interfaces.
package factory

import (
	"fmt"
	"net/http"
)

// HTTPBoss represents HTTP tasks generator.
type HTTPBoss struct{}

// Create creates HTTP tasks.
func (b *HTTPBoss) Create(line string) Task {
	h := &HTTPTask{}
	h.URL = line
	return h
}

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
