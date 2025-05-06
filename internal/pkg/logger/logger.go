package logger

import (
	"fmt"
	"log"
)

type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type Logger struct {
	prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Info(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    log.Printf("[INFO] %s %s", l.prefix, message)
}

func (l *Logger) Error(format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    log.Printf("[ERROR] %s %s", l.prefix, message)
}
