package work // import "github.com/jreisinger/work"

Package work is useful for building CLI tools that need to run many tasks
quickly. It concurrently generates and processes tasks. The tasks are
generated from lines supplied on STDIN and load balanced among workers for
processing. After each task is processed a result is printed on STDOUT. To
use it you just need to implement Factory and Task interfaces. See examples
folder for sample implementations. Adapted from John Graham-Cumming's talk:
https://github.com/cloudflare/jgc-talks/tree/master/dotGo/2014

func Run(f Factory, w int)
type Factory interface{ ... }
type Task interface{ ... }
