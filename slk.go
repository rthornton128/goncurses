// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <stdlib.h>
// #include <curses.h>
import "C"

import (
	"errors"
	"unsafe"
)

type SlkFormat byte

const (
	SLK_323        SlkFormat = iota // 8 labels; 3-2-3 arrangement
	SLK_44                          // 8 labels; 4-4 arrangement
	SLK_PC444                       // 12 labels; 4-4-4 arrangement
	SLK_PC444INDEX                  // 12 labels; 4-4-4 with index line
)

type SlkJustify byte

const (
	SLK_LEFT SlkJustify = iota
	SLK_CENTER
	SLK_RIGHT
)

// Initializes the soft-key labels with the given format; keys like the
// F1-F12 keys on most keyboards. After a call to SlkRefresh a bar at the
// bottom of the standard screen returned by Init will be displayed. This
// function MUST be called prior to Init()
func SlkInit(f SlkFormat) {
	C.slk_init(C.int(f))
}

// SlkSet sets the 'labnum' text to the supplied 'label'. Labels must not
// be greater than 8 characters
func SlkSet(labnum int, label string, just SlkJustify) error {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))

	if C.slk_set(C.int(labnum), (*C.char)(cstr), C.int(just)) == C.ERR {
		return errors.New("Soft-keys or terminal not initialized")
	}
	return nil
}

// SlkRefresh behaves the same as Window.Refresh. Most applications would use
// SlkNoutRefresh because a Window.Refresh is likely to follow
func SlkRefresh() error {
	if C.slk_refresh() == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkNoutFresh behaves like Window.NoutRefresh
func SlkNoutRefresh() error {
	if C.slk_noutrefresh() == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkLabel returns the label for the given key
func SlkLabel(labnum int) string {
	return C.GoString(C.slk_label(C.int(labnum)))
}

// SlkClear removes the soft-key labels from the screen
func SlkClear() error {
	if C.slk_clear() == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkRestore restores the soft-key labels to the screen after an SlkClear()
func SlkRestore() error {
	if C.slk_restore() == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkTouch behaves just like Window.Touch
func SlkTouch() error {
	if C.slk_touch() == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkColor sets the color pair for the soft-keys
func SlkColor(cp int16) error {
	if C.slk_color(C.short(cp)) == C.ERR {
		return errors.New("Invalid color pair or soft-keys not initialized.")
	}
	return nil
}

/* TODO: Not available in PDCurses
// SlkAttribute returns the currently set attributes
func SlkAttribute() Char {
	return Char(C.slk_attr())
}*/

// SlkSetAttribute sets the OR'd attributes to use
func SlkSetAttribute(attr Char) error {
	if C.slk_attrset(C.chtype(attr)) == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkAttributeOn turns on the given OR'd attributes without turning any off
func SlkAttributeOn(attr Char) error {
	if C.slk_attron(C.chtype(attr)) == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}

// SlkAttributeOff turns off the given OR'd attributes withoiut turning any on
func SlkAttributeOff(attr Char) error {
	if C.slk_attroff(C.chtype(attr)) == C.ERR {
		return errors.New("Soft-keys or terminal not initialized.")
	}
	return nil
}
