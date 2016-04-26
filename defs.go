// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: ncurses
// #include <curses.h>
import "C"

// Synconize options for Sync() function
const (
	SYNC_NONE   = iota
	SYNC_CURSOR // Sync cursor in all sub/derived windows
	SYNC_DOWN   // Sync changes in all parent windows
	SYNC_UP     // Sync change in all child windows
)

type Char C.chtype

// Text attributes
const (
	A_NORMAL     Char = C.A_NORMAL
	A_STANDOUT        = C.A_STANDOUT
	A_UNDERLINE       = C.A_UNDERLINE
	A_REVERSE         = C.A_REVERSE
	A_BLINK           = C.A_BLINK
	A_DIM             = C.A_DIM
	A_BOLD            = C.A_BOLD
	A_PROTECT         = C.A_PROTECT
	A_INVIS           = C.A_INVIS
	A_ALTCHARSET      = C.A_ALTCHARSET
	A_CHARTEXT        = C.A_CHARTEXT
)

var attrList = map[C.int]string{
	C.A_NORMAL:    "normal",
	C.A_STANDOUT:  "standout",
	C.A_UNDERLINE: "underline",
	C.A_REVERSE:   "reverse",
	C.A_BLINK:     "blink",
	//C.A_DIM:        "dim", TODO: pdcurses only, distinct in ncurses
	C.A_BOLD:       "bold",
	C.A_PROTECT:    "protect",
	C.A_INVIS:      "invis",
	C.A_ALTCHARSET: "altcharset",
	C.A_CHARTEXT:   "chartext",
}

// Definitions for printed characters not found on most keyboards.
const (
	/* VT100 symbols */
	ACS_ULCORNER Char = C.A_ALTCHARSET + 'l'
	ACS_LLCORNER      = C.A_ALTCHARSET + 'm'
	ACS_URCORNER      = C.A_ALTCHARSET + 'k'
	ACS_LRCORNER      = C.A_ALTCHARSET + 'j'
	ACS_LTEE          = C.A_ALTCHARSET + 't'
	ACS_RTEE          = C.A_ALTCHARSET + 'u'
	ACS_BTEE          = C.A_ALTCHARSET + 'v'
	ACS_TTEE          = C.A_ALTCHARSET + 'w'
	ACS_HLINE         = C.A_ALTCHARSET + 'q'
	ACS_VLINE         = C.A_ALTCHARSET + 'x'
	ACS_PLUS          = C.A_ALTCHARSET + 'n'
	ACS_S1            = C.A_ALTCHARSET + 'o'
	ACS_S9            = C.A_ALTCHARSET + 's'
	ACS_DIAMOND       = C.A_ALTCHARSET + '`'
	ACS_CKBOARD       = C.A_ALTCHARSET + 'a'
	ACS_DEGREE        = C.A_ALTCHARSET + 'f'
	ACS_PLMINUS       = C.A_ALTCHARSET + 'g'
	ACS_BULLET        = C.A_ALTCHARSET + '~'

	/* Teletype 5410v1 symbols */
	ACS_LARROW  = C.A_ALTCHARSET + ','
	ACS_RARROW  = C.A_ALTCHARSET + '+'
	ACS_DARROW  = C.A_ALTCHARSET + '.'
	ACS_UARROW  = C.A_ALTCHARSET + '-'
	ACS_BOARD   = C.A_ALTCHARSET + 'h'
	ACS_LANTERN = C.A_ALTCHARSET + 'i'
	ACS_BLOCK   = C.A_ALTCHARSET + '0'

	/* Undocumented, not well supported */
	ACS_S3       = C.A_ALTCHARSET + 'p'
	ACS_S7       = C.A_ALTCHARSET + 'r'
	ACS_LEQUAL   = C.A_ALTCHARSET + 'y'
	ACS_GEQUAL   = C.A_ALTCHARSET + 'z'
	ACS_PI       = C.A_ALTCHARSET + '{'
	ACS_NEQUAL   = C.A_ALTCHARSET + '|'
	ACS_STERLING = C.A_ALTCHARSET + '}'
)

// Colors available to ncurses. Combine these with the dim/bold attributes
// for bright/dark versions of each color. These colors can be used for
// both background and foreground colors.
const (
	C_BLACK   int16 = C.COLOR_BLACK
	C_BLUE          = C.COLOR_BLUE
	C_CYAN          = C.COLOR_CYAN
	C_GREEN         = C.COLOR_GREEN
	C_MAGENTA       = C.COLOR_MAGENTA
	C_RED           = C.COLOR_RED
	C_WHITE         = C.COLOR_WHITE
	C_YELLOW        = C.COLOR_YELLOW
)

type Key int

const (
	KEY_TAB       Key = 9               // tab
	KEY_RETURN        = 10              // enter key vs. KEY_ENTER
	KEY_DOWN          = C.KEY_DOWN      // down arrow key
	KEY_UP            = C.KEY_UP        // up arrow key
	KEY_LEFT          = C.KEY_LEFT      // left arrow key
	KEY_RIGHT         = C.KEY_RIGHT     // right arrow key
	KEY_HOME          = C.KEY_HOME      // home key
	KEY_BACKSPACE     = C.KEY_BACKSPACE // backpace
	KEY_F1            = C.KEY_F0 + 1    // F1 key
	KEY_F2            = C.KEY_F0 + 2    // F2 key
	KEY_F3            = C.KEY_F0 + 3    // F3 key
	KEY_F4            = C.KEY_F0 + 4    // F4 key
	KEY_F5            = C.KEY_F0 + 5    // F5 key
	KEY_F6            = C.KEY_F0 + 6    // F6 key
	KEY_F7            = C.KEY_F0 + 7    // F7 key
	KEY_F8            = C.KEY_F0 + 8    // F8 key
	KEY_F9            = C.KEY_F0 + 9    // F9 key
	KEY_F10           = C.KEY_F0 + 10   // F10 key
	KEY_F11           = C.KEY_F0 + 11   // F11 key
	KEY_F12           = C.KEY_F0 + 12   // F12 key
	KEY_DL            = C.KEY_DL        // delete-line key
	KEY_IL            = C.KEY_IL        // insert-line key
	KEY_DC            = C.KEY_DC        // delete-character key
	KEY_IC            = C.KEY_IC        // insert-character key
	KEY_EIC           = C.KEY_EIC       // sent by rmir or smir in insert mode
	KEY_CLEAR         = C.KEY_CLEAR     // clear-screen or erase key
	KEY_EOS           = C.KEY_EOS       // clear-to-end-of-screen key
	KEY_EOL           = C.KEY_EOL       // clear-to-end-of-line key
	KEY_SF            = C.KEY_SF        // scroll-forward key
	KEY_SR            = C.KEY_SR        // scroll-backward key
	KEY_PAGEDOWN      = C.KEY_NPAGE     // page-down key (next-page)
	KEY_PAGEUP        = C.KEY_PPAGE     // page-up key (prev-page)
	KEY_STAB          = C.KEY_STAB      // set-tab key
	KEY_CTAB          = C.KEY_CTAB      // clear-tab key
	KEY_CATAB         = C.KEY_CATAB     // clear-all-tabs key
	KEY_ENTER         = C.KEY_ENTER     // enter/send key
	KEY_PRINT         = C.KEY_PRINT     // print key
	KEY_LL            = C.KEY_LL        // lower-left key (home down)
	KEY_A1            = C.KEY_A1        // upper left of keypad
	KEY_A3            = C.KEY_A3        // upper right of keypad
	KEY_B2            = C.KEY_B2        // center of keypad
	KEY_C1            = C.KEY_C1        // lower left of keypad
	KEY_C3            = C.KEY_C3        // lower right of keypad
	KEY_BTAB          = C.KEY_BTAB      // back-tab key
	KEY_BEG           = C.KEY_BEG       // begin key
	KEY_CANCEL        = C.KEY_CANCEL    // cancel key
	KEY_CLOSE         = C.KEY_CLOSE     // close key
	KEY_COMMAND       = C.KEY_COMMAND   // command key
	KEY_COPY          = C.KEY_COPY      // copy key
	KEY_CREATE        = C.KEY_CREATE    // create key
	KEY_END           = C.KEY_END       // end key
	KEY_EXIT          = C.KEY_EXIT      // exit key
	KEY_FIND          = C.KEY_FIND      // find key
	KEY_HELP          = C.KEY_HELP      // help key
	KEY_MARK          = C.KEY_MARK      // mark key
	KEY_MESSAGE       = C.KEY_MESSAGE   // message key
	KEY_MOVE          = C.KEY_MOVE      // move key
	KEY_NEXT          = C.KEY_NEXT      // next key
	KEY_OPEN          = C.KEY_OPEN      // open key
	KEY_OPTIONS       = C.KEY_OPTIONS   // options key
	KEY_PREVIOUS      = C.KEY_PREVIOUS  // previous key
	KEY_REDO          = C.KEY_REDO      // redo key
	KEY_REFERENCE     = C.KEY_REFERENCE // reference key
	KEY_REFRESH       = C.KEY_REFRESH   // refresh key
	KEY_REPLACE       = C.KEY_REPLACE   // replace key
	KEY_RESTART       = C.KEY_RESTART   // restart key
	KEY_RESUME        = C.KEY_RESUME    // resume key
	KEY_SAVE          = C.KEY_SAVE      // save key
	KEY_SBEG          = C.KEY_SBEG      // shifted begin key
	KEY_SCANCEL       = C.KEY_SCANCEL   // shifted cancel key
	KEY_SCOMMAND      = C.KEY_SCOMMAND  // shifted command key
	KEY_SCOPY         = C.KEY_SCOPY     // shifted copy key
	KEY_SCREATE       = C.KEY_SCREATE   // shifted create key
	KEY_SDC           = C.KEY_SDC       // shifted delete-character key
	KEY_SDL           = C.KEY_SDL       // shifted delete-line key
	KEY_SELECT        = C.KEY_SELECT    // select key
	KEY_SEND          = C.KEY_SEND      // shifted end key
	KEY_SEOL          = C.KEY_SEOL      // shifted clear-to-end-of-line key
	KEY_SEXIT         = C.KEY_SEXIT     // shifted exit key
	KEY_SFIND         = C.KEY_SFIND     // shifted find key
	KEY_SHELP         = C.KEY_SHELP     // shifted help key
	KEY_SHOME         = C.KEY_SHOME     // shifted home key
	KEY_SIC           = C.KEY_SIC       // shifted insert-character key
	KEY_SLEFT         = C.KEY_SLEFT     // shifted left-arrow key
	KEY_SMESSAGE      = C.KEY_SMESSAGE  // shifted message key
	KEY_SMOVE         = C.KEY_SMOVE     // shifted move key
	KEY_SNEXT         = C.KEY_SNEXT     // shifted next key
	KEY_SOPTIONS      = C.KEY_SOPTIONS  // shifted options key
	KEY_SPREVIOUS     = C.KEY_SPREVIOUS // shifted previous key
	KEY_SPRINT        = C.KEY_SPRINT    // shifted print key
	KEY_SREDO         = C.KEY_SREDO     // shifted redo key
	KEY_SREPLACE      = C.KEY_SREPLACE  // shifted replace key
	KEY_SRIGHT        = C.KEY_SRIGHT    // shifted right-arrow key
	KEY_SRSUME        = C.KEY_SRSUME    // shifted resume key
	KEY_SSAVE         = C.KEY_SSAVE     // shifted save key
	KEY_SSUSPEND      = C.KEY_SSUSPEND  // shifted suspend key
	KEY_SUNDO         = C.KEY_SUNDO     // shifted undo key
	KEY_SUSPEND       = C.KEY_SUSPEND   // suspend key
	KEY_UNDO          = C.KEY_UNDO      // undo key
	KEY_MOUSE         = C.KEY_MOUSE     // any mouse event
	KEY_RESIZE        = C.KEY_RESIZE    // Terminal resize event
	//KEY_EVENT         = C.KEY_EVENT     // We were interrupted by an event
	KEY_MAX = C.KEY_MAX // Maximum key value is KEY_EVENT (0633)
)

var keyList = map[Key]string{
	KEY_TAB:       "tab",
	KEY_RETURN:    "enter", // On some keyboards?
	KEY_DOWN:      "down",
	KEY_UP:        "up",
	KEY_LEFT:      "left",
	KEY_RIGHT:     "right",
	KEY_HOME:      "home",
	KEY_BACKSPACE: "backspace",
	KEY_ENTER:     "enter", // And not others?
	KEY_F1:        "F1",
	KEY_F2:        "F2",
	KEY_F3:        "F3",
	KEY_F4:        "F4",
	KEY_F5:        "F5",
	KEY_F6:        "F6",
	KEY_F7:        "F7",
	KEY_F8:        "F8",
	KEY_F9:        "F9",
	KEY_F10:       "F10",
	KEY_F11:       "F11",
	KEY_F12:       "F12",
	KEY_MOUSE:     "mouse",
	KEY_PAGEUP:    "page up",
	KEY_PAGEDOWN:  "page down",
}

type MouseButton int

// Mouse button events
const (
	M_ALL            MouseButton = C.ALL_MOUSE_EVENTS
	M_ALT                        = C.BUTTON_ALT      // alt-click
	M_B1_PRESSED                 = C.BUTTON1_PRESSED // button 1
	M_B1_RELEASED                = C.BUTTON1_RELEASED
	M_B1_CLICKED                 = C.BUTTON1_CLICKED
	M_B1_DBL_CLICKED             = C.BUTTON1_DOUBLE_CLICKED
	M_B1_TPL_CLICKED             = C.BUTTON1_TRIPLE_CLICKED
	M_B2_PRESSED                 = C.BUTTON2_PRESSED // button 2
	M_B2_RELEASED                = C.BUTTON2_RELEASED
	M_B2_CLICKED                 = C.BUTTON2_CLICKED
	M_B2_DBL_CLICKED             = C.BUTTON2_DOUBLE_CLICKED
	M_B2_TPL_CLICKED             = C.BUTTON2_TRIPLE_CLICKED
	M_B3_PRESSED                 = C.BUTTON3_PRESSED // button 3
	M_B3_RELEASED                = C.BUTTON3_RELEASED
	M_B3_CLICKED                 = C.BUTTON3_CLICKED
	M_B3_DBL_CLICKED             = C.BUTTON3_DOUBLE_CLICKED
	M_B3_TPL_CLICKED             = C.BUTTON3_TRIPLE_CLICKED
	M_B4_PRESSED                 = C.BUTTON4_PRESSED // button 4
	M_B4_RELEASED                = C.BUTTON4_RELEASED
	M_B4_CLICKED                 = C.BUTTON4_CLICKED
	M_B4_DBL_CLICKED             = C.BUTTON4_DOUBLE_CLICKED
	M_B4_TPL_CLICKED             = C.BUTTON4_TRIPLE_CLICKED
	M_CTRL                       = C.BUTTON_CTRL           // ctrl-click
	M_SHIFT                      = C.BUTTON_SHIFT          // shift-click
	M_POSITION                   = C.REPORT_MOUSE_POSITION // mouse moved
)
