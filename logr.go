package logr

import (
	"io"
	"log"
	"os"
)

var defaultLogger = New(LevelInfo, false, os.Stdout)

type Logger struct {
	// Threshold is the minimum level output by this logger.
	//
	// For example, if the Threshold is set to LevelWarn, then only logs of LevelWarn and LevelError will be output
	// while logs of LevelDebug and LevelInfo will be ignored.
	//
	// Defaults to LevelInfo.
	Threshold Level

	// ShouldPrefixMessageWithLevel is whether to include the log level prefix in each log.
	//
	// For example, if this is set to true, then a debug log would output "<date> DEBUG <message>"
	//
	// Defaults to false.
	ShouldPrefixMessageWithLevel bool

	stdLogger *log.Logger
}

// New creates a new logger with the given threshold and output.
func New(threshold Level, shouldPrefixMessageWithLevel bool, output io.Writer) *Logger {
	return &Logger{
		Threshold:                    threshold,
		ShouldPrefixMessageWithLevel: shouldPrefixMessageWithLevel,
		stdLogger:                    log.New(output, "", log.LstdFlags),
	}
}

func (logger *Logger) SetOutput(output io.Writer) {
	// Because we're using the standard log package under the hood, we can just directly pass this
	// to the logger and call it a day. It already takes care of locking and discarding the previous writer.
	logger.stdLogger.SetOutput(output)
}

func (logger *Logger) SetThreshold(threshold Level) {
	logger.Threshold = threshold
}

func (logger *Logger) Log(level Level, message string) {
	logger.Logf(level, message)
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	if level < logger.Threshold {
		// The log level is below the threshold, so ignore the log
		return
	}
	if logger.ShouldPrefixMessageWithLevel {
		switch level {
		case LevelDebug:
			format = "- DEBUG - " + format
		case LevelInfo:
			format = "- INFO - " + format
		case LevelWarn:
			format = "- WARN - " + format
		case LevelError:
			format = "- ERROR - " + format
		}
	}
	logger.stdLogger.Printf(format, args...)
}

func (logger *Logger) Debug(message string) {
	logger.Log(LevelDebug, message)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logf(LevelDebug, format, args...)
}

func (logger *Logger) Info(message string) {
	logger.Log(LevelInfo, message)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logf(LevelInfo, format, args...)
}

func (logger *Logger) Warn(message string) {
	logger.Log(LevelWarn, message)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logf(LevelWarn, format, args...)
}

func (logger *Logger) Error(message string) {
	logger.Log(LevelError, message)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logf(LevelError, format, args...)
}

func Debug(message string) {
	defaultLogger.Log(LevelDebug, message)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Logf(LevelDebug, format, args...)
}

func Info(message string) {
	defaultLogger.Log(LevelInfo, message)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Logf(LevelInfo, format, args...)
}

func Warn(message string) {
	defaultLogger.Log(LevelWarn, message)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Logf(LevelWarn, format, args...)
}

func Error(message string) {
	defaultLogger.Log(LevelError, message)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Logf(LevelError, format, args...)
}

// SetOutput sets the output of the default logger
func SetOutput(output io.Writer) {
	defaultLogger.SetOutput(output)
}

// SetThreshold sets the minimum level output by the default logger
func SetThreshold(threshold Level) {
	defaultLogger.SetThreshold(threshold)
}
