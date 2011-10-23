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

/* Form library */
package goncurses

//#cgo LDFLAGS: -lform
//#include <form.h>
import "C"

import (
	"fmt"
	"os"
)

var errList = map[C.int]string{
	C.E_OK:              "Routine succeeded",
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

func error(e os.Error) os.Error {
	s, ok := errList[C.int(e.(os.Errno))]
	if !ok {
		return os.NewError(fmt.Sprintf("Error %d", int(e.(os.Errno))))
	}
	return os.NewError(s)
}

type Field C.FIELD

func NewField(h, w, tr, lc, oscr, nbuf int) (*Field, os.Error) {
	field, err := C.new_field(C.int(h), C.int(w), C.int(tr), C.int(lc),
		C.int(oscr), C.int(nbuf))
	return (*Field)(field), error(err)
}

func (f *Field) Background(ch int) {
	C.set_field_back((*C.FIELD)(f), C.chtype(ch))
}

func (f *Field) Foreground(ch int) {
	C.set_field_fore((*C.FIELD)(f), C.chtype(ch))
}

func (f *Field) Free() {
	C.free_field((*C.FIELD)(f))
	f = nil
}

func (f *Field) Options(opts int) {
	C.set_field_fore((*C.FIELD)(f), C.chtype(opts))
}

type Form C.FORM

func NewForm(fields []*Field) (*Form, os.Error) {
	cfields := make([]*C.FIELD, len(fields)+1)
	for index, field := range fields {
		cfields[index] = (*C.FIELD)(field)
	}
	cfields[len(fields)] = nil
	f, e := C.new_form((**C.FIELD)(&cfields[0]))
	return (*Form)(f), error(e)
}

func (f *Form) Free() {
	C.free_form((*C.FORM)(f))
	f = nil
}

func (f *Form) Post() {
	C.post_form((*C.FORM)(f))
}

func (f *Form) UnPost() {
	C.unpost_form((*C.FORM)(f))
}
