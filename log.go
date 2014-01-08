package log

import (
	"code.google.com/p/go.crypto/ssh/terminal"
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	l "log"
	"os"
)

const (
	MUTE = iota
	DEBUG
	INFO
	WARN
	CRITICAL
	ERROR = 99
)

// configurations
var LogLevel = DEBUG
var Logger = l.New(defaultOutput, "", l.LstdFlags|l.Lshortfile)

var Colors = map[int]string{
	DEBUG:    "",
	INFO:     "green",
	WARN:     "yellow",
	CRITICAL: "red+b",
	ERROR:    "red",
}

var Tags = map[int]string{
	DEBUG:    "debug",
	INFO:     "info",
	WARN:     "warn",
	CRITICAL: "critical",
	ERROR:    "error",
}

type Fder interface {
	Fd() uintptr
}

var defaultOutput io.Writer = os.Stderr

func SetOutput(out io.Writer, prefix string, flag int) {
	defaultOutput = out
	Logger = l.New(out, prefix, flag)
}

func Debugf(format string, args ...interface{}) {
	if levelOK(DEBUG) {
		logf(DEBUG, format, args...)
	}
}

func Debug(args ...interface{}) {
	if levelOK(DEBUG) {
		log(DEBUG, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if levelOK(INFO) {
		logf(INFO, format, args...)
	}
}

func Info(args ...interface{}) {
	if levelOK(INFO) {
		log(INFO, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if levelOK(WARN) {
		logf(WARN, format, args...)
	}
}

func Warn(args ...interface{}) {
	if levelOK(WARN) {
		log(WARN, args...)
	}
}

func Critf(format string, args ...interface{}) {
	if levelOK(CRITICAL) {
		logf(CRITICAL, format, args...)
	}
}

func Crit(args ...interface{}) {
	if levelOK(CRITICAL) {
		log(CRITICAL, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if levelOK(ERROR) {
		logf(ERROR, format, args...)
	}
}

func Error(args ...interface{}) {
	if levelOK(ERROR) {
		log(ERROR, args...)
	}
}

func levelOK(level int) bool {
	if level == MUTE {
		return false
	}
	return LogLevel <= level
}

func logf(level int, format string, args ...interface{}) {
	format = "[" + Tags[level] + "] " + format
	output(level, fmt.Sprintf(format, args...))
}

func log(level int, args ...interface{}) {
	_args := []interface{}{"[" + Tags[level] + "] "}
	args = append(_args, args...)
	output(level, fmt.Sprint(args...))
}

func output(level int, v string) {
	if file, ok := defaultOutput.(Fder); ok && terminal.IsTerminal(int(file.Fd())) {
		v = ansi.Color(v, Colors[level])
	}
	Logger.Output(4, v)
}
