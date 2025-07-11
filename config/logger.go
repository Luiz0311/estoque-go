package config

import (
	"io"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
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
		debug:   log.New(writter, Green+"DEBUG: "+Reset, logger.Flags()),
		info:    log.New(writter, Blue+"INFO: "+Reset, logger.Flags()),
		warning: log.New(writter, Yellow+"WARNING: "+Reset, logger.Flags()),
		err:     log.New(writter, Red+"ERROR: "+Reset, logger.Flags()),
		writer:  writter,
	}
}

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...any) {
	l.warning.Println(v...)
}

func (l *Logger) Err(v ...any) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errf(format string, v ...any) {
	l.err.Printf(format, v...)
}
