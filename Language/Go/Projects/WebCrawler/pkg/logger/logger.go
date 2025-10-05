package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Level represents the severity level of a log message.
type Level int

// Log levels in increasing order of severity.
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// levelNames maps Level constants to their string representations.
var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// Logger provides structured logging with configurable levels,
// output destinations, and time formatting.
type Logger struct {
	level      Level
	output     io.Writer
	mu         sync.Mutex
	timeFormat string
}

var (
	defaultLogger *Logger
	once          sync.Once
)

// Init initializes the default logger.
//
// Args:
//
//	level (string): The minimum log level. Accepted values: "debug", "info", "warn", "error", "fatal".
//	outputPath (string): The file path to output logs. If empty, logs are written only to stdout.
//
// Returns:
//
//	error: An error if file or directory creation fails, otherwise nil.
//
// Example:
//
//	err := logger.Init("info", "./logs/app.log")
//	if err != nil {
//	    fmt.Println("Failed to initialize logger:", err)
//	}
func Init(level string, outputPath string) error {
	var err error

	once.Do(func() {
		var l Level
		switch level {
		case "debug":
			l = DEBUG
		case "info":
			l = INFO
		case "warn":
			l = WARN
		case "error":
			l = ERROR
		case "fatal":
			l = FATAL
		default:
			l = INFO
		}

		var output io.Writer = os.Stdout
		if outputPath != "" {
			dir := filepath.Dir(outputPath)
			if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
				err = mkErr
				return
			}

			file, fileErr := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if fileErr != nil {
				err = fileErr
				return
			}

			output = io.MultiWriter(os.Stdout, file)
		}

		defaultLogger = &Logger{
			level:      l,
			output:     output,
			timeFormat: time.RFC3339,
		}
	})

	return err
}

// GetLogger returns the default logger instance.
//
// Returns:
//
//	*Logger: The singleton logger instance.
//	If Init() has not been called, it returns a default logger writing to stdout.
func GetLogger() *Logger {
	if defaultLogger == nil {
		defaultLogger = &Logger{
			level:      INFO,
			output:     os.Stdout,
			timeFormat: time.RFC3339,
		}
	}
	return defaultLogger
}

// log writes a formatted message if the log level is above the configured threshold.
//
// Args:
//
//	level (Level): The severity level of the log message.
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(l.timeFormat)
	levelName := levelNames[level]
	message := fmt.Sprintf(format, v...)
	logMsg := fmt.Sprintf("%s [%s] %s\n", timestamp, levelName, message)

	_, _ = l.output.Write([]byte(logMsg))

	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a debug-level message.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

// Info logs an info-level message.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

// Warn logs a warning-level message.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

// Error logs an error-level message.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

// Fatal logs a fatal-level message and terminates the program.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
}

// The following are convenience wrapper functions that call
// the corresponding methods on the default logger.

// Debug logs a debug-level message using the default logger.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func Debug(format string, v ...interface{}) {
	GetLogger().Debug(format, v...)
}

// Info logs an info-level message using the default logger.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func Info(format string, v ...interface{}) {
	GetLogger().Info(format, v...)
}

// Warn logs a warning-level message using the default logger.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func Warn(format string, v ...interface{}) {
	GetLogger().Warn(format, v...)
}

// Error logs an error-level message using the default logger.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func Error(format string, v ...interface{}) {
	GetLogger().Error(format, v...)
}

// Fatal logs a fatal-level message using the default logger
// and terminates the program.
//
// Args:
//
//	format (string): The message format string.
//	v (...interface{}): The arguments for the format string.
func Fatal(format string, v ...interface{}) {
	GetLogger().Fatal(format, v...)
}
