package logr

import (
	"errors"
	"strings"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	ErrInvalidLevelString = errors.New("invalid level, must be one of DEBUG, INFO, WARN or ERROR")
)

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
		return 0, ErrInvalidLevelString
	}
}
