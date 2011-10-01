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
    menu_opts := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}
    var [len(menu_opts)]items *MenuItem
    stdscr, _ := Initscr();
    defer Endwin()
    
    Raw()
    Noecho()
    CursSet(0)
    stdscr.Clear()
    stdscr.Keypad(true)
    
    for i := 0; i < len(menu_opts); {
        items[i] = NewItem(menu_opts[i], menu_opts[i])
        defer items[i].Free()
    }
    
    menu := NewMenu(items)
    defer menu.Free()
    stdscr.Print(20, 0, "'q' to exit")
    stdscr.Refresh()
    
    for {
        ch := stdscr.Getch()
        
        switch (Key(ch)) {
        case "q":
            return
        case "down":
            menu.Down()
        case "up":
            menu.Up()
        }
    }
}
