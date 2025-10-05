package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type Logger struct {
	level      Level
	output     io.Writer
	mu         sync.Mutex
	timeFormat string
}

var defaultLogger *Logger
var once sync.Once

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
		default:
			l = INFO
		}

		var output io.Writer = os.Stdout
		if outputPath != "" {
			dir := filepath.Dir(outputPath)
			if err = os.MkdirAll(dir, 0755); err != nil {
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

func (l *Logger) log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(l.timeFormat)
	levelName := levelNames[level]

	msg := fmt.Sprintf(format, v...)

	logMsg := fmt.Sprintf("%s [%s] %s\n", timestamp, levelName, msg)
	l.output.Write([]byte(logMsg))

	if level == FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func Debug(format string, v ...interface{}) {
	GetLogger().Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	GetLogger().Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	GetLogger().Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	GetLogger().Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	GetLogger().Fatal(format, v...)
}
