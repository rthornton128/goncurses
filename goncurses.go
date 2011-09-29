/* 1. No functions which operate only on stdscr have been implemented because 
 * it makes little sense to do so in a Go implementation. Stdscr is treated the
 * same as any other window.
 * 
 * 2. Whenever possible, versions of ncurses functions which could potentially
 * have a buffer overflow, like the getstr() family of functions, have been
 * implemented. Instead, only the mvwgetnstr() and wgetnstr() can be used. */
 
package goncurses

// #cgo LDFLAGS: -lncurses
// #include <ncurses.h>
// #include <stdlib.h>
import "C"

import (
    "fmt"
    "os"
    "unsafe"
)

type Attribute string

var attrList = map[Attribute] C.int {
    "normal": C.A_NORMAL,
    "standout": C.A_STANDOUT,
    "underline": C.A_UNDERLINE,
    "reverse": C.A_REVERSE,
    "blink": C.A_BLINK,
    "dim": C.A_DIM,
    "bold": C.A_BOLD,
    "protect": C.A_PROTECT,
    "invis": C.A_INVIS,
    "altcharset": C.A_ALTCHARSET,
    "chartext": C.A_CHARTEXT,
}

type Chtype C.chtype
type Color C.int

var colorList = map[string] Color {
    "black": C.COLOR_BLACK,
    "red": C.COLOR_RED,
    "green": C.COLOR_GREEN,
    "yellow": C.COLOR_YELLOW,
    "blue": C.COLOR_BLUE,
    "magenta": C.COLOR_MAGENTA,
    "cyan": C.COLOR_CYAN,
    "white": C.COLOR_WHITE,
}

var keyList = map[C.int] string {
    10: "enter", // On some keyboards?
    C.KEY_DOWN: "down",
    C.KEY_UP: "up",
    C.KEY_LEFT: "left",
    C.KEY_RIGHT: "right",
    C.KEY_HOME: "home",
    C.KEY_BACKSPACE: "backspace",
    C.KEY_F0: "F0",
    C.KEY_F0+1: "F1",
    C.KEY_F0+2: "F2",
    C.KEY_F0+3: "F3",
    C.KEY_F0+4: "F4",
    C.KEY_F0+5: "F5",
    C.KEY_F0+6: "F6",
    C.KEY_F0+7: "F7",
    C.KEY_F0+8: "F8",
    C.KEY_F0+9: "F9",
    C.KEY_F0+10: "F10",
    C.KEY_F0+11: "F11",
    C.KEY_F0+12: "F12",
    C.KEY_ENTER: "enter", // And not others?
    C.KEY_MOUSE: "mouse",
}

type MMask C.mmask_t

var mouseEvents = map[string] MMask {
    "button1-pressed": C.BUTTON1_PRESSED,
    "button1-released": C.BUTTON1_RELEASED,
    "button1-clicked": C.BUTTON1_CLICKED,
    "button1-double-clicked": C.BUTTON1_DOUBLE_CLICKED,
    "button1-triple-clicked": C.BUTTON1_TRIPLE_CLICKED,
    "button2-pressed": C.BUTTON2_PRESSED,
    "button2-released": C.BUTTON2_RELEASED,
    "button2-clicked": C.BUTTON2_CLICKED,
    "button2-double-clicked": C.BUTTON2_DOUBLE_CLICKED,
    "button2-triple-clicked": C.BUTTON2_TRIPLE_CLICKED,
    "button3-pressed": C.BUTTON3_PRESSED,
    "button3-released": C.BUTTON3_RELEASED,
    "button3-clicked": C.BUTTON3_CLICKED,
    "button3-double-clicked": C.BUTTON3_DOUBLE_CLICKED,
    "button3-triple-clicked": C.BUTTON3_TRIPLE_CLICKED,
    "button4-pressed": C.BUTTON4_PRESSED,
    "button4-released": C.BUTTON4_RELEASED,
    "button4-clicked": C.BUTTON4_CLICKED,
    "button4-double-clicked": C.BUTTON4_DOUBLE_CLICKED,
    "button4-triple-clicked": C.BUTTON4_TRIPLE_CLICKED,
//    "button5-pressed": C.BUTTON5_PRESSED,
//    "button5-released": C.BUTTON5_RELEASED,
//    "button5-clicked": C.BUTTON5_CLICKED,
//    "button5-double-clicked": C.BUTTON5_DOUBLE_CLICKED,
//    "button5-triple-clicked": C.BUTTON5_TRIPLE_CLICKED,
    "shift": C.BUTTON_SHIFT,
    "ctrl": C.BUTTON_CTRL,
    "alt": C.BUTTON_ALT,
    "all": C.ALL_MOUSE_EVENTS,
    "position": C.REPORT_MOUSE_POSITION,
}

// Turn on/off buffering; raw user signals are passed to the program for
// handling. Overrides raw mode
func Cbreak() os.Error {
    if (C.cbreak() == C.ERR) {
        return os.NewError("Failed to enable cbreak mode")
    }
    return nil
}

func CursSet(vis byte) os.Error {
    if (C.curs_set(C.int(vis)) == C.ERR) {
        return os.NewError("Failed to enable cbreak mode")
    }
    return nil
}

// Turn on echoing characters to the terminal
func Echo() os.Error {
    if (C.echo() == C.ERR) {
        return os.NewError("Failed to enable character echoing")
    }
    return nil
}

// Must be called prior to exiting the program in order to make sure the
// terminal returns to normal operation
func Endwin() {
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
// exceeded after a call to Getch() has been made then Getch will return
// with an error
func Halfdelay(delay int) os.Error {
    var cerr C.int
    if delay > 0 {
        cerr = C.halfdelay(C.int(delay))
    }
    if (cerr == C.ERR) {
        return os.NewError("Unable to set echo mode")
    }
    return nil
}

func InitColor(color string, r, g, b int) os.Error {
    col, ok := colorList[color];
    if !ok {
        return os.NewError("Failed to set new color definition")
    }
    if C.init_color(C.short(col), C.short(r), C.short(g), C.short(b)) == C.ERR {
        return os.NewError("Failed to set new color definition")
    }
    return nil
}

func InitPair(pair byte, fg, bg string) os.Error {
    if pair == 0 || C.int(pair) > (C.COLOR_PAIRS-1) {
        return os.NewError("Invalid color pair selected")
    }
    fg_color, fg_ok := colorList[fg]
    bg_color, bg_ok := colorList[bg]
    if !fg_ok {
        return os.NewError("Invalid foreground color")
    }
    if !bg_ok {
        return os.NewError("Invalid foreground color")
    }
    if C.init_pair(C.short(pair), C.short(fg_color), C.short(bg_color)) == C.ERR {
        return os.NewError("Failed to init color pair")
    }
    return nil
}

// Initialize the ncurses library. You must run this function prior to any 
// other goncurses function in order for the library to work
func Initscr() (stdscr *Window, err os.Error) {
    stdscr = (*Window)(C.initscr())
    err = nil
    if (unsafe.Pointer(stdscr) == nil) {
        err = os.NewError("An error occurred initializing ncurses")
    }
    return
}

// Returns a string representing the value of input returned by Getch
func Key(k C.int) (key string) {
    var ok bool
    key, ok = keyList[k]
    if !ok {
        key = fmt.Sprintf("%c", int(k))
    }
    return
}

func MouseMask(masks ...string) {
    var mousemask MMask
    for _, mask := range masks {
        mousemask |= mouseEvents[mask]
    }
    C.mousemask((C.mmask_t)(mousemask), (*C.mmask_t)(unsafe.Pointer(nil)))
}

func NewWin(h, w, y, x int) (new *Window, err os.Error) {
    new = (*Window)(C.newwin(C.int(h), C.int(w), C.int(y), C.int(x)))
    if (unsafe.Pointer(new) == unsafe.Pointer(nil)) {
        err = os.NewError("Failed to create a new window")
    }
    return
}

func Nocbreak() os.Error {
    if (C.nocbreak() == C.ERR) {
        return os.NewError("Failed to disable cbreak mode")
    }
    return nil
}

func Noecho() os.Error {
    if (C.noecho() == C.ERR) {
        return os.NewError("Failed to disable character echoing")
    }
    return nil
}

func Noraw() os.Error {
    if C.noraw() == C.ERR {
        return os.NewError("Failed to disable raw mode")
    }
    return nil
}

// Turn off input buffering; user signals are disabled and the key strokes are
// passed directly to input
func Raw() os.Error {
    if (C.raw() == C.ERR) {
        return os.NewError("Failed to enable raw mode")
    }
    return nil
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

func (w *Window) Addch(ch Chtype, attributes ...Attribute) {
    var cattr C.int
    for _, attr := range attributes {
        cattr |= attrList[attr]
    }
    C.waddch((*C.WINDOW)(w), C.chtype(C.int(ch) | cattr))
}

// Turn off character attribute TODO: range through Attribute array
func (w *Window) Attroff(attrstr Attribute) (err os.Error) {
    attr, ok := attrList[attrstr]
    if !ok {
        err = os.NewError(fmt.Sprintf("Invalid attribute: ", attrstr))
    }
    if C.wattroff((*C.WINDOW)(w), attr) == C.ERR {
        err = os.NewError(fmt.Sprintf("Failed to unset attribute: %s", attrstr))
    }
    return
}

// Turn on character attribute TODO: range through Attribute array
func (w *Window) Attron(attrstr Attribute) (err os.Error) {
    attr, ok := attrList[attrstr]
    if !ok {
        err = os.NewError(fmt.Sprintf("Invalid attribute: ", attrstr))
    }
    if C.wattron((*C.WINDOW)(w), attr) == C.ERR {
        err = os.NewError(fmt.Sprintf("Failed to set attribute: %s", attrstr))
    }
    return
}

func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) os.Error {
    res := C.wborder((*C.WINDOW)(w), C.chtype(ls), C.chtype(rs), C.chtype(ts),
            C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl),
            C.chtype(br))
    if res == C.ERR {
        return os.NewError("Failed to draw box around window")
    }
    return nil
}

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
        return os.NewError("Failed to clear screen")
    }
    return nil
}

// Clear from the current cursor position, moving to the right, to the end 
// of the line
func (w *Window) ClearToEOL() os.Error {
    if C.wclrtoeol((*C.WINDOW)(w)) == C.ERR {
        return os.NewError("Failed to clear screen")
    }
    return nil
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

func (w *Window) DelWin() os.Error {
    if C.delwin((*C.WINDOW)(w)) == C.ERR {
        return os.NewError("Failed to enable delete window")
    }
    return nil
}

func (w *Window) Erase() os.Error {
        if C.werase((*C.WINDOW)(w)) == C.ERR {
        return os.NewError("Failed to erase ??")
    }
    return nil
}

// Get a character from standard input
func (w *Window) Getch() (ch C.int, err os.Error) {
    if ch = C.wgetch((*C.WINDOW)(w)); ch == C.ERR {
        err = os.NewError("Failed to retrieve character from input stream")
    }
    return
}

// Returns the maximum size of the Window. Note that it uses ncurses idiom
// of returning y then x.
func (w *Window) Getmaxyx() (int, int) {
    // This hack is necessary to make cgo happy
    return int(w._maxy+1), int(w._maxx+1)
}

// Reads at most 'n' characters entered by the user from the Window. Attempts
// to enter greater than 'n' characters will elicit a 'beep'
func (w *Window) Getnstr(n int) (string, os.Error) {
    cstr := make([]C.char, n)
    if C.wgetnstr((*C.WINDOW)(w), (*C.char)(&cstr[0]), C.int(n)) == C.ERR {
        return "", os.NewError("Failed to retrieve string from input stream")
    }
    return C.GoString(&cstr[0]), nil
}

// Returns the current cursor location of the Window. Note that it uses 
// ncurses idiom of returning y then x.
func (w *Window) Getyx() (int, int) {
    // This hack is necessary to make cgo happy
    return int(w._cury), int(w._curx)
}

// Turn on/off accepting keypad characters like the F1-F12 keys and the 
// arrow keys
func (w *Window)Keypad(keypad bool) os.Error {
    var err C.int
    if err = C.keypad((*C.WINDOW)(w), C.bool(keypad)); err == C.ERR {
        return os.NewError("Unable to set keypad mode")
    }
    return nil
}

// Move the cursor to the specified coordinates
func (w *Window) Move(y, x int) os.Error {
    if C.wmove((*C.WINDOW)(w), C.int(y), C.int(x)) == C.ERR {
        return os.NewError("Failed to move cursor")
    }
    return nil
}
    
func (w *Window) Mvprint(y, x int, f string, s ...interface{}) {
    str := fmt.Sprintf(f, s...)
    C.wmove((*C.WINDOW)(w), C.int(y), C.int(x))
    for _, ch := range(str) {
        w.Addch(Chtype(ch))
    }
}

// Prints a formated string to the window similar to fmt.Printf family of 
// functions.
func (w *Window) Print(f string, a ...interface{}) {
// Currently, cgo doesn't play nice with varargs functions so this hack is
// used as a work around
    str := fmt.Sprintf(f, a...)
    for _, ch := range(str) {
        w.Addch(Chtype(ch))
    }
}

func (w *Window) Refresh() {
    C.wrefresh((*C.WINDOW)(w))
}
