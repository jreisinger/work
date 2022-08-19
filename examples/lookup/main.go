// Lookup prints domains that use Cloudflare's nameserver (NS).
package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/jreisinger/work"
)

type lookupFactory struct{}

func (lookupFactory) Generate(line string) work.Task {
	return &lookup{name: line}
}

type lookup struct {
	name       string
	err        error
	cloudflare bool
}

func (l *lookup) Process() {
	nss, err := net.LookupNS(l.name)
	if err != nil {
		l.err = err
	} else {
		for _, ns := range nss {
			if strings.HasSuffix(ns.Host, ".ns.cloudflare.com.") {
				l.cloudflare = true
				break
			}
		}
	}
}

func (l *lookup) Print() {
	if l.cloudflare {
		fmt.Println(l.name)
	}
}

func main() {
	work.Run(&lookupFactory{}, 100, []string{})
}
