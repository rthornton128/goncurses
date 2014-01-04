// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// Definitions for printed characters not found on most keyboards.
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
