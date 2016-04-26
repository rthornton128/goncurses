// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package goncurses

// #cgo !darwin,!openbsd pkg-config: form
// #cgo darwin openbsd LDFLAGS: -lform
// #include <form.h>
// #include <stdlib.h>
import "C"

import (
	"syscall"
	"unsafe"
)

type Field C.FIELD

type Form struct {
	form *C.FORM
}

func NewField(h, w, tr, lc, oscr, nbuf int32) (*Field, error) {
	f, err := C.new_field(C.int(h), C.int(w), C.int(tr), C.int(lc),
		C.int(oscr), C.int(nbuf))
	return (*Field)(f), ncursesError(err)
}

// Background returns the field's background character attributes
func (f *Field) Background() Char {
	return Char(C.field_back((*C.FIELD)(f)))
}

// Buffer returns a string containing the contents of the buffer. The returned
// string will contain whitespace up to the buffer size as set by SetMax or
// the value by the call to NewField
func (f *Field) Buffer() string {
	str := C.field_buffer((*C.FIELD)(f), C.int(0))

	return C.GoString(str)
}

// Duplicate the field at the specified coordinates, returning a pointer
// to the newly allocated object.
func (f *Field) Duplicate(y, x int32) (*Field, error) {
	nf, err := C.dup_field((*C.FIELD)(f), C.int(y), C.int(x))
	return (*Field)(nf), ncursesError(err)
}

// Foreground returns the field's foreground character attributes
func (f *Field) Foreground() Char {
	return Char(C.field_fore((*C.FIELD)(f)))
}

// Free field's allocated memory. This must be called to prevent memory
// leaks
func (f *Field) Free() error {
	err := C.free_field((*C.FIELD)(f))
	f = nil
	return ncursesError(syscall.Errno(err))
}

// Info retrieves the height, width, y, x, offset and buffer size of the
// given field. Pass the memory addess of the variable to store the data
// in or nil.
func (f *Field) Info(h, w, y, x, off, nbuf *int) error {
	err := C.field_info((*C.FIELD)(f), (*C.int)(unsafe.Pointer(h)),
		(*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(y)),
		(*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(off)),
		(*C.int)(unsafe.Pointer(nbuf)))
	return ncursesError(syscall.Errno(err))
}

// Just returns the justification type of the field
func (f *Field) Justification() int {
	return int(C.field_just((*C.FIELD)(f)))
}

// Move the field to the location of the specified coordinates
func (f *Field) Move(y, x int32) error {
	err := C.move_field((*C.FIELD)(f), C.int(y), C.int(x))
	return ncursesError(syscall.Errno(err))
}

// Options turns features on and off
func (f *Field) Options(opts int, on bool) {
	if on {
		C.field_opts_on((*C.FIELD)(f), C.Field_Options(opts))
		return
	}
	C.field_opts_off((*C.FIELD)(f), C.Field_Options(opts))
}

// Pad returns the padding character of the field
func (f *Field) Pad() int {
	return int(C.field_pad((*C.FIELD)(f)))
}

// SetBuffer sets the visible characters in the field. A buffer is empty by
// default.
func (f *Field) SetBuffer(s string) error {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	err := C.set_field_buffer((*C.FIELD)(f), C.int(0), cstr)
	return ncursesError(syscall.Errno(err))
}

// SetJustification of the field
func (f *Field) SetJustification(just int) error {
	err := C.set_field_just((*C.FIELD)(f), C.int(just))
	return ncursesError(syscall.Errno(err))
}

// SetMax sets the maximum size of a field
func (f *Field) SetMax(max int) error {
	err := C.set_max_field((*C.FIELD)(f), C.int(max))
	return ncursesError(syscall.Errno(err))
}

// OptionsOff turns feature(s) off
func (f *Field) SetOptionsOff(opts Char) error {
	err := int(C.field_opts_off((*C.FIELD)(f), C.Field_Options(opts)))
	if err != C.E_OK {
		return ncursesError(syscall.Errno(err))
	}
	return nil
}

// OptionsOn turns feature(s) on
func (f *Field) SetOptionsOn(opts Char) error {
	err := int(C.field_opts_on((*C.FIELD)(f), C.Field_Options(opts)))
	if err != C.E_OK {
		return ncursesError(syscall.Errno(err))
	}
	return nil
}

// SetPad sets the padding character of the field
func (f *Field) SetPad(padch int) error {
	err := C.set_field_pad((*C.FIELD)(f), C.int(padch))
	return ncursesError(syscall.Errno(err))
}

// SetBackground character and attributes (colours, etc)
func (f *Field) SetBackground(ch Char) error {
	err := C.set_field_back((*C.FIELD)(f), C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// SetForeground character and attributes (colours, etc)
func (f *Field) SetForeground(ch Char) error {
	err := C.set_field_fore((*C.FIELD)(f), C.chtype(ch))
	return ncursesError(syscall.Errno(err))
}

// NewForm returns a new form object using the fields array supplied as
// an argument
func NewForm(fields []*Field) (Form, error) {
	if fields[len(fields)-1] != nil {
		fields = append(fields, nil)
	}
	form, err := C.new_form((**C.FIELD)(unsafe.Pointer(&fields[0])))

	return Form{form}, ncursesError(err)
}

// FieldCount returns the number of fields attached to the Form
func (f *Form) FieldCount() int {
	return int(C.field_count(f.form))
}

// Driver issues the actions requested to the form itself. See the
// corresponding REQ_* constants
func (f *Form) Driver(drvract Key) error {
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
	//cfields := make([]*C.FIELD, len(fields)+1)
	//for index, field := range fields {
	//cfields[index] = field.field
	//}
	//cfields[len(fields)] = nil
	err := C.set_form_fields(f.form, (**C.FIELD)(unsafe.Pointer(&fields[0])))
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
