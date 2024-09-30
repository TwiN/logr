package logr

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	logger := New("what", false, nil)
	if logger.threshold != LevelInfo {
		t.Error("expected invalid threshold to be silently set to INFO, got", logger.threshold)
	}
}

func TestLogger(t *testing.T) {
	var writer bytes.Buffer
	logger := New(LevelWarn, false, &writer)
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
	output := writer.String()
	if numberOfNewLines := strings.Count(output, "\n"); numberOfNewLines != 2 {
		t.Error("expected 2 newlines, got", numberOfNewLines)
	}
	if strings.Contains(output, "this is debug") {
		t.Error("expected no debug message, got", output)
	}
	if strings.Contains(output, "this is info") {
		t.Error("expected no info message, got", output)
	}
	if !strings.Contains(output, "this is warn") {
		t.Error("expected warn message, got", output)
	}
	if !strings.Contains(output, "this is error") {
		t.Error("expected error message, got", output)
	}
}

func TestLogger_SetThreshold(t *testing.T) {
	var writer bytes.Buffer
	logger := New(LevelDebug, false, &writer)
	logger.SetThreshold(LevelError)
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
	output := writer.String()
	if numberOfNewLines := strings.Count(output, "\n"); numberOfNewLines != 1 {
		t.Error("expected 1 newline, got", numberOfNewLines)
	}
}

func TestLoggerFormatting(t *testing.T) {
	var writer bytes.Buffer
	logger := New(LevelDebug, true, &writer)
	logger.Debugf("hello, %s", "world")
	logger.Infof("%d", 11111)
	logger.Warnf("%s got %d%% in math", "John Doe", 87)
	logger.Errorf("%s sent %s$%.02f to %s", "John", "CAD", 69.1, "Jane")
	output := writer.String()
	if !strings.Contains(output, "- DEBUG - hello, world") {
		t.Error("expected '- DEBUG - hello, world.', got", output)
	}
	if !strings.Contains(output, "- INFO - 11111") {
		t.Error("expected '- INFO - John Doe got 87% in math', got", output)
	}
	if !strings.Contains(output, "- WARN - John Doe got 87% in math") {
		t.Error("expected '- WARN - John Doe got 87% in math', got", output)
	}
	if !strings.Contains(output, "- ERROR - John sent CAD$69.10 to Jane") {
		t.Error("expected '- ERROR - John sent USD$123.40 to Jane', got", output)
	}
}
