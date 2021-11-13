// Inspired by https://stackoverflow.com/a/15784352/587091

package main

// #include <sys/ioctl.h>
import "C"

import (
	"syscall"
	"unsafe"
)

func osTermSize() (int, int, error) {
	w := &C.struct_winsize{}
	// See http://www.delorie.com/djgpp/doc/libc/libc_495.html
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(w)),
	)
	if int(res) == -1 {
		return 0, 0, err
	}

	return int(w.ws_row), int(w.ws_col), nil
}
