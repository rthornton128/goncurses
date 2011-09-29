/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import . "goncurses"

const (
    HEIGHT = 10
    WIDTH = 30
)

func main() {
    var active int
    menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}
    
    stdscr, _ := Initscr();
    defer Endwin()
    
    Raw()
    Noecho()
    CursSet(0)
    stdscr.Clear()
    stdscr.Keypad(true)
    
    rows, cols := stdscr.Getmaxyx()
    y, x := (rows-HEIGHT)/2, (cols-WIDTH)/2
    
    win, _ := NewWin(HEIGHT, WIDTH, y, x)
    win.Keypad(true)
    
    stdscr.Mvprint(0, 0, "Use arrow keys to go up and down, Press enter to select")
    stdscr.Refresh()
    
    printmenu(win, menu, active)
    MouseMask("button1-clicked")
    
    for {
        ch, _ := stdscr.Getch()
        switch(Key(ch)) {
        case "q":
            return
        case "up":
            if active == 0 {
                active = len(menu)-1
            } else {
                active -= 1
            }
        case "down":
            if active == len(menu)-1 {
                active = 0
            } else {
                active += 1
            }
        case "mouse":
            md, _ := GetMouse()
            new := getactive(x, y, md[0], md[1], menu)
            if new != -1 {
                active = new
            }
            stdscr.Mvprint(23, 0, "Choice #%d: %s selected", active, 
                menu[active])
            stdscr.ClearToEOL()
            stdscr.Refresh()
        case "enter":
            stdscr.Mvprint(23, 0, "Choice #%d: %s selected", active, 
                menu[active])
            stdscr.ClearToEOL()
            stdscr.Refresh()
        default:
            stdscr.Mvprint(23, 0, "Character pressed = %3d/%c", ch, ch)
            stdscr.ClearToEOL()
            stdscr.Refresh()
        }

        printmenu(win, menu, active)
    }
}

func getactive(x, y, mx, my int, menu []string) int {
    row := my - y - 2
    col := mx - x - 2
    
    if row < 0 || row > len(menu)-1 {
        return -1
    }
    
    l := menu[row]
    
    if col >= 0 && col < len(l) {
        return row
    }
    return -1
}

func printmenu(w *Window, menu []string, active int) {
    y, x := 2, 2
    w.Box(0, 0)
    for i, s := range menu {
        if i == active {
            w.Attron("reverse")
            w.Mvprint(y+i, x, s)
            w.Attroff("reverse")
        } else {
            w.Mvprint(y+i, x, s)        
        }
    }
    w.Refresh()
}
