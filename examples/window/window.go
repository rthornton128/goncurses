/* This simnple example mirrors the "hello world" TLDP ncurses howto */

package main

import . "goncurses"

func main() {    
    stdscr, _ := Initscr()
    defer Endwin()
    
    Noecho()
    Cbreak()
    stdscr.Keypad(true)
    stdscr.Print("Press 'q' to exit")
    stdscr.Refresh()
    
    rows, cols := stdscr.Getmaxyx()
    height, width := 3, 10
    y, x := (rows-height)/2, (cols-width)/2
    win, _ := NewWin(height, width, y, x)
    win.Refresh()
    
    ch, _ := stdscr.Getch()
    
    for {
        switch(Key(ch)) {
        case "q":
            return
        case "left":
            x -= 1
        case "right":
            x += 1
        case "up":
            y -= 1
        case "down":
            y += 1
        }
        destroy(win)
        win = createWin(height, width, y, x)
        
        ch, _ = stdscr.Getch()
    }
}

func createWin(h, w, y, x int) *Window {
    new, _ := NewWin(h, w, y, x)
    new.Box(0, 0)
    new.Refresh()
    return new
}

func destroy(w *Window) {
    w.Border(' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ')
    w.Refresh()
    w.DelWin()
    return
}
