// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates how one might write to multiple terminals from a
// single program or redirect output. In order to run the program you must
// supply a single argument with is a path to a pseudo-terminal device.
//
// This example should compile on Windows but running it may be problematic
package main

import (
	"flag"
	gc "github.com/rthornton128/goncurses"
	"log"
	"os"
)

func main() {
	var term1, term2 *gc.Screen
	var err error

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Must supply path to a pseudo-terminal (ie. /dev/pts/0)")
	}

	// You can, for example, use the current psuedo-terminal to read/write to
	// You may need to change the  to one you have read/write access to
	var pts *os.File
	pts, err = os.OpenFile(flag.Arg(0), os.O_RDWR, os.FileMode(666))
	if err != nil {
		log.Fatal(err)
	}
	defer pts.Close()

	// Create a new terminal using the default $TERM type and the psuedo-terminal
	term1, err = gc.NewTerm("", pts, pts)
	if err != nil {
		log.Fatal("newterm:", err)
	}
	// Remember that defer is LIFO order and End must be called prior to Delete
	defer term1.Delete()
	defer term1.End()

	// Create a second terminal as if we had two to interact with
	term2, err = gc.NewTerm("", os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal("newterm:", err)
	}
	// It is important that End is called on term2 prior to term1
	defer term2.Delete()
	defer term2.End() // comment out for an alternative to End
	//defer gc.End()    // uncomment for alternate way to end
	//defer term2.Set() // uncomment

	// Set the active terminal to term1
	term1.Set()

	// Get the Standard Screen Window and write to the active terminal
	mw := gc.StdScr()
	mw.MovePrint(0, 0, "Term 1 works! Press any key to exit...")
	mw.Refresh()
	mw.GetChar()

	// activate term2 and send data to it
	term2.Set()
	mw.MovePrint(0, 0, "Term 2 works! Press any key to exit...")
	mw.Refresh()
	mw.GetChar()
}
