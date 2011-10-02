/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import . "goncurses"
import "os"

const (
    HEIGHT = 10
    WIDTH = 30
)

func main() {
    stdscr, _ := Initscr();
    defer Endwin()
    
    Raw()
    Noecho()
    CursSet(0)
    stdscr.Clear()
    stdscr.Keypad(true)
    
    menu_items := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", 
        "Exit"}
    items := make([]*MenuItem, len(menu_items))
    var err os.Error
    for i, val := range menu_items {
        items[i], err = NewItem(val, "")
        defer items[i].Free()
    }
    
    menu, err := NewMenu(items)
    defer menu.Free()
    menu.Post()
    
    if err != nil {
        stdscr.Print(err.String())
        return
    }

    stdscr.Mvprint(20, 0, "'q' to exit")
    stdscr.Refresh()
    
    for {
        DoUpdate()
        ch, _ := stdscr.Getch()
        
        switch (Key(ch)) {
        case "q":
            return
        case "down":
            menu.Driver("down")
        case "up":
            menu.Driver("up")
        }
    }
}
