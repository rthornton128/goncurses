/* A simmple example of how to use the panel ncurses library */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()

    StartColor()
    CBreak(true)
    Echo(true)
    stdscr.Keypad(true)
    stdscr.Print("Hit 'tab' to cycle through windows, 'q' to quit")
    
    InitPair(1, C_RED, C_BLACK)
    InitPair(2, C_GREEN, C_BLACK)
    InitPair(3, C_BLUE, C_BLACK)
    InitPair(4, C_CYAN, C_BLACK)
    
    var panels [3]*Panel
    y, x := 4, 10
    
    for i := 0; i < 3; i++ {
        h, w := 10, 40
        title := "Window Number %d"
        
        window, _ := NewWindow(h, w , y+(i*4), x+(i*10))
        window.Box(0, 0)
        window.AddChar(2, 0, ACS_LTEE)
        window.HLine(2, 1, ACS_HLINE, w - 2)
        window.AddChar(2, w-1, ACS_RTEE)
        window.ColorOn(byte(i+1))
        window.Print(1, (w/2)-(len(title)/2), title, i+1)
        window.ColorOff(byte(i+1))
        panels[i] = window.Panel()
        
    }

    active := 2
    
    for {
        UpdatePanels()
        Update()
        
        ch := stdscr.GetChar()
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
