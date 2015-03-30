// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates using goncurses with Go's built-in concurrency
// primatives. It is key to ensure no reads or writes occur concurrently to
// a window or screen. Another method would be to use a global mutex.
package main

import (
	gc "github.com/rthornton128/goncurses"
	"log"
)

func main() {
	scr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	gc.Echo(false)

	scr.Println("Type characters to have them appear on the screen.")
	scr.Println("Press 'q' to exit.")
	scr.Println()

	// Accept input concurrently via a goroutine and connect a channel
	in := make(chan gc.Char)
	ready := make(chan bool)
	go func(w *gc.Window, ch chan<- gc.Char) {
		for {
			// Block until all write operations are complete
			<-ready
			// Send typed character down the channel (which is blocking
			// in the main loop)
			ch <- gc.Char(w.GetChar())
		}
	}(scr, in)

	// Once a character has been received on the 'in' channel the
	// 'ready' channel will block until it recieves another piece of data.
	// This happens only once the received character has been written to
	// the screen. The 'in' channel then blocks on the next loop until
	// another 'true' is sent down the 'ready' channel signalling to the
	// input goroutine that it's okay to receive input
	for {
		var c gc.Char
		select {
		case c = <-in: // blocks while waiting for input from goroutine
			scr.Print(string(c))
			scr.Refresh()
		case ready <- true: // sends once above block completes
		}
		// Exit when 'q' is pressed
		if c == gc.Char('q') {
			break
		}
	}
}
