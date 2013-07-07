package main

import (
	gc "code.google.com/p/goncurses"
	"log"
	"os"
)

func main() {
	var term1, term2 *gc.Screen
	var err error

	// You can, for example, use the current psuedo-terminal to read/write to
	var pts *os.File
	pts, err = os.OpenFile("/dev/pts/0", os.O_RDWR, os.FileMode(666))
	if err != nil {
		log.Fatal(nil)
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
	mw := gc.StdScr() // misleading...use a StdScr() method instead?
	mw.MovePrint(0, 0, "Term 1 works! Press any key to exit...")
	mw.Refresh()
	mw.GetChar()

	// activate term2 and send data to it
	term2.Set()
	mw.MovePrint(0, 0, "Term 2 works! Press any key to exit...")
	mw.Refresh()
	mw.GetChar()
}
