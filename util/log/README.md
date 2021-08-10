# Log

DEPRECATED: use github.com/crypto-zero/go-micro/v2/logger interface

This is the global logger for all micro based libraries.

## Set Logger

Set the logger for micro libraries

```go
// import go-micro/util/log
import "github.com/crypto-zero/go-micro/util/log"

// SetLogger expects github.com/crypto-zero/go-micro/debug/log.Log interface
log.SetLogger(mylogger)
```
