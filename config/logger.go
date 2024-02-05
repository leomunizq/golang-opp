package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}

// The NewLogger function creates a new logger with different log levels and a specified prefix.
func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	
	return &Logger{
		debug: log.New(writer, "DEBUG: ", logger.Flags()),
		info: log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err: log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

// The `Debug` method of the `Logger` struct is used to log debug messages. It takes a variadic
// parameter `value` of type `interface{}` which allows you to pass in any number of values. The method
// then calls the `Println` method of the `debug` logger, which prints the values to the log output
// with the prefix "DEBUG:".
func (l *Logger) Debug(value ...interface{}) {
	l.debug.Println(value...)
}

func (l *Logger) Info(value ...interface{}) {
	l.info.Println(value...)
}

func (l *Logger) Warning(value ...interface{}) {
	l.warning.Println(value...)
}

func (l *Logger) Err(value ...interface{}) {
	l.err.Println(value...)
}

// The `Debugf` method of the `Logger` struct is used to log debug messages with a formatted string. It
// takes a `format` string as the first parameter, which specifies the format of the log message. The
// method also takes a variadic parameter `value` of type `interface{}`, which allows you to pass in
// any number of values to be formatted into the log message.
func (l *Logger) Debugf(format string, value ...interface{}) {
	l.debug.Printf(format, value...)
}

func (l *Logger) Infof(format string, value ...interface{}) {
	l.info.Printf(format, value...)
}

func (l *Logger) Warningf(format string, value ...interface{}) {
	l.warning.Printf(format, value...)
}

func (l *Logger) Errf(format string, value ...interface{}) {
	l.err.Printf(format, value...)
}


