// goncurses - ncurses library for Go.
//
// Copyright (c) 2011, Rob Thornton 
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without 
// modification, are permitted provided that the following conditions are met:
//
//   * Redistributions of source code must retain the above copyright notice, 
//     this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright notice, 
//     this list of conditions and the following disclaimer in the documentation 
//     and/or other materials provided with the distribution.
//  
//   * Neither the name of the copyright holder nor the names of its 
//     contributors may be used to endorse or promote products derived from this 
//     software without specific prior written permission.
//      
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" 
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE 
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE 
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE 
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR 
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF 
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS 
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN 
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) 
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE 
// POSSIBILITY OF SUCH DAMAGE.

/* ncurses library

   1. No functions which operate only on stdscr have been implemented because 
   it makes little sense to do so in a Go implementation. Stdscr is treated the
   same as any other window.

   2. Whenever possible, versions of ncurses functions which could potentially
   have a buffer overflow, like the getstr() family of functions, have not been
   implemented. Instead, only the mvwgetnstr() and wgetnstr() can be used. */
package goncurses

// #cgo LDFLAGS: -lncurses
// #include <ncurses.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

// Synconize options for Sync() function
const (
	SYNC_NONE   = iota
	SYNC_CURSOR // Sync cursor in all sub/derived windows
	SYNC_DOWN   // Sync changes in all parent windows
	SYNC_UP     // Sync change in all child windows
)

// Definitions for printed characters not found on most keyboards. Ideally, 
// these would not be hard-coded as they are potentially different on
// different systems. However, some ncurses implementations seem to be
// heavily reliant on macros which prevent these definitions from being
// handled by cgo properly. If they don't work for you, you won't be able
// to use them until either a) the Go team works out a way to overcome this
// limitation in godefs/cgo or b) an alternative method is found. Work is
// being done to find a solution from the ncurses source code.
const (
	ACS_DEGREE = iota + 4194406
	ACS_PLMINUS
	ACS_BOARD
	ACS_LANTERN
	ACS_LRCORNER
	ACS_URCORNER
	ACS_LLCORNER
	ACS_ULCORNER
	ACS_PLUS
	ACS_S1
	ACS_S3
	ACS_HLINE
	ACS_S7
	ACS_S9
	ACS_LTEE
	ACS_RTEE
	ACS_BTEE
	ACS_TTEE
	ACS_VLINE
	ACS_LEQUAL
	ACS_GEQUAL
	ACS_PI
	ACS_NEQUAL
	ACS_STERLING
	ACS_BULLET
	ACS_LARROW  = 4194347
	ACS_RARROW  = 4194348
	ACS_DARROW  = 4194349
	ACS_UARROW  = 4194350
	ACS_BLOCK   = 4194352
	ACS_DIAMOND = 4194400
	ACS_CKBOARD = 4194401
)

// Text attributes
const (
	A_NORMAL     = C.A_NORMAL
	A_STANDOUT   = C.A_STANDOUT
	A_UNDERLINE  = C.A_UNDERLINE
	A_REVERSE    = C.A_REVERSE
	A_BLINK      = C.A_BLINK
	A_DIM        = C.A_DIM
	A_BOLD       = C.A_BOLD
	A_PROTECT    = C.A_PROTECT
	A_INVIS      = C.A_INVIS
	A_ALTCHARSET = C.A_ALTCHARSET
	A_CHARTEXT   = C.A_CHARTEXT
)

var attrList = map[C.int]string{
	C.A_NORMAL:     "normal",
	C.A_STANDOUT:   "standout",
	C.A_UNDERLINE:  "underline",
	C.A_REVERSE:    "reverse",
	C.A_BLINK:      "blink",
	C.A_DIM:        "dim",
	C.A_BOLD:       "bold",
	C.A_PROTECT:    "protect",
	C.A_INVIS:      "invis",
	C.A_ALTCHARSET: "altcharset",
	C.A_CHARTEXT:   "chartext",
}

type Chtype C.chtype

// Colors available to ncurses. Combine these with the dim/bold attributes
// for bright/dark versions of each color. These colors can be used for
// both background and foreground colors.
const (
	C_BLACK   = C.COLOR_BLACK
	C_BLUE    = C.COLOR_BLUE
	C_CYAN    = C.COLOR_CYAN
	C_GREEN   = C.COLOR_GREEN
	C_MAGENTA = C.COLOR_MAGENTA
	C_RED     = C.COLOR_RED
	C_WHITE   = C.COLOR_WHITE
	C_YELLOW  = C.COLOR_YELLOW
)

/*var colorList = map[string]C.int{
	"black":   C.COLOR_BLACK,
	"red":     C.COLOR_RED,
	"green":   C.COLOR_GREEN,
	"yellow":  C.COLOR_YELLOW,
	"blue":    C.COLOR_BLUE,
	"magenta": C.COLOR_MAGENTA,
	"cyan":    C.COLOR_CYAN,
	"white":   C.COLOR_WHITE,
}*/

var keyList = map[C.int]string{
	9:               "tab",
	10:              "enter", // On some keyboards?
	C.KEY_DOWN:      "down",
	C.KEY_UP:        "up",
	C.KEY_LEFT:      "left",
	C.KEY_RIGHT:     "right",
	C.KEY_HOME:      "home",
	C.KEY_BACKSPACE: "backspace",
	C.KEY_ENTER:     "enter", // And not others?
	C.KEY_F0:        "F0",
	C.KEY_F0 + 1:    "F1",
	C.KEY_F0 + 2:    "F2",
	C.KEY_F0 + 3:    "F3",
	C.KEY_F0 + 4:    "F4",
	C.KEY_F0 + 5:    "F5",
	C.KEY_F0 + 6:    "F6",
	C.KEY_F0 + 7:    "F7",
	C.KEY_F0 + 8:    "F8",
	C.KEY_F0 + 9:    "F9",
	C.KEY_F0 + 10:   "F10",
	C.KEY_F0 + 11:   "F11",
	C.KEY_F0 + 12:   "F12",
	C.KEY_MOUSE:     "mouse",
	C.KEY_NPAGE:     "page down",
	C.KEY_PPAGE:     "page up",
}

type MMask C.mmask_t

// Mouse button events
const (
	M_ALL            = C.ALL_MOUSE_EVENTS
	M_ALT            = C.BUTTON_ALT      // alt-click
	M_B1_PRESSED     = C.BUTTON1_PRESSED // button 1
	M_B1_RELEASED    = C.BUTTON1_RELEASED
	M_B1_CLICKED     = C.BUTTON1_CLICKED
	M_B1_DBL_CLICKED = C.BUTTON1_DOUBLE_CLICKED
	M_B1_TPL_CLICKED = C.BUTTON1_TRIPLE_CLICKED
	M_B2_PRESSED     = C.BUTTON2_PRESSED // button 2
	M_B2_RELEASED    = C.BUTTON2_RELEASED
	M_B2_CLICKED     = C.BUTTON2_CLICKED
	M_B2_DBL_CLICKED = C.BUTTON2_DOUBLE_CLICKED
	M_B2_TPL_CLICKED = C.BUTTON2_TRIPLE_CLICKED
	M_B3_PRESSED     = C.BUTTON3_PRESSED // button 3
	M_B3_RELEASED    = C.BUTTON3_RELEASED
	M_B3_CLICKED     = C.BUTTON3_CLICKED
	M_B3_DBL_CLICKED = C.BUTTON3_DOUBLE_CLICKED
	M_B3_TPL_CLICKED = C.BUTTON3_TRIPLE_CLICKED
	M_B4_PRESSED     = C.BUTTON4_PRESSED // button 4
	M_B4_RELEASED    = C.BUTTON4_RELEASED
	M_B4_CLICKED     = C.BUTTON4_CLICKED
	M_B4_DBL_CLICKED = C.BUTTON4_DOUBLE_CLICKED
	M_B4_TPL_CLICKED = C.BUTTON4_TRIPLE_CLICKED
	M_CTRL           = C.BUTTON_CTRL           // ctrl-click
	M_SHIFT          = C.BUTTON_SHIFT          // shift-click
	M_POSITION       = C.REPORT_MOUSE_POSITION // mouse moved
)

var mouseEvents = map[string]MMask{
	"button1-pressed":        C.BUTTON1_PRESSED,
	"button1-released":       C.BUTTON1_RELEASED,
	"button1-clicked":        C.BUTTON1_CLICKED,
	"button1-double-clicked": C.BUTTON1_DOUBLE_CLICKED,
	"button1-triple-clicked": C.BUTTON1_TRIPLE_CLICKED,
	"button2-pressed":        C.BUTTON2_PRESSED,
	"button2-released":       C.BUTTON2_RELEASED,
	"button2-clicked":        C.BUTTON2_CLICKED,
	"button2-double-clicked": C.BUTTON2_DOUBLE_CLICKED,
	"button2-triple-clicked": C.BUTTON2_TRIPLE_CLICKED,
	"button3-pressed":        C.BUTTON3_PRESSED,
	"button3-released":       C.BUTTON3_RELEASED,
	"button3-clicked":        C.BUTTON3_CLICKED,
	"button3-double-clicked": C.BUTTON3_DOUBLE_CLICKED,
	"button3-triple-clicked": C.BUTTON3_TRIPLE_CLICKED,
	"button4-pressed":        C.BUTTON4_PRESSED,
	"button4-released":       C.BUTTON4_RELEASED,
	"button4-clicked":        C.BUTTON4_CLICKED,
	"button4-double-clicked": C.BUTTON4_DOUBLE_CLICKED,
	"button4-triple-clicked": C.BUTTON4_TRIPLE_CLICKED,
	//    "button5-pressed": C.BUTTON5_PRESSED,
	//    "button5-released": C.BUTTON5_RELEASED,
	//    "button5-clicked": C.BUTTON5_CLICKED,
	//    "button5-double-clicked": C.BUTTON5_DOUBLE_CLICKED,
	//    "button5-triple-clicked": C.BUTTON5_TRIPLE_CLICKED,
	"shift":    C.BUTTON_SHIFT,
	"ctrl":     C.BUTTON_CTRL,
	"alt":      C.BUTTON_ALT,
	"all":      C.ALL_MOUSE_EVENTS,
	"position": C.REPORT_MOUSE_POSITION,
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

// Return the value of a color pair which can be passed to functions which
// accept attributes like AddChar or AttrOn/Off.
func ColorPair(pair int) int {
	return int(C.COLOR_PAIR(C.int(pair)))
}

// Set the cursor visibility. Options are: 0 (invisible/hidden), 1 (normal)
// and 2 (extra-visible)
func Cursor(vis byte) os.Error {
	if C.curs_set(C.int(vis)) == C.ERR {
		return os.NewError("Failed to enable ")
	}
	return nil
}

// Update the screen, refreshing all windows
func Update() os.Error {
	if C.doupdate() == C.ERR {
		return os.NewError("Failed to update")
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

// Returns an array of integers representing the following, in order:
// x, y and z coordinates, id of the device, and a bit masked state of
// the devices buttons
func GetMouse() ([]int, os.Error) {
	var event C.MEVENT
	if C.getmouse(&event) != C.OK {
		return nil, os.NewError("Failed to get mouse event")
	}
	return []int{int(event.x), int(event.y), int(event.z), int(event.id),
		int(event.bstate)}, nil
}

// Behaves like cbreak() but also adds a timeout for input. If timeout is
// exceeded after a call to Getch() has been made then GetChar will return
// with an error.
func HalfDelay(delay int) os.Error {
	var cerr C.int
	if delay > 0 {
		cerr = C.halfdelay(C.int(delay))
	}
	if cerr == C.ERR {
		return os.NewError("Unable to set delay mode")
	}
	return nil
}

// InitColor is used to set 'color' to the specified RGB values. Values may
// be between 0 and 1000.
func InitColor(col int, r, g, b int) os.Error {
	if C.init_color(C.short(col), C.short(r), C.short(g), C.short(b)) == C.ERR {
		return os.NewError("Failed to set new color definition")
	}
	return nil
}

// InitPair sets a colour pair designated by 'pair' to fg and bg colors
func InitPair(pair byte, fg, bg int) os.Error {
	if pair == 0 || C.int(pair) > (C.COLOR_PAIRS-1) {
		return os.NewError("Invalid color pair selected")
	}
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == C.ERR {
		return os.NewError("Failed to init color pair")
	}
	return nil
}

// Initialize the ncurses library. You must run this function prior to any 
// other goncurses function in order for the library to work
func Init() (stdscr *Window, err os.Error) {
	stdscr = (*Window)(C.initscr())
	err = nil
	if unsafe.Pointer(stdscr) == nil {
		err = os.NewError("An error occurred initializing ncurses")
	}
	return
}

// Returns a string representing the value of input returned by Getch
func Key(k int) (key string) {
	var ok bool
	key, ok = keyList[C.int(k)]
	if !ok {
		key = fmt.Sprintf("%c", k)
	}
	return
}

// MouseMask accepts a single int of OR'd mouse events which should be
// accepted as input events to GetChar. 
func MouseMask(masks ...string) {
	var mousemask MMask
	for _, mask := range masks {
		mousemask |= mouseEvents[mask]
	}
	C.mousemask((C.mmask_t)(mousemask), (*C.mmask_t)(unsafe.Pointer(nil)))
}

// NewWindow creates a windows of size h(eight) and w(idth) at y, x
func NewWindow(h, w, y, x int) (new *Window, err os.Error) {
	new = (*Window)(C.newwin(C.int(h), C.int(w), C.int(y), C.int(x)))
	if unsafe.Pointer(new) == unsafe.Pointer(nil) {
		err = os.NewError("Failed to create a new window")
	}
	return
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

// Enables colors to be displayed. Will return an error if terminal is not
// capable of displaying colors
func StartColor() os.Error {
	if C.has_colors() == C.bool(false) {
		return os.NewError("Terminal does not support colors")
	}
	if C.start_color() == C.ERR {
		return os.NewError("Failed to enable color mode")
	}
	return nil
}

type Window C.WINDOW

// AddChar prints a single character to the window. The character can be
// OR'd together with attributes and colors. If optional first or second
// arguments are given they are the y and x coordinates on the screen
// respectively. If only y is given, x is assumed to be zero.
func (w *Window) AddChar(args ...int) {
	var cattr C.int
	var count, y, x int

	if len(args) > 1 {
		y = args[0]
		count += 1
	}
	if len(args) > 2 {
		x = args[1]
		count += 1
	}
	cattr |= C.int(args[count])
	if count > 0 {
		C.mvwaddch((*C.WINDOW)(w), C.int(y), C.int(x), C.chtype(cattr))
		return
	}
	C.waddch((*C.WINDOW)(w), C.chtype(cattr))
}

// Turn off character attribute.
func (w *Window) AttrOff(attr int) (err os.Error) {
	if C.wattroff((*C.WINDOW)(w), C.int(attr)) == C.ERR {
		err = os.NewError(fmt.Sprintf("Failed to unset attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

// Turn on character attribute
func (w *Window) AttrOn(attr int) (err os.Error) {
	if C.wattron((*C.WINDOW)(w), C.int(attr)) == C.ERR {
		err = os.NewError(fmt.Sprintf("Failed to set attribute: %s",
			attrList[C.int(attr)]))
	}
	return
}

func (w *Window) Background(attr int) {
	C.wbkgd((*C.WINDOW)(w), C.chtype(attr))
}

// Border uses the characters supplied to draw a border around the window.
// t, b, r, l, s correspond to top, bottom, right, left and side respectively.
func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) os.Error {
	res := C.wborder((*C.WINDOW)(w), C.chtype(ls), C.chtype(rs), C.chtype(ts),
		C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl),
		C.chtype(br))
	if res == C.ERR {
		return os.NewError("Failed to draw box around window")
	}
	return nil
}

// Box draws a border around the given window. For complete control over the
// characters used to draw the border use Border()
func (w *Window) Box(vch, hch int) os.Error {
	if C.box((*C.WINDOW)(w), C.chtype(vch), C.chtype(hch)) == C.ERR {
		return os.NewError("Failed to draw box around window")
	}
	return nil
}

// Clear the screen
func (w *Window) Clear() os.Error {
	if C.wclear((*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to clear screen")
	}
	return nil
}

// Clear starting at the current cursor position, moving to the right, to the 
// bottom of window
func (w *Window) ClearToBottom() os.Error {
	if C.wclrtobot((*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to clear bottom of window")
	}
	return nil
}

// Clear from the current cursor position, moving to the right, to the end 
// of the line
func (w *Window) ClearToEOL() os.Error {
	if C.wclrtoeol((*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to clear to end of line")
	}
	return nil
}

// Color sets the forground/background color pair for the entire window
func (w *Window) Color(pair byte) {
	C.wcolor_set((*C.WINDOW)(w), C.short(C.COLOR_PAIR(C.int(pair))), nil)
}

func (w *Window) ColorOff(pair byte) os.Error {
	if C.wattroff((*C.WINDOW)(w), C.COLOR_PAIR(C.int(pair))) == C.ERR {
		return os.NewError("Failed to enable color pair")
	}
	return nil
}

// Normally color pairs are turned on via attron() in ncurses but this
// implementation chose to make it seperate
func (w *Window) ColorOn(pair byte) os.Error {
	if C.wattron((*C.WINDOW)(w), C.COLOR_PAIR(C.int(pair))) == C.ERR {
		return os.NewError("Failed to enable color pair")
	}
	return nil
}

// Copy is similar to Overlay and Overwrite but provides a finer grain of
// control. 
func (w *Window) Copy(src *Window, sy, sx, dtr, dtc, dbr, dbc int,
overlay bool) os.Error {
	var ol int
	if overlay {
		ol = 1
	}
	if C.copywin((*C.WINDOW)(src), (*C.WINDOW)(w), C.int(sy), C.int(sx),
		C.int(dtr), C.int(dtc), C.int(dbr), C.int(dbc), C.int(ol)) ==
		C.ERR {
		return os.NewError("Failed to copy window")
	}
	return nil
}

// Delete the window
func (w *Window) Delete() os.Error {
	if C.delwin((*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to delete window")
	}
	w = nil
	return nil
}

// DerivedWindow creates a new window of height and width at the coordinates
// y, x.  These coordinates are relative to the original window thereby 
// confining the derived window to the area of original window. See the
// SubWindow function for additional notes.
func (w *Window) DerivedWindow(height, width, y, x int) *Window {
	res := C.derwin((*C.WINDOW)(w), C.int(height), C.int(width), C.int(y),
		C.int(x))
	return (*Window)(res)
}

// Duplicate the window, creating an exact copy.
func (w *Window) Duplicate() *Window {
	return (*Window)(C.dupwin((*C.WINDOW)(w)))
}

// Erase the contents of the window, effectively clearing it
func (w *Window) Erase() {
	C.werase((*C.WINDOW)(w))
}

// Get a character from standard input
func (w *Window) GetChar(coords ...int) int {
	var y, x, count int
	if len(coords) > 1 {
		y = coords[0]
		count++
	}
	if len(coords) > 2 {
		x = coords[1]
		count++
	}
	if count > 0 {
		return int(C.mvwgetch((*C.WINDOW)(w), C.int(y), C.int(x)))
	}
	return int(C.wgetch((*C.WINDOW)(w)))
}

// Returns the maximum size of the Window. Note that it uses ncurses idiom
// of returning y then x.
func (w *Window) Maxyx() (int, int) {
	// This hack is necessary to make cgo happy
	return int(w._maxy + 1), int(w._maxx + 1)
}

// Reads at most 'n' characters entered by the user from the Window. Attempts
// to enter greater than 'n' characters will elicit a 'beep'
func (w *Window) GetString(n int) (string, os.Error) {
	cstr := make([]C.char, n)
	if C.wgetnstr((*C.WINDOW)(w), (*C.char)(&cstr[0]), C.int(n)) == C.ERR {
		return "", os.NewError("Failed to retrieve string from input stream")
	}
	return C.GoString(&cstr[0]), nil
}

// Getyx returns the current cursor location in the Window. Note that it uses 
// ncurses idiom of returning y then x.
func (w *Window) Getyx() (int, int) {
	// In some cases, getxy() and family are macros which don't play well with
	// cgo
	return int(w._cury), int(w._curx)
}

// HLine draws a horizontal line starting at y, x and ending at width using 
// the specified character
func (w *Window) HLine(y, x, ch, wid int) {
	C.mvwhline((*C.WINDOW)(w), C.int(y), C.int(x), C.chtype(ch),
		C.int(wid))
	return
}

// Keypad turns on/off the keypad characters, including those like the F1-F12 
// keys and the arrow keys
func (w *Window) Keypad(keypad bool) os.Error {
	var err C.int
	if err = C.keypad((*C.WINDOW)(w), C.bool(keypad)); err == C.ERR {
		return os.NewError("Unable to set keypad mode")
	}
	return nil
}

// Move the cursor to the specified coordinates within the window
func (w *Window) Move(y, x int) {
	C.wmove((*C.WINDOW)(w), C.int(y), C.int(x))
	return
}

// Overlay copies overlapping sections of src window onto the destination
// window. Non-blank elements are not overwritten.
func (w *Window) Overlay(src *Window) os.Error {
	if C.overlay((*C.WINDOW)(src), (*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to overlay window")
	}
	return nil
}

// Overwrite copies overlapping sections of src window onto the destination
// window. This function is considered "destructive" by copying all
// elements of src onto the destination window.
func (w *Window) Overwrite(src *Window) os.Error {
	if C.overwrite((*C.WINDOW)(src), (*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to overwrite window")
	}
	return nil
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
func (w *Window) Print(args ...interface{}) {
	var count, y, x int

	if len(args) > 1 {
		if reflect.TypeOf(args[0]).String() == "int" {
			y = args[0].(int)
			count += 1
		}
	}
	if len(args) > 2 {
		if reflect.TypeOf(args[1]).String() == "int" {
			x = args[1].(int)
			count += 1
		}
	}

	cstr := C.CString(fmt.Sprintf(args[count].(string), args[count+1:]...))
	defer C.free(unsafe.Pointer(cstr))

	if count > 0 {
		C.mvwaddstr((*C.WINDOW)(w), C.int(y), C.int(x), cstr)
		return
	}
	C.waddstr((*C.WINDOW)(w), cstr)
}

// Refresh the window so it's contents will be displayed
func (w *Window) Refresh() {
	C.wrefresh((*C.WINDOW)(w))
}

// Resize the window to new height, width
func (w *Window) Resize(height, width int) {
	C.wresize((*C.WINDOW)(w), C.int(height), C.int(width))
}

// SubWindow creates a new window of height and width at the coordinates
// y, x.  This window shares memory with the original window so changes
// made to one window are reflected in the other. It is necessary to call
// Touch() on this window prior to calling Refresh in order for it to be
// displayed.
func (w *Window) SubWindow(height, width, y, x int) *Window {
	res := C.subwin((*C.WINDOW)(w), C.int(height), C.int(width), C.int(y),
		C.int(x))
	return (*Window)(res)
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
		C.wsyncdown((*C.WINDOW)(w))
	case SYNC_CURSOR:
		C.wcursyncup((*C.WINDOW)(w))
	case SYNC_UP:
		C.wsyncup((*C.WINDOW)(w))
	}
}

// Touch indicates that the window contains changes which should be updated
// on the next call to Refresh
func (w *Window) Touch() {
	C.touchwin((*C.WINDOW)(w))
}

// VLine draws a verticle line starting at y, x and ending at height using 
// the specified character
func (w *Window) VLine(y, x, ch, h int) {
	C.mvwvline((*C.WINDOW)(w), C.int(y), C.int(x), C.chtype(ch),
		C.int(wid))
	return
}