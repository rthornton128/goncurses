/* This example demonstrates the use of the menu library, similar to that 
 * found in the ncurses examples from TLDP */

package main

import . "goncurses.googlecode.com/hg/goncurses"
import "os"

const (
    HEIGHT = 10
    WIDTH = 30
)

func main() {
    stdscr, _ := Init();
    defer End()
    
    Raw(true)
    Echo(false)
    Cursor(0)
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

    stdscr.Print(20, 0, "'q' to exit")
    stdscr.Refresh()
    
    for {
        Update()
        ch := stdscr.GetChar()
        
        switch (Key(ch)) {
        case "q":
            return
        case "down":
            menu.Driver(REQ_DOWN)
        case "up":
            menu.Driver(REQ_UP)
        }
    }
}
