Domainguesser tries to find subdomains for a given domain. If a subdomain is
found its IP address is looked up. Stolen from "Black Hat Go: 5 Exploiting
DNS".

```
go build
./domainguesser -d example.com < wordlist.txt
```
