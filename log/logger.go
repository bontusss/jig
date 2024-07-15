package log
// This file defines a custom logger package with different log levels and handlers for logging messages to the console and files.
// It provides functions to create a new file handler, handle logging to console and file, and create a new logger instance with specified log level and handlers.
// Usage examples are provided for each function.

import (
	"context"
	"log/slog"
	"os"
)

type LogLevel slog.Level
type ConsoleHandler struct{}
type FileHandler struct {
    file *os.File
}
type Logger struct {
    level    LogLevel
    handlers []Handler
}

const (
	DEBUG LogLevel = LogLevel(slog.LevelDebug)
	INFO  LogLevel = LogLevel(slog.LevelInfo)
	WARN  LogLevel = LogLevel(slog.LevelWarn)
	ERROR LogLevel = LogLevel(slog.LevelError)
)

type Handler interface {
	Handle(level LogLevel, msg string, kv ...interface{})
}

// Handle function logs messages with the specified log level and key-value pairs to a file.
// This Handle function logs data to the console.
// Usage: fileHandler.Handle(logger.DEBUG, "This is a debug message", "key1", "value1", "key2", "value2")
func (c *ConsoleHandler) Handle(level LogLevel, msg string, kv ...interface{}) {
	slog.New(slog.NewTextHandler(os.Stdout, nil)).Log(context.Background(), slog.Level(level), msg, kv...)
}

// NewFileHandler creates a new file handler that logs messages to a specified file path.
//
// Usage: 
//        fileHandler, err := NewFileHandler("path/to/logfile.log")
//        if err != nil {
//            // handle error
//        }
//        defer fileHandler.file.Close()
//        fileHandler.Handle(logger.DEBUG, "This is a debug message", "key1", "value1", "key2", "value2")
func NewFileHandler(filePath string) (*FileHandler, error) {
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    return &FileHandler{file: file}, nil
}

// Handle function logs messages with the specified log level and key-value pairs to a file.
// This Handle function logs data to a file.
// Usage: fileHandler.Handle(logger.DEBUG, "This is a debug message", "key1", "value1", "key2", "value2")
func (h *FileHandler) Handle(level LogLevel, msg string, keysAndValues ...interface{}) {
    slog.New(slog.NewJSONHandler(h.file, nil)).Log(context.Background(), slog.Level(level), msg, keysAndValues...)
}

// NewLogger creates a new Logger instance with the specified log level and handlers.
// Usage: loggerInstance := NewLogger(logger.DEBUG, []Handler{&ConsoleHandler{}, fileHandler})
func NewLogger(level LogLevel, handlers []Handler) *Logger {
    return &Logger{level: level, handlers: handlers}
}

// log function logs messages with the specified log level and key-value pairs.
// Usage: loggerInstance.log(logger.DEBUG, "This is a debug message", "key1", "value1", "key2", "value2")
func (l *Logger) log(level LogLevel, msg string, keysAndValues ...interface{}) {
    if level < l.level {
        return
    }
    for _, handler := range l.handlers {
        handler.Handle(level, msg, keysAndValues...)
    }
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
    l.log(DEBUG, msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
    l.log(INFO, msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
    l.log(WARN, msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
    l.log(ERROR, msg, keysAndValues...)
}