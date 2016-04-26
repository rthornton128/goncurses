// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: ncurses
// #cgo windows CFLAGS: -DNCURSES_MOUSE_VERSION
// #cgo windows LDFLAGS: -lpdcurses
// #include <curses.h>
// #include "goncurses.h"
import "C"

import (
	"unsafe"
)

type MouseEvent struct {
	Id      int16       /* device ID */
	X, Y, Z int         /* event coordinates */
	State   MouseButton /* button state */
}

// GetMouse returns the MouseEvent associated with a KEY_MOUSE event returned
// by a call to GetChar(). Returns a new MouseEvent or nil on error or if no
// event is currently in the mouse event queue
func GetMouse() *MouseEvent {
	var event C.MEVENT
	if C.ncurses_getmouse(&event) != C.OK {
		return nil
	}
	return &MouseEvent{
		Id:    int16(event.id),
		Y:     int(event.y),
		X:     int(event.x),
		Z:     int(event.z),
		State: MouseButton(event.bstate),
	}
}

// MouseOk returns true if ncurses has built-in mouse support. On ncurses 5.7
// and earlier, this function is not present and so will always return false
func MouseOk() bool {
	return bool(C.ncurses_has_mouse())
}

// MouseInterval sets the maximum time in milliseconds that can elapse
// between press and release mouse events and returns the previous setting.
// Use a value of 0 (zero) to disable click resolution. Use a value of -1
// to get the previous value without changing the current value. Default
// value is 1/6 of a second.
func MouseInterval(ms int) int {
	return int(C.mouseinterval(C.int(ms)))
}

// MouseMask accepts a single int of OR'd mouse events. If a mouse event
// is triggered, GetChar() will return KEY_MOUSE. To retrieve the actual
// event use GetMouse() to pop it off the queue. Pass a pointer as the
// second argument to store the prior events being monitored or nil.
func MouseMask(mask MouseButton, old *MouseButton) MouseButton {
	return MouseButton(C.mousemask((C.mmask_t)(mask),
		(*C.mmask_t)(unsafe.Pointer(old))))
}
