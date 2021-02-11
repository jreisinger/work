Urlchecker checks whether URL returns 200.

```bash
$ go build

$ ./urlchecker -w 10 < urls.txt
OK    https://reisinge.net/notes/go/basics
NOTOK https://nonexistent.net
OK    https://golang.org/doc
OK    https://perl.org
NOTOK https://perl.org/python
```
