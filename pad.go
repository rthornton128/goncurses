// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <curses.h>
// #include "goncurses.h"
import "C"

import "errors"

type Pad struct {
	*Window
}

// NewPad creates a window which is not restricted by the terminal's
// dimentions (unlike a Window). Pads accept all functions which can be
// called on a window. It returns a pointer to a new Pad of h(eight) by
// w(idth).
func NewPad(h, w int) (*Pad, error) {
	p := C.newpad(C.int(h), C.int(w))
	if p == nil {
		return nil, errors.New("Failed to create pad")
	}
	return &Pad{&Window{p}}, nil
}

// NoutRefresh indicates that a section of the screen should be redrawn but
// does not update the phsyical screen until Update() is called. See
// Pad.Refresh() for details on the arguments and Window.NoutRefresh for
// more details on the workings of this function
func (p *Pad) NoutRefresh(py, px, sy, sx, h, w int) error {
	ok := C.pnoutrefresh(p.win, C.int(py), C.int(px), C.int(sy),
		C.int(sx), C.int(h), C.int(w))
	if ok != C.OK {
		return errors.New("Failed to refresh pad")
	}
	return nil
}

// Refresh will calculate how to update the physical screen in the most
// efficient manor and update it. See Window.Refresh for more details.
// The coordinates py, px specify the location on the pad from which the
// characters we want to display are located. sy and sx specify the location
// on the screen where this data should be displayed. h and w are the height
// and width of the rectangle to be displayed. The coodinates and the size
// of the rectangle must be contained within both the Pad's and Window's
// respective areas
func (p *Pad) Refresh(py, px, sy, sx, h, w int) error {
	if C.prefresh(p.win, C.int(py), C.int(px), C.int(sy), C.int(sx),
		C.int(h), C.int(w)) != C.OK {
		return errors.New("Failed to refresh pad")
	}
	return nil
}

// Sub creates a sub-pad h(eight) by w(idth) in size starting at the location
// y, x in the parent pad. Changes to a sub-pad will also change it's parent
func (p *Pad) Sub(y, x, h, w int) *Pad {
	return &Pad{&Window{C.subpad(p.win, C.int(h), C.int(w), C.int(y),
		C.int(x))}}
}
