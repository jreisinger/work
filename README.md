```
package work // import "github.com/jreisinger/work"

Package work concurrently generates and processes tasks. The tasks are
generated from lines supplied on STDIN. The results of tasks processing are
then printed on STDOUT. To use it you just need to implement Factory and
Task interfaces.

func Run(f Factory, n int)
type Factory interface{ ... }
type Task interface{ ... }
```

See `examples` folder sample implementations.

```bash
cd examples
go run urlchecker.go < urls.txt
```

Adapted from John Graham-Cumming's [talk](https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014).