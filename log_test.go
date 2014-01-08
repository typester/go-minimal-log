package log

import (
	"bytes"
	l "log"
	"testing"
)

func TestOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf, "", l.Lshortfile)

	Debug("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:13: [debug] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Debugf("foo %s", "bar")
	if string(buf.Bytes()) != "log_test.go:20: [debug] foo bar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Info("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:27: [info] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Infof("foo %s", "bar")
	if string(buf.Bytes()) != "log_test.go:34: [info] foo bar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Warn("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:41: [warn] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Warnf("foo %s", "bar")
	if string(buf.Bytes()) != "log_test.go:48: [warn] foo bar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Crit("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:55: [critical] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Critf("foo %s", "bar")
	if string(buf.Bytes()) != "log_test.go:62: [critical] foo bar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Error("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:69: [error] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	Errorf("foo %s", "bar")
	if string(buf.Bytes()) != "log_test.go:76: [error] foo bar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()
}

func TestLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf, "", l.Lshortfile)

	Debug("foo", "bar")
	if string(buf.Bytes()) != "log_test.go:88: [debug] foobar\n" {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	LogLevel = ERROR
	Debug("foo", "bar")
	Info("foo")
	Warn("foo")
	if buf.Len() > 0 {
		print(string(buf.Bytes()))
		t.Fail()
	}

	Error("foo")
	if buf.Len() == 0 {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

	LogLevel = WARN
	Error("foo")
	if buf.Len() == 0 {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()
	Warn("foo")
	if buf.Len() == 0 {
		print(string(buf.Bytes()))
		t.Fail()
	}
	buf.Reset()

}
