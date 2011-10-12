/* Panels implementation

The following functions have not been implemented because there are far more
effective manners by which they can be done in Go: set_panel_userptr() and
panel_userptr()
*/

package goncurses

// #cgo LDFLAGS: -lpanel
// #include <panel.h>
// #include <ncurses.h>
import "C"

import (
	"os"
)

type Panel C.PANEL

// Delete panel, removing from the stack. 
func DeletePanel(p *Panel) os.Error {
	if C.del_panel((*C.PANEL)(p)) == C.ERR {
		return os.NewError("Failed to delete panel")
	}
	return nil
}

// Add a window to the panel stack
func NewPanel(w *Window) *Panel {
	p := (*Panel)(C.new_panel((*C.WINDOW)(w)))
	if p != nil {
		return p
	}
	return nil
}

// Update the panel stack. Must be called prior to using ncurses's DoUpdate()
// Never use Refresh() with the panel library
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

// Returns true if panel is visible, false if not
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

// Replace panel's window with a new one.
func (p *Panel) Replace(w *Window) os.Error {
	if C.replace_panel((*C.PANEL)(p), (*C.WINDOW)(w)) == C.ERR {
		return os.NewError("Failed to replace window")
	}
	return nil
}

// Make the window visible. Also places it on the top of the stack.
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

// Return the window governed by panel
func (p *Panel) Window() *Window {
	return (*Window)(C.panel_window((*C.PANEL)(p)))
}
