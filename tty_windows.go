package log

import (
	"syscall"
	"unsafe"
)

var getConsoleMode = syscall.MustLoadDLL("kernel32.dll").MustFindProc("GetConsoleMode")

func isatty(fd int) bool {
	var st uint32
	r1, _, err := getConsoleMode.Call(uintptr(fd), uintptr(unsafe.Pointer(&st)))
	return r1 != 0 && err != nil
}
