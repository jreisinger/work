Urlchecker checks whether URL returns 200.

```bash
$ go build

$ ./urlchecker -w 10 nonexistent.txt urls.txt
2021/02/24 19:49:22 open nonexistent.txt: no such file or directory
NOTOK https://nonexistent.net
OK    https://reisinge.net/notes/go/basics
OK    https://perl.org
NOTOK https://perl.org/python
OK    https://golang.org/doc
```
