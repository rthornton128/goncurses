// goncurses - ncurses library for Go.
//
// Copyright (c) 2011, Rob Thornton 
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without 
// modification, are permitted provided that the following conditions are met:
//
//   * Redistributions of source code must retain the above copyright notice, 
//     this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright notice, 
//     this list of conditions and the following disclaimer in the documentation 
//     and/or other materials provided with the distribution.
//  
//   * Neither the name of the copyright holder nor the names of its 
//     contributors may be used to endorse or promote products derived from this 
//     software without specific prior written permission.
//      
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" 
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE 
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE 
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE 
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR 
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF 
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS 
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN 
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) 
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE 
// POSSIBILITY OF SUCH DAMAGE.

/* ncurses panel extension

The following functions have not been implemented because there are far more
effective manners by which they can be done in Go: set_panel_userptr() and
panel_userptr() */
package goncurses

// #cgo LDFLAGS: -lpanel
// #include <panel.h>
// #include <ncurses.h>
import "C"

import (
	"os"
)

type Panel C.PANEL

// Panel creates a new panel derived from the window, adding it to the 
// panel stack. The pointer to the original window can still be used to
// excute most window functions with the exception of Refresh(). Always
// use panel's Refresh() function.
func (w *Window) Panel() *Panel {
	p := (*Panel)(C.new_panel((*C.WINDOW)(w)))
	if p != nil {
		return p
	}
	return nil
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
	return (*Panel)(C.panel_above((*C.PANEL)(p)))
}

// Returns a pointer to the panel below in the stack or nil. Passing nil will
// return the bottom panel in the stack
func Below(p *Panel) *Panel {
	return (*Panel)(C.panel_above((*C.PANEL)(p)))
}

// Move the panel to the bottom of the stack.
func (p *Panel) Bottom() os.Error {
	if C.bottom_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to move panel to bottom of stack")
	}
	return nil
}

// Delete panel, removing from the stack. 
func (p *Panel) Delete() os.Error {
	if C.del_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to delete panel")
	}
	p = nil
	return nil
}

// Hidden returns true if panel is visible, false if not
func (p *Panel) Hidden() bool {
	return C.panel_hidden((*C.PANEL)(p)) == C.TRUE
}

// Hide the panel
func (p *Panel) Hide() os.Error {
	if C.hide_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to hide panel")
	}
	return nil
}

// Move the panel to the specified location. It is important to never use
// ncurses movement functions on the window governed by panel. Always use
// this function
func (p *Panel) Move(y, x int) os.Error {
	if C.move_panel((*C.PANEL)(p), C.int(y), C.int(x)) == C.ERR {
		return os.NewError("Failed to move panel")
	}
	return nil
}

// Replace panel's associated window with a new one.
func (p *Panel) Replace(w *Window) os.Error {
	if C.replace_panel((*C.PANEL)(p), (*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to replace window")
	}
	return nil
}

// Show the panel, if hidden, and place it on the top of the stack.
func (p *Panel) Show() os.Error {
	if C.show_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to show panel")
	}
	return nil
}

// Move panel to the top of the stack
func (p *Panel) Top() os.Error {
	if C.top_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to move panel to top of stack")
	}
	return nil
}

// Window returns the window governed by panel
func (p *Panel) Window() *Window {
	return (*Window)(C.panel_window((*C.PANEL)(p)))
}
