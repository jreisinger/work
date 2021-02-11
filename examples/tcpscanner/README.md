Tcpscanner scans TCP ports on a host and reports the open ones.

```bash
$ go build

$ for i in $(seq 1 1024); do echo "$i"; done | ./tcpscanner scanme.nmap.org
scanme.nmap.org:22
scanme.nmap.org:80
```