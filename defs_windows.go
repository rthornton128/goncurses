// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <curses.h>
import "C"

// Definitions for printed characters not found on most keyboards. Ideally,
// these would not be hard-coded as they are potentially different on
// different systems. However, some ncurses implementations seem to be
// heavily reliant on macros which prevent these definitions from being
// handled by cgo properly. If they don't work for you, you won't be able
// to use them until either a) the Go team works out a way to overcome this
// limitation in godefs/cgo or b) an alternative method is found. Work is
// being done to find a solution from the ncurses source code.
const (
	ACS_DEGREE   Char = C.ACS_DEGREE
	ACS_PLMINUS       = C.ACS_PLMINUS
	ACS_BOARD         = C.ACS_BOARD
	ACS_LANTERN       = C.ACS_LANTERN
	ACS_LRCORNER      = C.ACS_LRCORNER
	ACS_URCORNER      = C.ACS_URCORNER
	ACS_LLCORNER      = C.ACS_LLCORNER
	ACS_ULCORNER      = C.ACS_ULCORNER
	ACS_PLUS          = C.ACS_PLUS
	ACS_S1            = C.ACS_S1
	ACS_S3            = C.ACS_S3
	ACS_HLINE         = C.ACS_HLINE
	ACS_S7            = C.ACS_S7
	ACS_S9            = C.ACS_S9
	ACS_LTEE          = C.ACS_LTEE
	ACS_RTEE          = C.ACS_RTEE
	ACS_BTEE          = C.ACS_BTEE
	ACS_TTEE          = C.ACS_TTEE
	ACS_VLINE         = C.ACS_VLINE
	ACS_LEQUAL        = C.ACS_LEQUAL
	ACS_GEQUAL        = C.ACS_GEQUAL
	ACS_PI            = C.ACS_PI
	ACS_NEQUAL        = C.ACS_NEQUAL
	ACS_STERLING      = C.ACS_STERLING
	ACS_BULLET        = C.ACS_BULLET
	ACS_LARROW        = C.ACS_LARROW
	ACS_RARROW        = C.ACS_RARROW
	ACS_DARROW        = C.ACS_DARROW
	ACS_UARROW        = C.ACS_UARROW
	ACS_BLOCK         = C.ACS_BLOCK
	ACS_DIAMOND       = C.ACS_DIAMOND
	ACS_CKBOARD       = C.ACS_CKBOARD
)
