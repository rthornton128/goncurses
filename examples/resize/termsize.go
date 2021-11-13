// Inspired by https://stackoverflow.com/a/15784352/587091

package main

import (
	"syscall"
	"unsafe"
)

type winsize struct {
	Row     uint16
	Col     uint16
	XOffset uint16
	YOffset uint16
}

func osTermSize() (int, int, error) {
	w := &winsize{}
	// See http://www.delorie.com/djgpp/doc/libc/libc_495.html
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(w)),
	)
	if int(res) == -1 {
		return 0, 0, err
	}

	return int(w.Row), int(w.Col), nil
}
