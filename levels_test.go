package logr

import (
	"testing"
)

func TestLevelFromString(t *testing.T) {
	scenarios := []struct {
		input       string
		expected    Level
		expectedErr error
	}{
		{"DEBUG", LevelDebug, nil},
		{"debug", LevelDebug, nil},
		{"INFO", LevelInfo, nil},
		{"info", LevelInfo, nil},
		{"WARN", LevelWarn, nil},
		{"warn", LevelWarn, nil},
		{"ERROR", LevelError, nil},
		{"error", LevelError, nil},
		{"", 0, ErrInvalidLevelString},
		{"invalid", 0, ErrInvalidLevelString},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.input, func(t *testing.T) {
			actual, err := LevelFromString(scenario.input)
			if actual != scenario.expected {
				t.Errorf("expected %d, got %d", scenario.expected, actual)
			}
			if err != scenario.expectedErr {
				t.Errorf("expected %v, got %v", scenario.expectedErr, err)
			}
		})
	}
}
