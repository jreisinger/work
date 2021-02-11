package main

import (
	"flag"
	"fmt"

	"github.com/jreisinger/work"
)

type factory struct {
	domain     string
	dnsSrvAddr string
}

func (f *factory) Generate(line string) work.Task {
	subdomain := line
	t := &task{
		fqdn:          fmt.Sprintf("%s.%s", subdomain, f.domain),
		dnsServerAddr: f.dnsSrvAddr,
	}
	return t
}

type task struct {
	dnsServerAddr string
	fqdn          string
	ipAddrs       []string
}

func (t *task) Process() {
	results := lookup(t.fqdn, t.dnsServerAddr)
	t.ipAddrs = results
}

func (t *task) Print() {
	for _, ip := range t.ipAddrs {
		fmt.Printf("%s\t%s\n", t.fqdn, ip)
	}
}

func main() {
	w := flag.Int("w", 100, "number of concurrent workers")
	d := flag.String("d", "example.com", "domain to guess")
	s := flag.String("s", "8.8.8.8:53", "DNS server address")
	flag.Parse()

	f := &factory{
		domain:     *d,
		dnsSrvAddr: *s,
	}
	work.Run(f, *w)
}
