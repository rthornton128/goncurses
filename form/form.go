// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


/* ncurses form extension */
package goncurses

//#cgo LDFLAGS: -lform
//#include <form.h>
import "C"

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	FO_VISIBLE  = C.O_VISIBLE  // Field visibility
	FO_ACTIVE   = C.O_ACTIVE   // Field is sensitive/accessable
	FO_PUBLIC   = C.O_PUBLIC   // Typed characters are echoed
	FO_EDIT     = C.O_EDIT     // Editable
	FO_WRAP     = C.O_WRAP     // Line wrapping
	FO_BLANK    = C.O_BLANK    // Clear on entry
	FO_AUTOSKIP = C.O_AUTOSKIP // Skip to next field when current filled
	FO_NULLOK   = C.O_NULLOK   // Blank ok 
	FO_STATIC   = C.O_STATIC   // Fixed size
	FO_PASSOK   = C.O_PASSOK   // Field validation
)

var errList = map[C.int]string{
	C.E_SYSTEM_ERROR:    "System error occurred",
	C.E_BAD_ARGUMENT:    "Incorrect or out-of-range argument",
	C.E_POSTED:          "Form has already been posted",
	C.E_CONNECTED:       "Field is already connected to a form",
	C.E_BAD_STATE:       "Bad state",
	C.E_NO_ROOM:         "No room",
	C.E_NOT_POSTED:      "Form has not been posted",
	C.E_UNKNOWN_COMMAND: "Unknown command",
	C.E_NO_MATCH:        "No match",
	C.E_NOT_SELECTABLE:  "Not selectable",
	C.E_NOT_CONNECTED:   "Field is not connected to a form",
	C.E_REQUEST_DENIED:  "Request denied",
	C.E_INVALID_FIELD:   "Invalid field",
	C.E_CURRENT:         "Current",
}

// Form Driver Requests
const (
	REQ_NEXT_PAGE    = C.REQ_NEXT_PAGE    // next page
	REQ_PREV_PAGE    = C.REQ_PREV_PAGE    // previous page
	REQ_FIRST_PAGE   = C.REQ_FIRST_PAGE   // first page
	REQ_LAST_PAGE    = C.REQ_LAST_PAGE    // last page
	REQ_NEXT_FIELD   = C.REQ_NEXT_FIELD   // next field
	REQ_PREV_FIELD   = C.REQ_PREV_FIELD   // previous field
	REQ_FIRST_FIELD  = C.REQ_FIRST_FIELD  // first field
	REQ_LAST_FIELD   = C.REQ_LAST_FIELD   // last field
	REQ_SNEXT_FIELD  = C.REQ_SNEXT_FIELD  // sorted next field
	REQ_SPREV_FIELD  = C.REQ_SPREV_FIELD  // sorted previous field
	REQ_SFIRST_FIELD = C.REQ_SFIRST_FIELD // sorted first field
	REQ_SLAST_FIELD  = C.REQ_SLAST_FIELD  // sorted last field
	REQ_LEFT_FIELD   = C.REQ_LEFT_FIELD   // left field
	REQ_RIGHT_FIELD  = C.REQ_RIGHT_FIELD  // right field
	REQ_UP_FIELD     = C.REQ_UP_FIELD     // up to a field
	REQ_DOWN_FIELD   = C.REQ_DOWN_FIELD   // down to a field
	REQ_NEXT_CHAR    = C.REQ_NEXT_CHAR    // next character in field
	REQ_PREV_CHAR    = C.REQ_PREV_CHAR    // previous character in field
	REQ_NEXT_LINE    = C.REQ_NEXT_LINE    // next line
	REQ_PREV_LINE    = C.REQ_PREV_LINE    // previous line
	REQ_NEXT_WORD    = C.REQ_NEXT_WORD    // next word
	REQ_PREV_WORD    = C.REQ_PREV_WORD    // previous word
	REQ_BEG_FIELD    = C.REQ_BEG_FIELD    // beginning of field
	REQ_END_FIELD    = C.REQ_END_FIELD    // end of field
	REQ_BEG_LINE     = C.REQ_BEG_LINE     // beginning of line
	REQ_END_LINE     = C.REQ_END_LINE     // end of line
	REQ_LEFT_CHAR    = C.REQ_LEFT_CHAR    // character to the left
	REQ_RIGHT_CHAR   = C.REQ_RIGHT_CHAR   // character to the right
	REQ_UP_CHAR      = C.REQ_UP_CHAR      // up a character
	REQ_DOWN_CHAR    = C.REQ_DOWN_CHAR    // down a character
	REQ_NEW_LINE     = C.REQ_NEW_LINE     // insert of overlay a new line
	REQ_INS_CHAR     = C.REQ_INS_CHAR     // insert a blank character at cursor
	REQ_INS_LINE     = C.REQ_INS_LINE     // insert a blank line at cursor
	REQ_DEL_CHAR     = C.REQ_DEL_CHAR     // delete character at cursor
	REQ_DEL_PREV     = C.REQ_DEL_PREV     // delete character before cursor
	REQ_DEL_LINE     = C.REQ_DEL_LINE     // delete line
	REQ_DEL_WORD     = C.REQ_DEL_WORD     // delete word
	REQ_CLR_EOL      = C.REQ_CLR_EOL      // clear from cursor to end of line
	REQ_CLR_EOF      = C.REQ_CLR_EOF      // clear from cursor to end of field
	REQ_CLR_FIELD    = C.REQ_CLR_FIELD    // clear field
	REQ_OVL_MODE     = C.REQ_OVL_MODE     // overlay mode
	REQ_INS_MODE     = C.REQ_INS_MODE     // insert mode
	REQ_SCR_FLINE    = C.REQ_SCR_FLINE    // scroll field forward a line
	REQ_SCR_BLINE    = C.REQ_SCR_BLINE    // scroll field back a line
	REQ_SCR_FPAGE    = C.REQ_SCR_FPAGE    // scroll field forward a page
	REQ_SCR_BPAGE    = C.REQ_SCR_BPAGE    // scroll field back a page
	REQ_SCR_FHPAGE   = C.REQ_SCR_FHPAGE   // scroll field forward half a page
	REQ_SCR_BHPAGE   = C.REQ_SCR_BHPAGE   // scroll field back half a page
	REQ_SCR_FCHAR    = C.REQ_SCR_FCHAR    // scroll field forward a character
	REQ_SCR_BCHAR    = C.REQ_SCR_BCHAR    // scroll field back a character
	REQ_SCR_HFLINE   = C.REQ_SCR_HFLINE   // horisontal scroll field forward a line
	REQ_SCR_HBLINE   = C.REQ_SCR_HBLINE   // horisontal scroll field back a line
	REQ_SCR_HFHALF   = C.REQ_SCR_HFHALF   // horisontal scroll field forward half a line
	REQ_SCR_HBHALF   = C.REQ_SCR_HBHALF   // horisontal scroll field back half a line
	REQ_VALIDATION   = C.REQ_VALIDATION   // validate field
	REQ_NEXT_CHOICE  = C.REQ_NEXT_CHOICE  // display next field choice
	REQ_PREV_CHOICE  = C.REQ_PREV_CHOICE  // display previous field choice
)

type Field struct {
	field *C.FIELD
}

type Form struct {
	form *C.FORM
}

func ncursesError(e error) error {
	errno, ok := e.(syscall.Errno)
	if int(errno) == C.OK {
		e = nil
	}
	if ok {
		errstr, ok := errList[C.int(errno)]
		if ok {
			return errors.New(errstr)
		}
	}
	return e
}

func NewField(h, w, tr, lc, oscr, nbuf int) (*Field, error) {
	var new_field Field
	var err error
	new_field.field, err = C.new_field(C.int(h), C.int(w), C.int(tr), C.int(lc),
		C.int(oscr), C.int(nbuf))
	return &new_field, ncursesError(err)
}

// Background returns the field's background character attributes
func (f *Field) Background() int {
	return int(C.field_back(f.field))
}

// Duplicate the field at the specified coordinates, returning a pointer 
// to the newly allocated object.
func (f *Field) Duplicate(y, x int) (*Field, error) {
	var new_field Field
	var err error
	new_field.field, err = C.dup_field(f.field, C.int(y), C.int(x))
	return &new_field, ncursesError(err)
}

// Foreground returns the field's foreground character attributes
func (f *Field) Foreground() int {
	return int(C.field_fore(f.field))
}

// Free field's allocated memory. This must be called to prevent memory
// leaks
func (f *Field) Free() error {
	err := C.free_field(f.field);
	f = nil
	return ncursesError(syscall.Errno(err))
}

// Info retrieves the height, width, y, x, offset and buffer size of the 
// given field. Pass the memory addess of the variable to store the data
// in or nil.
func (f *Field) Info(h, w, y, x, off, nbuf *int) error {
	err := C.field_info(f.field, (*C.int)(unsafe.Pointer(h)),
		(*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(y)),
		(*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(off)),
		(*C.int)(unsafe.Pointer(nbuf)))
	return ncursesError(syscall.Errno(err))
}

// Just returns the justification type of the field
func (f *Field) Justification() int {
	return int(C.field_just(f.field))
}

// Move the field to the location of the specified coordinates
func (f *Field) Move(y, x int) error {
	err := C.move_field(f.field, C.int(y), C.int(x));
	return ncursesError(syscall.Errno(err))
}

// Options turns features on and off
func (f *Field) Options(opts int, on bool) {
	if on {
		C.field_opts_on(f.field, C.Field_Options(opts))
		return
	}
	C.field_opts_off(f.field, C.Field_Options(opts))
}

// Pad returns the padding character of the field
func (f *Field) Pad() int {
	return int(C.field_pad(f.field))
}

// SetJustification of the field
func (f *Field) SetJustification(just int) error {
	err := C.set_field_just(f.field, C.int(just))
	return ncursesError(syscall.Errno(err))
}

// OptionsOff turns feature(s) off
func (f *Field) OptionsOff(opts int) error {
	err := int(C.field_opts_off(f.field, C.Field_Options(opts)))
	if err != C.E_OK {
		return ncursesError(syscall.Errno(err))
	}
	return nil
}

// OptionsOn turns feature(s) on
func (f *Field) OptionsOn(opts int) error {
	err := int(C.field_opts_on(f.field, C.Field_Options(opts)))
	if err != C.E_OK {
		return ncursesError(syscall.Errno(err))
	}
	return nil
}

// SetPad sets the padding character of the field
func (f *Field) SetPad(padch int) error {
	err := C.set_field_pad(f.field, C.int(padch))
	return ncursesError(syscall.Errno(err))
}

// SetBackground character and attributes (colours, etc)
func (f *Field) SetBackground(ch int) error {
	err := C.set_field_back(f.field, C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// SetForeground character and attributes (colours, etc)
func (f *Field) SetForeground(ch int) error {
	err := C.set_field_fore(f.field, C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// NewForm returns a new form object using the fields array supplied as
// an argument
func NewForm(fields []*Field) (Form, error) {
	cfields := make([]*C.FIELD, len(fields)+1)
	for index, field := range fields {
		cfields[index] = field.field
	}
	cfields[len(fields)] = nil
	
	var form *C.FORM
	var err error
	form, err = C.new_form((**C.FIELD)(&cfields[0]))
	
	return Form{form}, ncursesError(err)
}

// FieldCount returns the number of fields attached to the Form
func (f *Form) FieldCount() int {
	return int(C.field_count(f.form))
}

// Driver issues the actions requested to the form itself. See the
// corresponding REQ_* constants
func (f *Form) Driver(drvract int) error {
	err := C.form_driver(f.form, C.int(drvract))
	return ncursesError(syscall.Errno(err))
}

// Free the memory allocated to the form. Forms are not automatically
// free'd by Go's garbage collection system so the memory allocated to
// it must be explicitely free'd
func (f *Form) Free() error {
	err := C.free_form(f.form)
	f = nil
	return ncursesError(syscall.Errno(err))
}

// Post the form, making it visible and interactive
func (f *Form) Post() error {
	err := C.post_form(f.form)
	return ncursesError(syscall.Errno(err))
}

// SetFields overwrites the current fields for the Form with new ones.
// It is important to make sure all prior fields have been freed otherwise
// this action will result in a memory leak
func (f *Form) SetFields(fields []*Field) error {
	cfields := make([]*C.FIELD, len(fields)+1)
	for index, field := range fields {
		cfields[index] = field.field
	}
	cfields[len(fields)] = nil
	err := C.set_form_fields(f.form, (**C.FIELD)(&cfields[0]))
	return ncursesError(syscall.Errno(err))
}

// SetOptions for the form
func (f *Form) SetOptions(opts int) error {
	_, err := C.set_form_opts(f.form, (C.Form_Options)(opts))
	return ncursesError(err)
}

// SetSub sets the subwindow associated with the form
func (f *Form) SetSub(w *Window) error {
	err := int(C.set_form_sub(f.form, w.win))
	return ncursesError(syscall.Errno(err))
}

// SetWindow sets the window associated with the form
func (f *Form) SetWindow(w *Window) error {
	err := int(C.set_form_win(f.form, w.win))
	return ncursesError(syscall.Errno(err))
}

// Sub returns the subwindow assocaiated with the form
func (f *Form) Sub() Window {
	return Window{C.form_sub(f.form)}
}

// UnPost the form, removing it from the interface
func (f *Form) UnPost() error {
	err := C.unpost_form(f.form)
	return ncursesError(syscall.Errno(err))
}
