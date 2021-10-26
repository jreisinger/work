Package work is useful for building CLI tools that need to run many tasks
quickly. It concurrently generates and processes tasks. The tasks are
generated from lines read from file(s) or STDIN and load balanced among
workers for processing. After each task is processed a result is printed on
STDOUT.

To use it you just need to implement
[Factory](https://pkg.go.dev/github.com/jreisinger/work#Factory) and
[Task](https://pkg.go.dev/github.com/jreisinger/work#Task) interfaces. Then
Run the Factory. See [examples](examples) folder for sample implementations.

Adapted from John Graham-Cumming's [talk](https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014).