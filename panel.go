// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #cgo !darwin,!openbsd,!windows pkg-config: panel
// #cgo darwin openbsd LDFLAGS: -lpanel
// #include <panel.h>
// #include <curses.h>
import "C"

import "errors"

type Panel struct {
	pan *C.PANEL
}

// Panel creates a new panel derived from the window, adding it to the
// panel stack. The pointer to the original window can still be used to
// excute most window functions with the exception of Refresh(). Always
// use panel's Refresh() function.
func NewPanel(w *Window) *Panel {
	return &Panel{C.new_panel(w.win)}
}

// UpdatePanels refreshes the panel stack. It must be called prior to
// using ncurses's DoUpdate()
func UpdatePanels() {
	C.update_panels()
	return
}

// Returns a pointer to the panel above in the stack or nil. Passing nil will
// return the top panel in the stack
func (p *Panel) Above() *Panel {
	return &Panel{C.panel_above(p.pan)}
}

// Returns a pointer to the panel below in the stack or nil. Passing nil will
// return the bottom panel in the stack
func Below(p *Panel) *Panel {
	return &Panel{C.panel_above(p.pan)}
}

// Move the panel to the bottom of the stack.
func (p *Panel) Bottom() error {
	if C.bottom_panel(p.pan) == C.ERR {
		return errors.New("Failed to move panel to bottom of stack")
	}
	return nil
}

// Delete panel, removing from the stack.
func (p *Panel) Delete() error {
	if C.del_panel(p.pan) == C.ERR {
		return errors.New("Failed to delete panel")
	}
	p = nil
	return nil
}

// Hidden returns true if panel is visible, false if not
func (p *Panel) Hidden() bool {
	return C.panel_hidden(p.pan) == C.TRUE
}

// Hide the panel
func (p *Panel) Hide() error {
	if C.hide_panel(p.pan) == C.ERR {
		return errors.New("Failed to hide panel")
	}
	return nil
}

// Move the panel to the specified location. It is important to never use
// ncurses movement functions on the window governed by panel. Always use
// this function
func (p *Panel) Move(y, x int) error {
	if C.move_panel(p.pan, C.int(y), C.int(x)) == C.ERR {
		return errors.New("Failed to move panel")
	}
	return nil
}

// Replace panel's associated window with a new one.
func (p *Panel) Replace(w *Window) error {
	if C.replace_panel(p.pan, w.win) == C.ERR {
		return errors.New("Failed to replace window")
	}
	return nil
}

// Show the panel, if hidden, and place it on the top of the stack.
func (p *Panel) Show() error {
	if C.show_panel(p.pan) == C.ERR {
		return errors.New("Failed to show panel")
	}
	return nil
}

// Move panel to the top of the stack
func (p *Panel) Top() error {
	if C.top_panel(p.pan) == C.ERR {
		return errors.New("Failed to move panel to top of stack")
	}
	return nil
}

// Window returns the window governed by panel
func (p *Panel) Window() *Window {
	return &Window{C.panel_window(p.pan)}
}
