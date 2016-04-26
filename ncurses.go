// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: ncurses
// #cgo windows CFLAGS: -DNCURSES_MOUSE_VERSION
// #cgo windows LDFLAGS: -lpdcurses
// #cgo darwin openbsd LDFLAGS: -lncurses
// #include <curses.h>
// #include "goncurses.h"
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

func TypeAhead(fd int) int {

	return int(C.typeahead(C.int(fd)))
}

// Test whether colour values can be changed
func CanChangeColor() bool {
	return bool(C.bool(C.can_change_color()))
}

// Get RGB values for specified colour
func ColorContent(col int16) (int16, int16, int16) {
	var r, g, b C.short
	C.color_content(C.short(col), (*C.short)(&r), (*C.short)(&g),
		(*C.short)(&b))
	return int16(r), int16(g), int16(b)
}

// Return the value of a color pair which can be passed to functions which
// accept attributes like AddChar, AttrOn/Off and Background.
func ColorPair(pair int16) Char {
	return Char(C.ncurses_COLOR_PAIR(C.int(pair)))
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

// HasInsertChar return true if the terminal has insert and delete
// character capabilities
func HasInsertChar() bool {
	return bool(C.has_ic())
}

// HasInsertLine returns true if the terminal has insert and delete line
// capabilities. See ncurses documentation for more details
func HasInsertLine() bool {
	return bool(C.has_il())
}

// HasKey returns true if terminal recognized the given character
func HasKey(ch Key) bool {
	if C.ncurses_has_key(C.int(ch)) == 1 {
		return true
	}
	return false
}

// InitColor is used to set 'color' to the specified RGB values. Values may
// be between 0 and 1000.
func InitColor(col, r, g, b int16) error {
	if C.init_color(C.short(col), C.short(r), C.short(g),
		C.short(b)) == C.ERR {
		return errors.New("Failed to set new color definition")
	}
	return nil
}

// InitPair sets a colour pair designated by 'pair' to fg and bg colors
func InitPair(pair, fg, bg int16) error {
	if pair <= 0 || C.int(pair) > C.int(C.COLOR_PAIRS-1) {
		return errors.New("Color pair out of range")
	}
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == C.ERR {
		return errors.New("Failed to init color pair")
	}
	return nil
}

// Initialize the ncurses library. You must run this function prior to any
// other goncurses function in order for the library to work
func Init() (stdscr *Window, err error) {
	stdscr = &Window{C.initscr()}
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

// PairContent returns the current foreground and background colours
// associated with the given pair
func PairContent(pair int16) (fg int16, bg int16, err error) {
	var f, b C.short
	if C.pair_content(C.short(pair), &f, &b) == C.ERR {
		return -1, -1, errors.New("Invalid color pair")
	}
	return int16(f), int16(b), nil
}

// Nap (sleep; halt execution) for 'ms' milliseconds
func Nap(ms int) {
	C.napms(C.int(ms))
}

// NewLines turns newline translation on/off.
func NewLines(on bool) {
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

// StdScr returns a Window for the underlying stdscr object which represents
// the physical screen. This is the same Window returned by Init and therefore
// not useful unless using NewTerm and other multi-screen related functions.
func StdScr() *Window {
	return &Window{C.stdscr}
}

// UnGetChar places the character back into the input queue
func UnGetChar(ch Char) {
	C.ncurses_ungetch(C.int(ch))
}

// Update the screen, refreshing all windows
func Update() error {
	if C.doupdate() == C.ERR {
		return errors.New("Failed to update")
	}
	return nil
}

// UseDefaultColors tells the curses library to assign the terminal's default
// foreground and background colors to color number -1. This will allow you to
// call InitPair(x, -1, -1) to set both the foreground and backgroun colours
// of pair x to the terminal's default. This function can fail if the terminal
// does not support certain ncurses features like orig_pair or initialize_pair.
func UseDefaultColors() error {
	if C.use_default_colors() == C.ERR {
		return errors.New("Failed to assume default colours.")
	}
	return nil
}

// UseEnvironment specifies whether the LINES and COLUMNS environmental
// variables should be used or not
func UseEnvironment(use bool) {
	C.use_env(C.bool(use))
}
