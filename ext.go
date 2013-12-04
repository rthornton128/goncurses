// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// #include <curses.h>
// #include <form.h>
// #include <menu.h>
import "C"

import (
	"errors"
	"syscall"
)

// DriverActions is a convenience mapping for common responses
// to keyboard input
var DriverActions = map[Key]int{
	KEY_DOWN:     C.REQ_DOWN_ITEM,
	KEY_HOME:     C.REQ_FIRST_ITEM,
	KEY_END:      C.REQ_LAST_ITEM,
	KEY_LEFT:     C.REQ_LEFT_ITEM,
	KEY_PAGEDOWN: C.REQ_SCR_DPAGE,
	KEY_PAGEUP:   C.REQ_SCR_UPAGE,
	KEY_RIGHT:    C.REQ_RIGHT_ITEM,
	KEY_UP:       C.REQ_UP_ITEM,
}

var errList = map[C.int]string{
	C.E_SYSTEM_ERROR:    "System error occurred",
	C.E_BAD_ARGUMENT:    "Incorrect or out-of-range argument",
	C.E_POSTED:          "Form has already been posted",
	C.E_CONNECTED:       "Field is already connected to a form",
	C.E_BAD_STATE:       "Bad state",
	C.E_NO_ROOM:         "No room",
	C.E_NOT_POSTED:      "Form has not been posted",
	C.E_UNKNOWN_COMMAND: "Unknown command",
	C.E_NO_MATCH:        "No match",
	C.E_NOT_SELECTABLE:  "Not selectable",
	C.E_NOT_CONNECTED:   "Field is not connected to a form",
	C.E_REQUEST_DENIED:  "Request denied",
	C.E_INVALID_FIELD:   "Invalid field",
	C.E_CURRENT:         "Current",
}

func ncursesError(e error) error {
	errno, ok := e.(syscall.Errno)
	if int(errno) == C.OK {
		e = nil
	}
	if ok {
		errstr, ok := errList[C.int(errno)]
		if ok {
			return errors.New(errstr)
		}
	}
	return e
}
