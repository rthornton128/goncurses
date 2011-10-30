/* A simmple example of how to use the panel ncurses library */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()

    var panels [3]*Panel
    y, x := 2, 4
    
    for i := 0; i < 3; i++ {
        window, _ := NewWindow(10, 40 , y+i, x+(i*5))
        window.Box(0, 0)
        panels[i] = NewPanel(window)
    }
    
    UpdatePanels()
    Update()

    stdscr.GetChar()
}
