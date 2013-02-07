// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <stdlib.h>
// #include <ncurses.h>
// #include "goncurses.h"
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type Window struct {
	win *C.WINDOW
}

// AddChar prints a single character to the window. The character can be
// OR'd together with attributes and colors.
func (w *Window) AddChar(ach Character) {
	C.waddch(w.win, C.chtype(ach))
}

// MoveAddChar prints a single character to the window at the specified 
// y x coordinates. See AddChar for more info.
func (w *Window) MoveAddChar(y, x int, ach Character) {
	C.mvwaddch(w.win, C.int(y), C.int(x), C.chtype(ach))
}

// Turn off character attribute.
func (w *Window) AttrOff(attr int) (err error) {
	if C.wattroff(w.win, C.int(attr)) == C.ERR {
		err = errors.New(fmt.Sprintf("Failed to unset attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

// Turn on character attribute
func (w *Window) AttrOn(attr int) (err error) {
	if C.wattron(w.win, C.int(attr)) == C.ERR {
		err = errors.New(fmt.Sprintf("Failed to set attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

// SetBackground fills the background with the supplied attributes and/or
// characters.
func (w *Window) SetBackground(attr Character) {
	C.wbkgd(w.win, C.chtype(attr))
}

// Background returns the current background attributes
func (w *Window) Background() Character {
	return Character(C.ncurses_getbkgd(w.win))
}

// Border uses the characters supplied to draw a border around the window.
// t, b, r, l, s correspond to top, bottom, right, left and side respectively.
func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) error {
	res := C.wborder(w.win, C.chtype(ls), C.chtype(rs), C.chtype(ts),
		C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl),
		C.chtype(br))
	if res == C.ERR {
		return errors.New("Failed to draw box around window")
	}
	return nil
}

// Box draws a border around the given window. For complete control over the
// characters used to draw the border use Border()
func (w *Window) Box(vch, hch int) error {
	if C.box(w.win, C.chtype(vch), C.chtype(hch)) == C.ERR {
		return errors.New("Failed to draw box around window")
	}
	return nil
}

// Clear the screen
func (w *Window) Clear() error {
	if C.wclear(w.win) == C.ERR {
		return errors.New("Failed to clear screen")
	}
	return nil
}

// ClearOk clears the window completely prior to redrawing it. If called
// on stdscr then the whole screen is redrawn no matter which window has
// Refresh() called on it. Defaults to False.
func (w *Window) ClearOk(ok bool) {
	C.clearok(w.win, C.bool(ok))
}

// Clear starting at the current cursor position, moving to the right, to the 
// bottom of window
func (w *Window) ClearToBottom() error {
	if C.wclrtobot(w.win) == C.ERR {
		return errors.New("Failed to clear bottom of window")
	}
	return nil
}

// Clear from the current cursor position, moving to the right, to the end 
// of the line
func (w *Window) ClearToEOL() error {
	if C.wclrtoeol(w.win) == C.ERR {
		return errors.New("Failed to clear to end of line")
	}
	return nil
}

// Color sets the forground/background color pair for the entire window
func (w *Window) Color(pair byte) {
	C.wcolor_set(w.win, C.short(C.COLOR_PAIR(C.int(pair))), nil)
}

// ColorOff turns the specified color pair off
func (w *Window) ColorOff(pair byte) error {
	if C.wattroff(w.win, C.COLOR_PAIR(C.int(pair))) == C.ERR {
		return errors.New("Failed to enable color pair")
	}
	return nil
}

// Normally color pairs are turned on via attron() in ncurses but this
// implementation chose to make it seperate
func (w *Window) ColorOn(pair byte) error {
	if C.wattron(w.win, C.COLOR_PAIR(C.int(pair))) == C.ERR {
		return errors.New("Failed to enable color pair")
	}
	return nil
}

// Copy is similar to Overlay and Overwrite but provides a finer grain of
// control. 
func (w *Window) Copy(src *Window, sy, sx, dtr, dtc, dbr, dbc int,
	overlay bool) error {
	var ol int
	if overlay {
		ol = 1
	}
	if C.copywin(src.win, w.win, C.int(sy), C.int(sx),
		C.int(dtr), C.int(dtc), C.int(dbr), C.int(dbc), C.int(ol)) ==
		C.ERR {
		return errors.New("Failed to copy window")
	}
	return nil
}

// DelChar
func (w *Window) DelChar(coord ...int) error {
	if len(coord) > 2 {
		return errors.New(fmt.Sprintf("Invalid number of arguments, "+
			"expected 2, got %d", len(coord)))
	}
	var err C.int
	if len(coord) > 1 {
		var x int
		y := coord[0]
		if len(coord) > 2 {
			x = coord[1]
		}
		err = C.mvwdelch(w.win, C.int(y), C.int(x))
	} else {
		err = C.wdelch(w.win)
	}
	if err != C.OK {
		return errors.New("An error occurred when trying to delete " +
			"character")
	}
	return nil
}

// Delete the window
func (w *Window) Delete() error {
	if C.delwin(w.win) == C.ERR {
		return errors.New("Failed to delete window")
	}
	w = nil
	return nil
}

// Derived creates a new window of height and width at the coordinates
// y, x.  These coordinates are relative to the original window thereby 
// confining the derived window to the area of original window. See the
// SubWindow function for additional notes.
func (w *Window) Derived(height, width, y, x int) Window {
	return Window{C.derwin(w.win, C.int(height), C.int(width), C.int(y),
		C.int(x))}
}

// Duplicate the window, creating an exact copy.
func (w *Window) Duplicate() Window {
	return Window{C.dupwin(w.win)}
}

// Test whether the given mouse coordinates are within the window or not
func (w *Window) Enclose(y, x int) bool {
	return bool(C.wenclose(w.win, C.int(y), C.int(x)))
}

// Erase the contents of the window, effectively clearing it
func (w *Window) Erase() {
	C.werase(w.win)
}

// GetChar retrieves a character from standard input stream and returns it
func (w *Window) GetChar() Key {
	return Key(C.wgetch(w.win))
}

// MoveGetChar moves the cursor to the given position and gets a character
// from the input stream
func (w *Window) MoveGetChar(y, x int) Key {
	return Key(C.mvwgetch(C.int(y), C.int(x)))
}

// GetString reads at most 'n' characters entered by the user from the Window. 
// Attempts to enter greater than 'n' characters will elicit a 'beep'
func (w *Window) GetString(n int) (string, error) {
	cstr := make([]C.char, n)
	if C.wgetnstr(w.win, (*C.char)(&cstr[0]), C.int(n)) == C.ERR {
		return "", errors.New("Failed to retrieve string from input stream")
	}
	return C.GoString(&cstr[0]), nil
}

// Getyx returns the current cursor location in the Window. Note that it uses 
// ncurses idiom of returning y then x.
func (w *Window) Getyx() (int, int) {
	// In some cases, getxy() and family are macros which don't play well with
	// cgo
	var cy, cx C.int
	C.ncurses_getyx(w.win, &cy, &cx)
	return int(cy), int(cx)
}

// HLine draws a horizontal line starting at y, x and ending at width using 
// the specified character
func (w *Window) HLine(y, x int, ch Character, wid int) {
	C.mvwhline(w.win, C.int(y), C.int(x), C.chtype(ch), C.int(wid))
	return
}

// InChar returns the character at the current position in the curses window
func (w *Window) InChar() Character {
	return Character(C.winch(w.win))
}

// MoveInChar returns the character at the designated coordates in the curses
// window
func (w *Window) MoveInChar(y, x int) Character {
	return Character(C.mvwinch(w.win, C.int(y), C.int(x)))
}

// IsCleared returns the value set in ClearOk
func (w *Window) IsCleared() bool {
	return bool(C.ncurses_is_cleared(w.win))
}

// IsKeypad returns the value set in Keypad
func (w *Window) IsKeypad() bool {
	return bool(C.ncurses_is_keypad(w.win))
}

// Keypad turns on/off the keypad characters, including those like the F1-F12 
// keys and the arrow keys
func (w *Window) Keypad(keypad bool) error {
	var err C.int
	if err = C.keypad(w.win, C.bool(keypad)); err == C.ERR {
		return errors.New("Unable to set keypad mode")
	}
	return nil
}

// Returns the maximum size of the Window. Note that it uses ncurses idiom
// of returning y then x.
func (w *Window) Maxyx() (int, int) {
	var cy, cx C.int
	C.ncurses_getmaxyx(w.win, &cy, &cx)
	return int(cy), int(cx)
}

// Move the cursor to the specified coordinates within the window
func (w *Window) Move(y, x int) {
	C.wmove(w.win, C.int(y), C.int(x))
	return
}

// MoveWindow moves the location of the window to the specified coordinates
func (w *Window) MoveWindow(y, x int) {
	C.mvwin(w.win, C.int(y), C.int(x))
	return
}

// NoutRefresh flags the window for redrawing. In order to actually perform
// the changes, Update() must be called. This function when coupled with
// Update() provides a speed increase over using Refresh() on each window.
func (w *Window) NoutRefresh() {
	C.wnoutrefresh(w.win)
	return
}

// Overlay copies overlapping sections of src window onto the destination
// window. Non-blank elements are not overwritten.
func (w *Window) Overlay(src *Window) error {
	if C.overlay(src.win, w.win) == C.ERR {
		return errors.New("Failed to overlay window")
	}
	return nil
}

// Overwrite copies overlapping sections of src window onto the destination
// window. This function is considered "destructive" by copying all
// elements of src onto the destination window.
func (w *Window) Overwrite(src *Window) error {
	if C.overwrite(src.win, w.win) == C.ERR {
		return errors.New("Failed to overwrite window")
	}
	return nil
}

func (w *Window) Parent() *Window {
	return &Window{C.ncurses_wgetparent(w.win)}
}

// Print a string to the given window. The first two arguments may be
// coordinates to print to. If only one integer is supplied, it is assumed to
// be the y coordinate, x therefore defaults to 0. In order to simulate the 'n' 
// versions of functions like addnstr use a string slice.
// Examples:
// goncurses.Print("hello!")
// goncurses.Print("hello %s!", "world")
// goncurses.Print(23, "hello!") // moves to 23, 0 and prints "hello!"
// goncurses.Print(5, 10, "hello %s!", "world") // move to 5, 10 and print
//                                              // "hello world!"
func (w *Window) Print(format string, args ...interface{}) {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))

	C.waddstr(w.win, cstr)
}

// MovePrint moves the cursor to the specified coordinates and prints the supplied message.
// See Print for more details
func (w *Window) MovePrint(y, x int, format string, args ...interface{}) {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))

	C.mvwaddstr(w.win, C.int(y), C.int(x), cstr)
}

// Refresh the window so it's contents will be displayed
func (w *Window) Refresh() {
	C.wrefresh(w.win)
}

// Resize the window to new height, width
func (w *Window) Resize(height, width int) {
	C.wresize(w.win, C.int(height), C.int(width))
}

// Scroll the contents of the window. Use a negative number to scroll up,
// a positive number to scroll down. ScrollOk Must have been called prior.
func (w *Window) Scroll(n int) {
	C.wscrl(w.win, C.int(n))
}

// ScrollOk sets whether scrolling will work
func (w *Window) ScrollOk(ok bool) {
	C.scrollok(w.win, C.bool(ok))
}

// SubWindow creates a new window of height and width at the coordinates
// y, x.  This window shares memory with the original window so changes
// made to one window are reflected in the other. It is necessary to call
// Touch() on this window prior to calling Refresh in order for it to be
// displayed.
func (w *Window) Sub(height, width, y, x int) Window {
	return Window{C.subwin(w.win, C.int(height), C.int(width), C.int(y),
		C.int(x))}
}

// Sync updates all parent or child windows which were created via
// SubWindow() or DerivedWindow(). Argument can be one of: SYNC_DOWN, which
// syncronizes all parent windows (done by Refresh() by default so should
// rarely, if ever, need to be called); SYNC_UP, which updates all child
// windows to match any updates made to the parent; and, SYNC_CURSOR, which
// updates the cursor position only for all windows to match the parent window
func (w *Window) Sync(sync int) {
	switch sync {
	case SYNC_DOWN:
		C.wsyncdown(w.win)
	case SYNC_CURSOR:
		C.wcursyncup(w.win)
	case SYNC_UP:
		C.wsyncup(w.win)
	}
}

// Touch indicates that the window contains changes which should be updated
// on the next call to Refresh
func (w *Window) Touch() error {
	// may not use touchwin() directly. cgo does not handle macros well.
	//y, _ := w.Maxyx()
	//C.wtouchln(w.win, 0, C.int(y), 1)
	if C.ncurses_touchwin(w.win) == C.ERR {
		return errors.New("Failed to Touch window")
	}
	return nil
}

// Touched returns true if window will be updated on the next Refresh
func (w *Window) Touched() bool {
	return bool(C.is_wintouched(w.win))
}

func (w *Window) TouchLine(start, count int) error {
	if C.touchline(w.win, C.int(start), C.int(count)) == C.ERR {
		return errors.New("Error in call to TouchLine")
	}
	return nil
}

// UnTouch indicates the window should not be updated on the next call to
// Refresh
func (w *Window) UnTouch() {
	C.ncurses_untouchwin(w.win)
}

// VLine draws a verticle line starting at y, x and ending at height using 
// the specified character
func (w *Window) VLine(y, x, ch, h int) {
	// TODO: move portion
	C.mvwvline(w.win, C.int(y), C.int(x), C.chtype(ch), C.int(h))
	return
}
