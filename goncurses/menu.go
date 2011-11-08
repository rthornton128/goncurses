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

/* ncurses menu extension */
package goncurses
/*
#cgo LDFLAGS: -lmenu
#include <menu.h>
#include <stdlib.h>

ITEM* menu_item_at(ITEM** ilist, int i) {
	return ilist[i];
}*/
import "C"

import (
	"os"
	"unsafe"
)

// Menu Driver Requests
const (
	REQ_LEFT          = C.REQ_LEFT_ITEM
	REQ_RIGHT         = C.REQ_RIGHT_ITEM
	REQ_UP            = C.REQ_UP_ITEM
	REQ_DOWN          = C.REQ_DOWN_ITEM
	REQ_ULINE         = C.REQ_SCR_ULINE
	REQ_DLINE         = C.REQ_SCR_DLINE
	REQ_PAGE_DOWN     = C.REQ_SCR_DPAGE
	REQ_PAGE_UP       = C.REQ_SCR_UPAGE
	REQ_FIRST         = C.REQ_FIRST_ITEM
	REQ_LAST          = C.REQ_LAST_ITEM
	REQ_NEXT          = C.REQ_NEXT_ITEM
	REQ_PREV          = C.REQ_PREV_ITEM
	REQ_TOGGLE        = C.REQ_TOGGLE_ITEM
	REQ_CLEAR_PATTERN = C.REQ_CLEAR_PATTERN
	REQ_BACK_PATTERN  = C.REQ_BACK_PATTERN
	REQ_NEXT_MATCH    = C.REQ_NEXT_MATCH
	REQ_PREV_MATCH    = C.REQ_PREV_MATCH
)

// Menu Options
const (
	O_ONEVALUE   = C.O_ONEVALUE   // Only one item can be selected
	O_SHOWDESC   = C.O_SHOWDESC   // Display item descriptions
	O_ROWMAJOR   = C.O_ROWMAJOR   // Display in row-major order
	O_IGNORECASE = C.O_IGNORECASE // Ingore case when pattern-matching
	O_SHOWMATCH  = C.O_SHOWMATCH  // Move cursor to item when pattern-matching
	O_NONCYCLIC  = C.O_NONCYCLIC  // Don't wrap next/prev item
)

// Menu Item Options
const O_SELECTABLE = C.O_SELECTABLE

// DriverActions is a convenience mapping for common responses
// to keyboard input
var DriverActions = map[int]int{
	KEY_DOWN:     C.REQ_DOWN_ITEM,
	KEY_HOME:     C.REQ_FIRST_ITEM,
	KEY_END:      C.REQ_LAST_ITEM,
	KEY_LEFT:     C.REQ_LEFT_ITEM,
	KEY_PAGEDOWN: C.REQ_SCR_DPAGE,
	KEY_PAGEUP:   C.REQ_SCR_UPAGE,
	KEY_RIGHT:    C.REQ_RIGHT_ITEM,
	KEY_UP:       C.REQ_UP_ITEM,
}

type Menu C.MENU

// NewMenu returns a pointer to a new menu.
func NewMenu(items []*MenuItem) (*Menu, os.Error) {
	citems := make([]*C.ITEM, len(items)+1)
	for index, item := range items {
		citems[index] = (*C.ITEM)(item)
	}
	citems[len(items)] = nil
	menu, errno := C.new_menu((**C.ITEM)(&citems[0]))
	if menu == nil {
		return nil, errno
	}
	return (*Menu)(menu), nil
}

// Background returns the menu's background character setting
func (m *Menu) Background() int {
	return int(C.menu_back((*C.MENU)(m)))
}

// Count returns the number of MenuItems in the Menu
func (m *Menu) Count() int {
	return int(C.item_count((*C.MENU)(m)))
}

// Current returns the selected item in the menu
func (m *Menu) Current(mi *MenuItem) *MenuItem {
	if mi == nil {
		return (*MenuItem)(C.current_item((*C.MENU)(m)))
	}
	C.set_current_item((*C.MENU)(m), (*C.ITEM)(mi))
	return nil
}

// Driver controls how the menu is activated. Action usually corresponds
// to the string return by the Key() function in goncurses.
func (m *Menu) Driver(daction int) os.Error {
	if err := C.menu_driver((*C.MENU)(m), C.int(daction)); err != C.E_OK {
		return error(os.Errno(err))
	}
	return nil
}

// Foreground gets the attributes of highlighted items in the menu
func (m *Menu) Foreground() int {
	return int(C.menu_fore((*C.MENU)(m)))
}

// Format sets the menu format. See the O_* menu options.
func (m *Menu) Format(r, c int) os.Error {
	if res := C.set_menu_format((*C.MENU)(m), C.int(r), C.int(c)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// Free deallocates memory set aside for the menu. This must be called
// before exiting.
func (m *Menu) Free() os.Error {
	if res := C.free_menu((*C.MENU)(m)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// Grey sets the attributes of non-selectable items in the menu
func (m *Menu) Grey(ch int) {
	C.set_menu_grey((*C.MENU)(m), C.chtype(ch))
}

// Items will return the items in the menu.
func (m *Menu) Items() []*MenuItem {
	citems := C.menu_items((*C.MENU)(m))
	count := m.Count()
	mitems := make([]*MenuItem, count)
	for index := 0; index < count; index++ {
		mitems[index] = (*MenuItem)(C.menu_item_at(citems, C.int(index)))
	}
	return mitems
}

// Mark sets the indicator for the currently selected menu item
func (m *Menu) Mark(mark string) os.Error {
	cmark := C.CString(mark)
	defer C.free(unsafe.Pointer(cmark))

	if res := C.set_menu_mark((*C.MENU)(m), cmark); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// Option sets the options for the menu. See the O_* definitions for
// a list of values which can be OR'd together
func (m *Menu) Option(opts int, on bool) os.Error {
	var res C.int
	if on {
		res = C.menu_opts_on((*C.MENU)(m), C.Menu_Options(opts))
	} else {
		res = C.menu_opts_off((*C.MENU)(m), C.Menu_Options(opts))
	}
	if res != 0 {
		return error(os.Errno(res))
	}
	return nil
}

// Pad sets the padding character for menu items.
func (m *Menu) Pad() int {
	return int(C.menu_pad((*C.MENU)(m)))
}

// Pattern returns the menu's pattern buffer
func (m *Menu) Pattern() string {
	return C.GoString(C.menu_pattern((*C.MENU)(m)))
}

// PositionCursor sets the cursor over the currently selected menu item.
func (m *Menu) PositionCursor() {
	C.pos_menu_cursor((*C.MENU)(m))
}

// Post the menu, making it visible
func (m *Menu) Post() os.Error {
	if res := C.post_menu((*C.MENU)(m)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// SetBackground set the attributes of the un-highlighted items in the 
// menu
func (m *Menu) SetBackground(ch int) os.Error {
	err := C.set_menu_back((*C.MENU)(m), C.chtype(ch))
	if err != C.E_OK {
		return error(os.Errno(err))
	}
	return nil
}

// SetForeground sets the attributes of the highlighted items in the menu
func (m *Menu) SetForeground(ch int) os.Error {
	err := C.set_menu_fore((*C.MENU)(m), C.chtype(ch))
	if err != C.E_OK {
		return error(os.Errno(err))
	}
	return nil
}

// SetItems will either set the items in the menu. When setting
// items you must make sure the prior menu items will be freed.
func (m *Menu) SetItems(items []*MenuItem) os.Error {
	citems := make([]*C.ITEM, len(items)+1)
	for index, item := range items {
		citems[index] = (*C.ITEM)(item)
	}
	citems[len(items)] = nil
	err := C.set_menu_items((*C.MENU)(m), (**C.ITEM)(&citems[0]))
	if err != C.E_OK {
		return error(os.Errno(err))
	}
	return nil
}

// SetPad sets the padding character for menu items.
func (m *Menu) SetPad(ch int) os.Error {
	if res := C.set_menu_pad((*C.MENU)(m), C.int(ch)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// SetPattern sets the padding character for menu items.
func (m *Menu) SetPattern(pattern string) os.Error {
	cpattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cpattern))
	res := C.set_menu_pattern((*C.MENU)(m), (*C.char)(cpattern))
	if res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// SetSpacing of the the menu's items. 'desc' is the space between the
// item and it's description andmay not be larger than TAB_SIZE. 'row' 
// is the number of rows separating each item and may not be larger than 
// three. 'col' is the spacing between each column of items in 
// multi-column mode. Use values of 0 or 1 to reset spacing to default, 
// which is one
func (m *Menu) SetSpacing(desc, row, col int) os.Error {
	err := C.set_menu_spacing((*C.MENU)(m), C.int(desc), C.int(row),
		C.int(col))
	if err != C.E_OK {
		return error(os.Errno(err))
	}
	return nil
}

// Spacing returns the menu item spacing. See SetSpacing for a description
func (m *Menu) Spacing() (int, int, int) {
	var desc, row, col C.int
	err := C.menu_spacing((*C.MENU)(m), (*C.int)(&desc), (*C.int)(&row),
		(*C.int)(&col))
	if err != C.E_OK {
		return int(desc), int(row), int(col)
	}
	return 0, 0, 0
}

// SubWindow for the menu
func (m *Menu) SubWindow(sub *Window) os.Error {
	if res := C.set_menu_sub((*C.MENU)(m), (*C.WINDOW)(sub)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// UnPost the menu, effectively hiding it.
func (m *Menu) UnPost() os.Error {
	if res := C.unpost_menu((*C.MENU)(m)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

// Window container for the menu
func (m *Menu) Window(win *Window) os.Error {
	if res := C.set_menu_win((*C.MENU)(m), (*C.WINDOW)(win)); res != C.E_OK {
		return error(os.Errno(res))
	}
	return nil
}

type MenuItem C.ITEM

// NewItem creates a new menu item with name and description.
func NewItem(name, desc string) (*MenuItem, os.Error) {
	cname := C.CString(name)
	cdesc := C.CString(desc)

	item, err := C.new_item(cname, cdesc)
	if item == nil {
		return nil, err
	}
	return (*MenuItem)(item), nil
}

// Description returns the second value passed to NewItem 
func (mi *MenuItem) Description() string {
	return C.GoString(C.item_description((*C.ITEM)(mi)))
}

// Free must be called on all menu items to avoid memory leaks
func (mi *MenuItem) Free() {
	C.free(unsafe.Pointer(C.item_name((*C.ITEM)(mi))))
	C.free_item((*C.ITEM)(mi))
}

// Name of the menu item
func (mi *MenuItem) Name() string {
	return C.GoString(C.item_name((*C.ITEM)(mi)))
}

// Selectable turns on/off whether a menu option is "greyed out"
func (mi *MenuItem) Selectable(on bool) {
	if on {
		C.item_opts_on((*C.ITEM)(mi), O_SELECTABLE)
	} else {
		C.item_opts_off((*C.ITEM)(mi), O_SELECTABLE)
	}
}

// Value returns true if menu item is toggled/active, otherwise false
func (mi *MenuItem) Value() bool {
	return bool(C.item_value((*C.ITEM)(mi)))
}
