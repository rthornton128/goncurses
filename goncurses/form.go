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

type Field C.FIELD

func NewField(h, w, tr, lc, os, nbuf int) *Field {
	return (*Field)(C.new_field(C.int(h), C.int(w), C.int(tr), C.int(lc),
			C.int(os), C.int(nbuf)))
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

func NewForm(fields []*Field) {
	
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
