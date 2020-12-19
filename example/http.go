// Package example contains some implementions of Generator and Task interfaces.
package example

import (
	"fmt"
	"net/http"

	"github.com/jreisinger/work"
)

// HTTTGenerator represents HTTP tasks generator.
type HTTTGenerator struct{}

// Generate generates HTTP tasks.
func (b *HTTTGenerator) Generate(line string) work.Task {
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
