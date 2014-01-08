// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <stdlib.h>
// #include <curses.h>
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

// NewWindow creates a window of size h(eight) and w(idth) at y, x
func NewWindow(h, w, y, x int) (window *Window, err error) {
	window = &Window{C.newwin(C.int(h), C.int(w), C.int(y), C.int(x))}
	if window.win == nil {
		err = errors.New("Failed to create a new window")
	}
	return
}

// AddChar prints a single character to the window. The character can be
// OR'd together with attributes and colors.
func (w *Window) AddChar(ach Char) {
	C.waddch(w.win, C.chtype(ach))
}

// MoveAddChar prints a single character to the window at the specified
// y x coordinates. See AddChar for more info.
func (w *Window) MoveAddChar(y, x int, ach Char) {
	C.mvwaddch(w.win, C.int(y), C.int(x), C.chtype(ach))
}

// Turn off character attribute.
func (w *Window) AttrOff(attr Char) (err error) {
	if C.ncurses_wattroff(w.win, C.int(attr)) == C.ERR {
		err = errors.New(fmt.Sprintf("Failed to unset attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

// Turn on character attribute
func (w *Window) AttrOn(attr Char) (err error) {
	if C.ncurses_wattron(w.win, C.int(attr)) == C.ERR {
		err = errors.New(fmt.Sprintf("Failed to set attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

// AttrSet sets the attributes to the given value
func (w *Window) AttrSet(attr Char) error {
	if C.ncurses_wattrset(w.win, C.int(attr)) == C.ERR {
		return errors.New("Failed to set attributes")
	}
	return nil
}

// SetBackground fills the background with the supplied attributes and/or
// characters.
func (w *Window) SetBackground(attr Char) {
	C.wbkgd(w.win, C.chtype(attr))
}

// Background returns the current background attributes
func (w *Window) Background() Char {
	return Char(C.ncurses_getbkgd(w.win))
}

// Border uses the characters supplied to draw a border around the window.
// t, b, r, l, s correspond to top, bottom, right, left and side respectively.
func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br Char) error {
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
func (w *Window) Box(vch, hch Char) error {
	if C.box(w.win, C.chtype(vch), C.chtype(hch)) == C.ERR {
		return errors.New("Failed to draw box around window")
	}
	return nil
}

// Clears the screen and the underlying virtual screen. This forces the entire
// screen to be rewritten from scratch. This will cause likely cause a
// noticeable flicker because the screen is completely cleared before
// redrawing it. This is probably not what you want. Instead, you should
// probably use the Erase() function. It is the same as called Erase() followed
// by a call to ClearOk().
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
func (w *Window) Color(pair int16) {
	C.wcolor_set(w.win, C.short(ColorPair(pair)), nil)
}

// ColorOff turns the specified color pair off
func (w *Window) ColorOff(pair int16) error {
	if C.ncurses_wattroff(w.win, C.int(ColorPair(pair))) == C.ERR {
		return errors.New("Failed to enable color pair")
	}
	return nil
}

// Normally color pairs are turned on via attron() in ncurses but this
// implementation chose to make it seperate
func (w *Window) ColorOn(pair int16) error {
	if C.ncurses_wattron(w.win, C.int(ColorPair(pair))) == C.ERR {
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

// DelChar deletes the character at the current cursor position, moving all
// characters to the right of that position one space to the left and appends
// a blank character at the end.
func (w *Window) DelChar() error {
	if err := C.wdelch(w.win); err != C.OK {
		return errors.New("An error occurred when trying to delete " +
			"character")
	}
	return nil
}

// MoveDelChar deletes the character at the givin cursor coordinates, moving all
// characters to the right of that position one space to the left and appends
// a blank character at the end.
func (w *Window) MoveDelChar(y, x int) error {
	if err := C.mvwdelch(w.win, C.int(y), C.int(x)); err != C.OK {
		return errors.New("An error occurred when trying to delete " +
			"character")
	}
	return nil
}

// Delete the window. This function must be called to ensure memory is freed
// to prevent memory leaks once you are done with the window.
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
func (w *Window) Derived(height, width, y, x int) *Window {
	return &Window{C.derwin(w.win, C.int(height), C.int(width), C.int(y),
		C.int(x))}
}

// Duplicate the window, creating an exact copy.
func (w *Window) Duplicate() *Window {
	return &Window{C.dupwin(w.win)}
}

// Test whether the given coordinates are within the window or not
func (w *Window) Enclose(y, x int) bool {
	return bool(C.wenclose(w.win, C.int(y), C.int(x)))
}

// Erase the contents of the window, clearing it. This function allows the
// underlying structures to be updated efficiently and thereby provide smooth
// updates to the terminal when frequently clearing and re-writing the window
// or screen.
func (w *Window) Erase() {
	C.werase(w.win)
}

// GetChar retrieves a character from standard input stream and returns it.
// In the event of an error or if the input timeout has expired (ie. if
// Timeout() has been set to zero or a positive value and no characters have
// been received) the value returned will be zero (0)
func (w *Window) GetChar() Key {
	ch := C.wgetch(w.win)
	if ch == C.ERR {
		ch = 0
	}
	return Key(ch)
}

// MoveGetChar moves the cursor to the given position and gets a character
// from the input stream
func (w *Window) MoveGetChar(y, x int) Key {
	return Key(C.mvwgetch(w.win, C.int(y), C.int(x)))
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
func (w *Window) CursorYX() (int, int) {
	var cy, cx C.int
	C.ncurses_getyx(w.win, &cy, &cx)
	return int(cy), int(cx)
}

// HLine draws a horizontal line starting at y, x and ending at width using
// the specified character
func (w *Window) HLine(y, x int, ch Char, wid int) {
	C.mvwhline(w.win, C.int(y), C.int(x), C.chtype(ch), C.int(wid))
	return
}

// InChar returns the character at the current position in the curses window
func (w *Window) InChar() Char {
	return Char(C.winch(w.win))
}

// MoveInChar returns the character at the designated coordates in the curses
// window
func (w *Window) MoveInChar(y, x int) Char {
	return Char(C.mvwinch(w.win, C.int(y), C.int(x)))
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

// LineTouched returns true if the line has been touched; returns false
// otherwise
func (w *Window) LineTouched(line int) bool {
	return bool(C.is_linetouched(w.win, C.int(line)))
}

// Returns the maximum size of the Window. Note that it uses ncurses idiom
// of returning y then x.
func (w *Window) MaxYX() (int, int) {
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

// NoutRefresh, or No Output Refresh, flags the window for redrawing but does
// not output the changes to the terminal (screen). Essentially, the output is
// buffered and a call to Update() flushes the buffer to the terminal. This
// function provides a speed increase over calling Refresh() when multiple
// windows are involved because only the final output is
// transmitted to the terminal.
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

// Parent returns a pointer to a Sub-window's parent, or nil if the window
// has no parent
func (w *Window) Parent() *Window {
	p := C.ncurses_wgetparent(w.win)
	if p == nil {
		return nil
	}
	return &Window{p}
}

// Print a string to the given window. See the fmt package in the standard
// library for more information. In order to simulate the 'n' version
// of functions (like addnstr) just slice your string to the maximum
// length before passing it as an argument.
// window.Print("My line which should be clamped to 20 characters"[:20])
func (w *Window) Print(args ...interface{}) {
	w.Printf("%s", fmt.Sprint(args...))
}

// Printf functions the same as the stardard library's fmt package. See Print
// for more details.
func (w *Window) Printf(format string, args ...interface{}) {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))

	C.waddstr(w.win, cstr)
}

// Println behaves the s as Println in the stanard library's fmt package.
// See Print for more information.
func (w *Window) Println(args ...interface{}) {
	w.Printf("%s", fmt.Sprintln(args...))
}

// MovePrint moves the cursor to the specified coordinates and prints the
// supplied message. See Print for more details.The first two arguments are the
// coordinates to print to.
func (w *Window) MovePrint(y, x int, args ...interface{}) {
	w.MovePrintf(y, x, "%s", fmt.Sprint(args...))
}

// MovePrintf moves the cursor to coordinates and prints the message using
// the specified format. See Printf and MovePrint for more information.
func (w *Window) MovePrintf(y, x int, format string, args ...interface{}) {
	cstr := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cstr))

	C.mvwaddstr(w.win, C.int(y), C.int(x), cstr)
}

// MovePrintln moves the cursor to coordinates and prints the message. See
// Println and MovePrint for more details.
func (w *Window) MovePrintln(y, x int, args ...interface{}) {
	w.MovePrintf(y, x, "%s", fmt.Sprintln(args...))
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
func (w *Window) Sub(height, width, y, x int) *Window {
	return &Window{C.subwin(w.win, C.int(height), C.int(width), C.int(y),
		C.int(x))}
}

// Standend turns off Standout mode, which is equivalent AttrSet(A_NORMAL)
func (w *Window) Standend() error {
	if C.ncurses_wstandend(w.win) == C.ERR {
		return errors.New("Failed to set standend")
	}
	return nil
}

// Standout is equivalent to AttrSet(A_STANDOUT)
func (w *Window) Standout() error {
	if C.ncurses_wstandout(w.win) == C.ERR {
		return errors.New("Failed to set standout")
	}
	return nil
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

// Timeout sets the window to blocking or non-blocking read mode. Calls to
// GetCh will behave in the following manor depending on the value of delay:
// <= -1 - blocking mode is set (blocks indefinately)
// ==  0 - non-blocking; returns zero (0)
// >=  1 - blocks for delay in milliseconds; returns zero (0)
func (w *Window) Timeout(delay int) {
	C.wtimeout(w.win, C.int(delay))
}

// Touch indicates that the window contains changes which should be updated
// on the next call to Refresh
func (w *Window) Touch() error {
	if C.ncurses_touchwin(w.win) == C.ERR {
		return errors.New("Failed to Touch window")
	}
	return nil
}

// Touched returns true if window will be updated on the next Refresh
func (w *Window) Touched() bool {
	return bool(C.is_wintouched(w.win))
}

// Touchline behaves like Touch but only effects count number of lines,
// beginning at start
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
func (w *Window) VLine(y, x int, ch Char, wid int) {
	C.mvwvline(w.win, C.int(y), C.int(x), C.chtype(ch), C.int(wid))
}

// YX returns the current coordinates of the Window. Note that it uses
// ncurses idiom of returning y then x.
func (w *Window) YX() (int, int) {
	var y, x C.int
	C.ncurses_getbegyx(w.win, &y, &x)
	return int(y), int(x)
}
