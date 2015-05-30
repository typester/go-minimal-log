package log

import (
	"fmt"
	"io"
	l "log"
	"os"

	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/go-isatty"
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

type color struct {
	value  ct.Color
	bright bool
}

var Colors = map[int]color{
	DEBUG:    {ct.White, false},
	INFO:     {ct.Green, false},
	WARN:     {ct.Yellow, false},
	CRITICAL: {ct.Red, true},
	ERROR:    {ct.Red, false},
}

var Tags = map[int]string{
	DEBUG:    "debug",
	INFO:     "info",
	WARN:     "warn",
	CRITICAL: "critical",
	ERROR:    "error",
}

type fder interface {
	Fd() uintptr
}

var defaultOutput io.Writer = os.Stderr

func SetFlags(flags int) {
	Logger.SetFlags(flags)
}

func SetPrefix(prefix string) {
	Logger.SetPrefix(prefix)
}

func SetOutput(out io.Writer) {
	defaultOutput = out
	Logger = l.New(out, Logger.Prefix(), Logger.Flags())
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

func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
	os.Exit(1)
}

func Fatal(args ...interface{}) {
	Error(args...)
	os.Exit(1)
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
	if file, ok := defaultOutput.(fder); ok && isatty.IsTerminal(uintptr(file.Fd())) {
		c := Colors[level]
		ct.ChangeColor(c.value, c.bright, ct.None, false)
		Logger.Output(4, v)
		ct.ResetColor()
	} else {
		Logger.Output(4, v)
	}

}
