package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/jreisinger/work"
)

type factory struct {
	host string
}

func (f *factory) Generate(line string) work.Task {
	t := &task{
		host: f.host,
		port: line,
	}
	return t
}

type task struct {
	host string
	port string
	open bool
}

func (t *task) Process() {
	addr := fmt.Sprintf("%s:%s", t.host, t.port)
	conn, err := net.Dial("tcp", addr)
	if err != nil { // can't connect
		return
	}
	conn.Close()
	t.open = true
}

func (t *task) Print() {
	if t.open {
		fmt.Printf("%s:%s\n", t.host, t.port)
	}
}

func main() {
	w := flag.Int("w", 100, "number of concurrent workers")
	flag.Parse()

	host := flag.Args()[0]
	f := &factory{host: host}
	work.Run(f, *w, []string{})
}
