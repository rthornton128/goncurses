// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// Definitions for printed characters not found on most keyboards. Ideally,
// these would not be hard-coded as they are potentially different on
// different systems. However, some ncurses implementations seem to be
// heavily reliant on macros which prevent these definitions from being
// handled by cgo properly. If they don't work for you, you won't be able
// to use them until either a) the Go team works out a way to overcome this
// limitation in godefs/cgo or b) an alternative method is found. Work is
// being done to find a solution from the ncurses source code.
const (
	ACS_DEGREE Char = iota + 4194406
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
