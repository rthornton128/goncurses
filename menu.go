// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

/*
#cgo !darwin,!openbsd pkg-config: menu
#cgo darwin openbsd LDFLAGS: -lmenu
#include <menu.h>
#include <stdlib.h>

ITEM* menu_item_at(ITEM** ilist, int i) {
	return ilist[i];
}*/
import "C"

import (
	"syscall"
	"unsafe"
)

type Menu struct {
	menu *C.MENU
}

type MenuItem struct {
	item *C.ITEM
}

// NewMenu returns a pointer to a new menu.
func NewMenu(items []*MenuItem) (*Menu, error) {
	citems := make([]*C.ITEM, len(items)+1)
	for index, item := range items {
		citems[index] = item.item
	}
	citems[len(items)] = nil
	var menu *C.MENU
	var err error
	menu, err = C.new_menu((**C.ITEM)(&citems[0]))
	return &Menu{menu}, ncursesError(err)
}

// RequestName of menu request code
func RequestName(request int) (string, error) {
	cstr, err := C.menu_request_name(C.int(request))
	return C.GoString(cstr), ncursesError(err)
}

// RequestByName returns the request ID of the provide request
func RequestByName(request string) (res int, err error) {
	cstr := C.CString(request)
	defer C.free(unsafe.Pointer(cstr))

	res = int(C.menu_request_by_name(cstr))
	err = ncursesError(syscall.Errno(res))
	return
}

// Background returns the menu's background character setting
func (m *Menu) Background() int {
	return int(C.menu_back(m.menu))
}

// Count returns the number of MenuItems in the Menu
func (m *Menu) Count() int {
	return int(C.item_count(m.menu))
}

// Current returns the selected item in the menu
func (m *Menu) Current(mi *MenuItem) *MenuItem {
	if mi == nil {
		return &MenuItem{C.current_item(m.menu)}
	}
	C.set_current_item(m.menu, mi.item)
	return nil
}

// Driver controls how the menu is activated. Action usually corresponds
// to the string return by the Key() function in goncurses.
func (m *Menu) Driver(daction int) error {
	err := C.menu_driver(m.menu, C.int(daction))
	return ncursesError(syscall.Errno(err))
}

// Foreground gets the attributes of highlighted items in the menu
func (m *Menu) Foreground() int {
	return int(C.menu_fore(m.menu))
}

// Format sets the menu format. See the O_* menu options.
func (m *Menu) Format(r, c int) error {
	err := C.set_menu_format(m.menu, C.int(r), C.int(c))
	return ncursesError(syscall.Errno(err))
}

// Free deallocates memory set aside for the menu. This must be called
// before exiting.
func (m *Menu) Free() error {
	err := C.free_menu(m.menu)
	m = nil
	return ncursesError(syscall.Errno(err))
}

// Grey sets the attributes of non-selectable items in the menu
func (m *Menu) Grey(ch Char) {
	C.set_menu_grey(m.menu, C.chtype(ch))
}

// Items will return the items in the menu.
func (m *Menu) Items() []*MenuItem {
	citems := C.menu_items(m.menu)
	count := m.Count()
	mitems := make([]*MenuItem, count)
	for index := 0; index < count; index++ {
		mitems[index] = &MenuItem{C.menu_item_at(citems, C.int(index))}
	}
	return mitems
}

// Mark sets the indicator for the currently selected menu item
func (m *Menu) Mark(mark string) error {
	cmark := C.CString(mark)
	defer C.free(unsafe.Pointer(cmark))

	err := C.set_menu_mark(m.menu, cmark)
	return ncursesError(syscall.Errno(err))
}

// Option sets the options for the menu. See the O_* definitions for
// a list of values which can be OR'd together
func (m *Menu) Option(opts int, on bool) error {
	var err C.int
	if on {
		err = C.menu_opts_on(m.menu, C.Menu_Options(opts))
	} else {
		err = C.menu_opts_off(m.menu, C.Menu_Options(opts))
	}
	return ncursesError(syscall.Errno(err))
}

// Pad sets the padding character for menu items.
func (m *Menu) Pad() int {
	return int(C.menu_pad(m.menu))
}

// Pattern returns the menu's pattern buffer
func (m *Menu) Pattern() string {
	return C.GoString(C.menu_pattern(m.menu))
}

// PositionCursor sets the cursor over the currently selected menu item.
func (m *Menu) PositionCursor() {
	C.pos_menu_cursor(m.menu)
}

// Post the menu, making it visible
func (m *Menu) Post() error {
	err := C.post_menu(m.menu)
	return ncursesError(syscall.Errno(err))
}

// Scale
func (m *Menu) Scale() (int, int, error) {
	var y, x C.int
	err := C.scale_menu(m.menu, (*C.int)(&y), (*C.int)(&x))
	return int(y), int(x), ncursesError(syscall.Errno(err))
}

// SetBackground set the attributes of the un-highlighted items in the
// menu
func (m *Menu) SetBackground(ch Char) error {
	err := C.set_menu_back(m.menu, C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// SetForeground sets the attributes of the highlighted items in the menu
func (m *Menu) SetForeground(ch Char) error {
	err := C.set_menu_fore(m.menu, C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// SetItems will either set the items in the menu. When setting
// items you must make sure the prior menu items will be freed.
func (m *Menu) SetItems(items []*MenuItem) error {
	citems := make([]*C.ITEM, len(items)+1)
	for index, item := range items {
		citems[index] = item.item
	}
	citems[len(items)] = nil
	err := C.set_menu_items(m.menu, (**C.ITEM)(&citems[0]))
	return ncursesError(syscall.Errno(err))
}

// SetPad sets the padding character for menu items.
func (m *Menu) SetPad(ch Char) error {
	err := C.set_menu_pad(m.menu, C.int(ch))
	return ncursesError(syscall.Errno(err))
}

// SetPattern sets the padding character for menu items.
func (m *Menu) SetPattern(pattern string) error {
	cpattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cpattern))
	err := C.set_menu_pattern(m.menu, (*C.char)(cpattern))
	return ncursesError(syscall.Errno(err))
}

// SetSpacing of the the menu's items. 'desc' is the space between the
// item and it's description andmay not be larger than TAB_SIZE. 'row'
// is the number of rows separating each item and may not be larger than
// three. 'col' is the spacing between each column of items in
// multi-column mode. Use values of 0 or 1 to reset spacing to default,
// which is one
func (m *Menu) SetSpacing(desc, row, col int) error {
	err := C.set_menu_spacing(m.menu, C.int(desc), C.int(row),
		C.int(col))
	return ncursesError(syscall.Errno(err))
}

// SetWindow container for the menu
func (m *Menu) SetWindow(w *Window) error {
	err := C.set_menu_win(m.menu, w.win)
	return ncursesError(syscall.Errno(err))
}

// Spacing returns the menu item spacing. See SetSpacing for a description
func (m *Menu) Spacing() (int, int, int) {
	var desc, row, col C.int
	C.menu_spacing(m.menu, (*C.int)(&desc), (*C.int)(&row),
		(*C.int)(&col))

	return int(desc), int(row), int(col)
}

// SubWindow for the menu
func (m *Menu) SubWindow(sub *Window) error {
	err := C.set_menu_sub(m.menu, sub.win)
	return ncursesError(syscall.Errno(err))
}

// UnPost the menu, effectively hiding it.
func (m *Menu) UnPost() error {
	err := C.unpost_menu(m.menu)
	return ncursesError(syscall.Errno(err))
}

// Window container for the menu. Returns nil on failure
func (m *Menu) Window() *Window {
	return &Window{C.menu_win(m.menu)}
}

// NewItem creates a new menu item with name and description.
func NewItem(name, desc string) (*MenuItem, error) {
	cname := C.CString(name)
	cdesc := C.CString(desc)

	var item *C.ITEM
	var err error
	item, err = C.new_item(cname, cdesc)
	return &MenuItem{item}, ncursesError(err)
}

// Description returns the second value passed to NewItem
func (mi *MenuItem) Description() string {
	return C.GoString(C.item_description(mi.item))
}

// Free must be called on all menu items to avoid memory leaks
func (mi *MenuItem) Free() {
	C.free(unsafe.Pointer(C.item_name(mi.item)))
	C.free_item(mi.item)
}

// Index of the menu item in it's parent menu
func (mi *MenuItem) Index() int {
	return int(C.item_index(mi.item))
}

// Name of the menu item
func (mi *MenuItem) Name() string {
	return C.GoString(C.item_name(mi.item))
}

// Selectable turns on/off whether a menu option is "greyed out"
func (mi *MenuItem) Selectable(on bool) {
	if on {
		C.item_opts_on(mi.item, O_SELECTABLE)
	} else {
		C.item_opts_off(mi.item, O_SELECTABLE)
	}
}

// SetValue sets whether an item is active or not
func (mi *MenuItem) SetValue(val bool) error {
	err := int(C.set_item_value(mi.item, C.bool(val)))
	return ncursesError(syscall.Errno(err))
}

// Value returns true if menu item is toggled/active, otherwise false
func (mi *MenuItem) Value() bool {
	return bool(C.item_value(mi.item))
}

// Visible returns true if the item is visible, false if not
func (mi *MenuItem) Visible() bool {
	return bool(C.item_visible(mi.item))
}
