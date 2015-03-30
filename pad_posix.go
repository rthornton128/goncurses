// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// #include <curses.h>
import "C"

// Echo prints a single character to the pad immediately. This has the
// same effect of calling AddChar() + Refresh() but has a significant
// speed advantage
func (p *Pad) Echo(ch int) {
	C.pechochar(p.win, C.chtype(ch))
}
