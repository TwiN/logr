# logr
Go logging library with levels.

## Usage
```go
logr.Debug("This is a debug message")
logr.Infof("This is an %s message", "info")
logr.Warn("This is a warn message")
logr.Error("This is an error message")
```

You can set the default logger's threshold like so:
```go
logr.SetThreshold(logr.WarnLevel)
```
The above would make it so only `Warn` and `Error` messages are logged, while `Debug` and `Info` messages are ignored.

TODO: Finish documentation