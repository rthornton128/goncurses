// goncurses - ncurses library for Go.

/* This example demonstrates the ability to resize. Only one of detecting SIGWINCH or KEY_RESIZE
 * is strictly needed, but depending on the options ncurses was built with, one or the other may
 * work better. */
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var stdscr *gc.Window

func main() {
	start := time.Now()
	sigWinChCount := 0
	keyResizeCount := 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGWINCH)

	stdscr, _ = gc.Init()
	stdscr.Timeout(0)
	defer gc.End()

	for {
		select {
		case <-sigs:
			sigWinChCount++
			resize()
		default:
			c := stdscr.GetChar()
			if c == 'q' {
				return
			} else if c == gc.KEY_RESIZE {
				keyResizeCount++
				//resize()
			}
		}
		row, col := stdscr.MaxYX()
		tRow, tCol, _ := osTermSize()
		runtime := int(time.Since(start).Seconds())
		stdscr.MovePrintf(1, 1, "     MaxYX shows %d rows and %d columns", row, col)
		stdscr.MovePrintf(2, 1, "osTermSize shows %d rows and %d columns", tRow, tCol)
		stdscr.MovePrintf(3, 1, "  SIGWINCH has been sent %d times", sigWinChCount)
		stdscr.MovePrintf(4, 1, "KEY_RESIZE has been sent %d times", keyResizeCount)
		stdscr.MovePrintf(5, 1, "The program has been running for %d seconds", runtime)
		_ = stdscr.Box(0, 0)
		stdscr.Refresh()
	}
}

func resize() {
	gc.End()

	row, col, _ := osTermSize()
	_ = gc.ResizeTerm(row, col)

	stdscr, _ = gc.Init()
	_ = stdscr.Clear()
	stdscr.Timeout(0)
}
