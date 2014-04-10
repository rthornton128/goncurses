// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// #include <curses.h>
import "C"

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
