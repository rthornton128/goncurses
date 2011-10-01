/* A simmple example of how to use the panel ncurses library */

package main

import . "goncurses"
import "fmt"

func main() {
    stdscr, _ := Initscr()
    defer Endwin()

    StartColor()
    Cbreak()
    Noecho()
    stdscr.Keypad(true)
    stdscr.Print("Hit 'tab' to cycle through windows, 'q' to quit")
    
    InitPair(1, "red", "black")
    InitPair(2, "green", "black")
    InitPair(3, "blue", "black")
    InitPair(4, "cyan", "black")
    
    var panels [3]*Panel
    y, x := 4, 10
    
    for i := 0; i < 3; i++ {
        h, w := 10, 40
        title := fmt.Sprintf("Window Number %d", i+1)
        
        window, _ := NewWin(h, w , y+(i*4), x+(i*10))
        window.Box(0, 0)
        //window.MvAddch(2, 0, ACS_LTEE)
        window.HLine(2, 1, '-', w - 2) //ACS_HLINE
        //window.MvAddch(2, 0, ACS_RTEE)
        window.ColorOn(byte(i+1))
        window.Mvprint(1, (w/2)-(len(title)/2), title)
        window.ColorOff(byte(i+1))
        panels[i] = NewPanel(window)
        
    }

    active := 2
    
    for {
        UpdatePanels()
        DoUpdate()
        
        ch, _ := stdscr.Getch()
        switch(Key(ch)) {
        case "q":
            return
        case "tab":
            active += 1
            if active > 2 {
                active = 0
            }
            panels[active].Top()
        }
    }
}
