// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <curses.h>
// #include "goncurses.h"
import "C"

type Pad struct {
	*Window
}

// NewPad creates a window which is not restricted by the terminal's
// dimentions (unlike a Window)
func NewPad(lines, cols int) *Pad {
	return &Pad{&Window{C.newpad(C.int(lines), C.int(cols))}}
}

// NoutRefresh() refreshes the pad for drawing but the screen will not
// be updated until Update() has been called on the underlying window.
func (p *Pad) NoutRefresh(py, px, ty, tx, by, bx int) {
	C.pnoutrefresh(p.win, C.int(py), C.int(px), C.int(ty),
		C.int(tx), C.int(by), C.int(bx))
}

// Refresh the pad at location py, px using the rectangle specified by
// ty, tx, by, bx (bottom/top y/x)
func (p *Pad) Refresh(py, px, ty, tx, by, bx int) {
	C.prefresh(p.win, C.int(py), C.int(px), C.int(ty), C.int(tx),
		C.int(by), C.int(bx))
}

// Sub creates a sub-pad lines by columns in size
func (p *Pad) Sub(y, x, h, w int) Pad {
	return Pad{&Window{C.subpad(p.win, C.int(h), C.int(w), C.int(y),
		C.int(x))}}
}

// Window is a helper function for calling Window functions on a pad like
// Print(). Convention would be to use pad.Window().Print() rather than to
// cast the pad to a window with (*Window)(&pad).Print().
/*func (p Pad) Window() *Window {
	return (*Window)(&p)
}*/
