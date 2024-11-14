# logr
Go logging library with levels.

## Usage
```console
go get -u github.com/TwiN/logr
```

```go
import "github.com/TwiN/logr"

func main() {
    logr.Debug("This is a debug message")
    logr.Infof("This is an %s message", "info")
    logr.Warn("This is a warn message")
    logr.Error("This is an error message")
    logr.Fatal("This is a fatal message") // Exits with code 1
}
```

You can set the default logger's threshold like so:
```go
logr.SetThreshold(logr.LevelWarn)
```
The above would make it so only `WARN`, `ERROR` and `FATAL` messages are logged, while `DEBUG` and `INFO` messages are ignored.

TODO: Finish documentation