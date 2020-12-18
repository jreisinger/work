Work is a scalable work system using goroutines, channels and interfaces. A
common use case for Go is to take a stream of jobs of work and perform them
concurrently, automatically scaling up and down as work becomes available.

Adapted from the
[video](https://learning.oreilly.com/videos/intermediate-go-programming/9781491944073/9781491944073-video234754)
by John Graham-Cumming: Intermediate Go programming - Building a scalable
work system.

Usage:

```
go build
./work < urls.txt
```