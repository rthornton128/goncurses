// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goncurses

// #include <stdio.h>
// #if defined(__MINGW32__) || defined(__MINGW64__)
// FILE *fdopen(int fildes, const char *mode) { return _fdopen(fildes, mode); }
// #endif
// #include <stdlib.h>
// #include <curses.h>
import "C"

import (
	"errors"
	"os"
	"unsafe"
)

type Screen struct{ scrPtr *C.SCREEN }

// NewTerm returns a new Screen, representing a physical terminal. If using
// this function to generate a new Screen you should not call Init().
// Unlike Init(), NewTerm does not call Refresh() to clear the screen so this
// will need to be done manually. When finished with a terminal, you must
// call End() in reverse order that each terminal was created in. After you
// are finished with the screen you must call Delete to free the memory
// allocated to it. This function is usually only useful for programs using
// multiple terminals or test for terminal capabilites. The argument termType
// is the type of terminal to be used ($TERM is used if value is "" which also
// has the same effect of using os.Getenv("TERM"))
func NewTerm(termType string, out, in *os.File) (*Screen, error) {
	var tt, wr, rd *C.char
	if termType == "" {
		tt, wr, rd = (*C.char)(nil), C.CString("w"), C.CString("r")
	} else {
		tt, wr, rd = C.CString(termType), C.CString("w"), C.CString("r")
		defer C.free(unsafe.Pointer(tt))
	}
	defer C.free(unsafe.Pointer(wr))
	defer C.free(unsafe.Pointer(rd))

	cout, cin := C.fdopen(C.int(out.Fd()), wr), C.fdopen(C.int(in.Fd()), rd)
	screen := C.newterm(tt, cout, cin)
	if screen == nil {
		return nil, errors.New("Failed to create new screen")
	}
	return &Screen{screen}, nil
}

// Set the screen to be the current, active screen
func (s *Screen) Set() (*Screen, error) {
	screen := C.set_term(s.scrPtr)
	if screen == nil {
		return nil, errors.New("Failed to set screen")
	}
	return &Screen{screen}, nil
}

// Delete frees memory allocated to the screen. This function
func (s *Screen) Delete() {
	C.delscreen(s.scrPtr)
}

// End is just a wrapper for the global End function. This helper function
// has been provided to help ensure that new terminals are closed in the
// proper, reverse order they were created. It makes the terminal active via
// set then called End so it is closed properly. You must make sure that
// Delete is called once done with the screen/terminal.
func (s *Screen) End() {
	s.Set()
	End()
}
