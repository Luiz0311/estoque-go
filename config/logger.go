package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writter := io.Writer(os.Stdout)
	logger := log.New(writter, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writter, "DEBUG: ", logger.Flags()),
		info:    log.New(writter, "INFO: ", logger.Flags()),
		warning: log.New(writter, "WARNING: ", logger.Flags()),
		err:     log.New(writter, "ERROR: ", logger.Flags()),
		writer:  writter,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}
