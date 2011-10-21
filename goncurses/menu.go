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

/* Menu library */
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

// Not currently used
type Errno int

var menuerrors = map[Errno]string{
	C.E_OK:              "The routine succeeded.",
	C.E_BAD_ARGUMENT:    "Routine detected an incorrect or out-of-range argument.",
	C.E_BAD_STATE:       "Routine was called from an initialization or termination function.",
	C.E_NO_MATCH:        "Character failed to match.",
	C.E_NO_ROOM:         "Menu is too large for its window.",
	C.E_NOT_CONNECTED:   "No items are connected to the menu.",
	C.E_NOT_POSTED:      "The menu has not been posted.",
	C.E_NOT_SELECTABLE:  "The designated item cannot be selected.",
	C.E_POSTED:          "The menu is already posted.",
	C.E_REQUEST_DENIED:  "The menu driver could not process the request.",
	C.E_SYSTEM_ERROR:    "System error occurred (see errno).",
	C.E_UNKNOWN_COMMAND: "The menu driver code saw an unknown request code.",
}

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

func (e Errno) String() string {
	return menuerrors[e]
}

// The strings in this mapping correspond to those found in the keyList mapping in
// ncurses.go. If changing any of these values, the corresponding value should be
// modified too.
var driveractions = map[string]C.int{
	"down":        C.REQ_DOWN_ITEM,
	"first":       C.REQ_FIRST_ITEM,
	"last":        C.REQ_LAST_ITEM,
	"left":        C.REQ_LEFT_ITEM,
	"next":        C.REQ_NEXT_ITEM,
	"page down":   C.REQ_SCR_DPAGE,
	"page up":     C.REQ_SCR_UPAGE,
	"prev":        C.REQ_PREV_ITEM,
	"right":       C.REQ_RIGHT_ITEM,
	"scroll down": C.REQ_SCR_DLINE,
	"scroll up":   C.REQ_SCR_ULINE,
	"toggle":      C.REQ_TOGGLE_ITEM,
	"up":          C.REQ_UP_ITEM,
}

type Menu C.MENU
type MenuItem C.ITEM

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

// Background sets the attributes of un-highlighted items in the menu
func (m *Menu) Background(ch int) {
	C.set_menu_back((*C.MENU)(m), C.chtype(ch))
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
func (m *Menu) Driver(action string) {
	if da, ok := driveractions[action]; ok {
		C.menu_driver((*C.MENU)(m), da)
	}
	return
}

// Foreground sets the attributes of highlighted items in the menu
func (m *Menu) Foreground(ch int) {
	C.set_menu_fore((*C.MENU)(m), C.chtype(ch))
}

// Format sets the menu format. See the O_* menu options.
func (m *Menu) Format(r, c int) os.Error {
	if res := C.set_menu_format((*C.MENU)(m), C.int(r), C.int(c)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// Free deallocates memory set aside for the menu. This must be called
// before exiting.
func (m *Menu) Free() os.Error {
	if res := C.free_menu((*C.MENU)(m)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// Grey sets the attributes of non-selectable items in the menu
func (m *Menu) Grey(ch int) {
	C.set_menu_grey((*C.MENU)(m), C.chtype(ch))
}

// Items will either set or return the items in the menu. When setting
// items you must make sure the prior menu items will be freed. Pass
// nil to get the items in the menu.
func (m *Menu) Items(items []*MenuItem) []*MenuItem {
	if items == nil {
		type ItemArray [10]*C.ITEM
		citems := C.menu_items((*C.MENU)(m))
		count := m.Count()
		mitems := make([]*MenuItem, count)
		for index := 0; index < count; index++ {
			mitems[index] = (*MenuItem)(C.menu_item_at(citems, C.int(index)))
		}
		return mitems
	}
	citems := make([]*C.ITEM, len(items)+1)
	for index, item := range items {
		citems[index] = (*C.ITEM)(item)
	}
	citems[len(items)] = nil
	C.set_menu_items((*C.MENU)(m), (**C.ITEM)(&citems[0]))
	return nil
}

// Mark sets the indicator for the currently selected menu item
func (m *Menu) Mark(mark string) os.Error {
	cmark := C.CString(mark)
	defer C.free(unsafe.Pointer(cmark))

	if res := C.set_menu_mark((*C.MENU)(m), cmark); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
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
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// PositionCursor sets the cursor over the currently selected menu item.
func (m *Menu) PositionCursor() {
	C.pos_menu_cursor((*C.MENU)(m))
}

// Post the menu, making it visible
func (m *Menu) Post() os.Error {
	if res := C.post_menu((*C.MENU)(m)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// SubWindow for the menu
func (m *Menu) SubWindow(sub *Window) os.Error {
	if res := C.set_menu_sub((*C.MENU)(m), (*C.WINDOW)(sub)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// UnPost the menu, effectively hiding it.
func (m *Menu) UnPost() os.Error {
	if res := C.unpost_menu((*C.MENU)(m)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

// Window container for the menu
func (m *Menu) Window(win *Window) os.Error {
	if res := C.set_menu_win((*C.MENU)(m), (*C.WINDOW)(win)); res != C.E_OK {
		return os.NewError(menuerrors[Errno(res)])
	}
	return nil
}

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
