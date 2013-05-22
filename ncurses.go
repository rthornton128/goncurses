// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* ncurses library

   1. No functions which operate only on stdscr have been implemented because 
   it makes little sense to do so in a Go implementation. Stdscr is treated the
   same as any other window.

   2. Whenever possible, versions of ncurses functions which could potentially
   have a buffer overflow, like the getstr() family of functions, have not been
   implemented. Instead, only the mvwgetnstr() and wgetnstr() can be used. */
package goncurses

/* 
#cgo pkg-config: ncurses
#include <ncurses.h>
#include "goncurses.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// BaudRate returns the speed of the terminal in bits per second
func BaudRate() int {
	return int(C.baudrate())
}

// Beep requests the terminal make an audible bell or, if not available,
// flashes the screen. Note that screen flashing doesn't work on all
// terminals
func Beep() {
	C.beep()
}

// Turn on/off buffering; raw user signals are passed to the program for
// handling. Overrides raw mode
func CBreak(on bool) {
	if on {
		C.cbreak()
		return
	}
	C.nocbreak()
}

// Test whether colour values can be changed
func CanChangeColor() bool {
	return bool(C.bool(C.can_change_color()))
}

// Get RGB values for specified colour
func ColorContent(col int) (int, int, int) {
	var r, g, b C.short
	C.color_content(C.short(col), (*C.short)(&r), (*C.short)(&g),
		(*C.short)(&b))
	return int(r), int(g), int(b)
}

// Return the value of a color pair which can be passed to functions which
// accept attributes like AddChar or AttrOn/Off.
func ColorPair(pair int) int {
	return int(C.COLOR_PAIR(C.int(pair)))
}

// CursesVersion returns the version of the ncurses library currently linked to
func CursesVersion() string {
	return C.GoString(C.curses_version())
}

// Set the cursor visibility. Options are: 0 (invisible/hidden), 1 (normal)
// and 2 (extra-visible)
func Cursor(vis byte) error {
	if C.curs_set(C.int(vis)) == C.ERR {
		return errors.New("Failed to enable ")
	}
	return nil
}

// Echo turns on/off the printing of typed characters
func Echo(on bool) {
	if on {
		C.echo()
		return
	}
	C.noecho()
}

// Must be called prior to exiting the program in order to make sure the
// terminal returns to normal operation
func End() {
	C.endwin()
}

// Flash requests the terminal flashes the screen or, if not available,
// make an audible bell. Note that screen flashing doesn't work on all
// terminals
func Flash() {
	C.flash()
}

// FlushInput flushes all input
func FlushInput() error {
	if C.flushinp() == C.ERR {
		return errors.New("Flush input failed")
	}
	return nil
}

// Returns an array of integers representing the following, in order:
// x, y and z coordinates, id of the device, and a bit masked state of
// the devices buttons
func GetMouse() ([]int, error) {
	if bool(C.ncurses_has_mouse()) != true {
		return nil, errors.New("Mouse support not enabled")
	}
	var event C.MEVENT
	if C.getmouse(&event) != C.OK {
		return nil, errors.New("Failed to get mouse event")
	}
	return []int{int(event.x), int(event.y), int(event.z), int(event.id),
		int(event.bstate)}, nil
}

// Behaves like cbreak() but also adds a timeout for input. If timeout is
// exceeded after a call to Getch() has been made then GetChar will return
// with an error.
func HalfDelay(delay int) error {
	var cerr C.int
	if delay > 0 {
		cerr = C.halfdelay(C.int(delay))
	}
	if cerr == C.ERR {
		return errors.New("Unable to set delay mode")
	}
	return nil
}

// HasColors returns true if terminal can display colors
func HasColors() bool {
	return bool(C.has_colors())
}

// HasInsertCharacter return true if the terminal has insert and delete
// character capabilities
func HasInsertCharacter() bool {
	return bool(C.has_ic())
}

// HasInsertLine returns true if the terminal has insert and delete line
// capabilities. See ncurses documentation for more details
func HasInsertLine() bool {
	return bool(C.has_il())
}

// HasKey returns true if terminal recognized the given character
func HasKey(ch Key) bool {
	if C.has_key(C.int(ch)) == 1 {
		return true
	}
	return false
}

// InitColor is used to set 'color' to the specified RGB values. Values may
// be between 0 and 1000.
func InitColor(col int, r, g, b int) error {
	if C.init_color(C.short(col), C.short(r), C.short(g),
		C.short(b)) == C.ERR {
		return errors.New("Failed to set new color definition")
	}
	return nil
}

// InitPair sets a colour pair designated by 'pair' to fg and bg colors
func InitPair(pair byte, fg, bg int) error {
	if pair == 0 || C.int(pair) > (C.COLOR_PAIRS-1) {
		return errors.New("Invalid color pair selected")
	}
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == C.ERR {
		return errors.New("Failed to init color pair")
	}
	return nil
}

// Initialize the ncurses library. You must run this function prior to any 
// other goncurses function in order for the library to work
func Init() (stdscr Window, err error) {
	stdscr = Window{C.initscr()}
	if unsafe.Pointer(stdscr.win) == nil {
		err = errors.New("An error occurred initializing ncurses")
	}
	return
}

// IsEnd returns true if End() has been called, otherwise false
func IsEnd() bool {
	return bool(C.isendwin())
}

// IsTermResized returns true if ResizeTerm would modify any current Windows 
// if called with the given parameters
func IsTermResized(nlines, ncols int) bool {
	return bool(C.is_term_resized(C.int(nlines), C.int(ncols)))
}

// Returns a string representing the value of input returned by Getch
func KeyString(k Key) string {
	key, ok := keyList[k]
	if !ok {
		key = fmt.Sprintf("%c", int(k))
	}
	return key
}

func Mouse() bool {
	return bool(C.ncurses_has_mouse())
}

func MouseInterval() {
}

// MouseMask accepts a single int of OR'd mouse events. If a mouse event
// is triggered, GetChar() will return KEY_MOUSE. To retrieve the actual
// event use GetMouse() to pop it off the queue. Pass a pointer as the 
// second argument to store the prior events being monitored or nil.
func MouseMask(mask int, old *int) (m int) {
	if bool(C.ncurses_has_mouse()) {
		m = int(C.mousemask((C.mmask_t)(mask),
			(*C.mmask_t)(unsafe.Pointer(old))))
	}
	return
}

// NapMilliseconds is used to sleep for ms milliseconds
func NapMilliseconds(ms int) {
	C.napms(C.int(ms))
}

// NewWindow creates a window of size h(eight) and w(idth) at y, x
func NewWindow(h, w, y, x int) (window Window, err error) {
	window = Window{C.newwin(C.int(h), C.int(w), C.int(y), C.int(x))}
	if window.win == nil {
		err = errors.New("Failed to create a new window")
	}
	return
}

// NL turns newline translation on/off.
func NL(on bool) {
	if on {
		C.nl()
		return
	}
	C.nonl()
}

// Raw turns on input buffering; user signals are disabled and the key strokes 
// are passed directly to input. Set to false if you wish to turn this mode
// off
func Raw(on bool) {
	if on {
		C.raw()
		return
	}
	C.noraw()
}

// ResizeTerm will attempt to resize the terminal. This only has an effect if
// the terminal is in an XWindows (GUI) environment.
func ResizeTerm(nlines, ncols int) error {
	if C.resizeterm(C.int(nlines), C.int(ncols)) == C.ERR {
		return errors.New("Failed to resize terminal")
	}
	return nil
}

// Enables colors to be displayed. Will return an error if terminal is not
// capable of displaying colors
func StartColor() error {
	if C.has_colors() == C.bool(false) {
		return errors.New("Terminal does not support colors")
	}
	if C.start_color() == C.ERR {
		return errors.New("Failed to enable color mode")
	}
	return nil
}

// UnGetChar places the character back into the input queue
func UnGetChar(ch Character) {
	C.ungetch(C.int(ch))
}

// Update the screen, refreshing all windows
func Update() error {
	if C.doupdate() == C.ERR {
		return errors.New("Failed to update")
	}
	return nil
}

// UseEnvironment specifies whether the LINES and COLUMNS environmental
// variables should be used or not
func UseEnvironment(use bool) {
	C.use_env(C.bool(use))
}
