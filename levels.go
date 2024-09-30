package logr

import (
	"errors"
	"strings"
)

type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

var (
	ErrInvalidLevelString = errors.New("invalid level, must be one of DEBUG, INFO, WARN or ERROR")
)

func (level Level) Value() int {
	switch level {
	case LevelDebug:
		return 0
	case LevelInfo:
		return 1
	case LevelWarn:
		return 2
	case LevelError:
		return 3
	default:
		return -1
	}
}

func LevelFromString(level string) (Level, error) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return LevelDebug, nil
	case "INFO":
		return LevelInfo, nil
	case "WARN":
		return LevelWarn, nil
	case "ERROR":
		return LevelError, nil
	default:
		return LevelDebug, ErrInvalidLevelString
	}
}
