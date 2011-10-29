/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init();
    defer End()
    
    StartColor()
    Raw(true)
    Echo(false)
    Cursor(0)
    stdscr.Keypad(true)
    InitPair(1, C_RED, C_BLACK)
    
    // build the menu items
    menu_items := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", 
        "Exit"}
    items := make([]*MenuItem, len(menu_items))
    for i, val := range menu_items {
        items[i], _ = NewItem(val, "")
        defer items[i].Free()
    }
    
    // create the menu
    menu, _ := NewMenu(items)
    defer menu.Free()
    
    menuwin, _ := NewWindow(10, 40, 4, 14)
    menuwin.Keypad(true)
    
    menu.Window(menuwin)
    menu.Sub(menuwin.Derived(6, 38, 3, 1))
    menu.Mark(" * ")
    
    // Print centered menu title
    y, x := menuwin.Maxyx()
    title := "My Menu"
    menuwin.Box(0, 0)
    menuwin.ColorOn(1)
    menuwin.Print(1, (x/2)-(len(title)/2), title)
    menuwin.ColorOff(1)
    menuwin.AddChar(2, 0, ACS_LTEE)
    menuwin.HLine(2, 1, ACS_HLINE, x-2)
    menuwin.AddChar(2, x-1, ACS_RTEE)
    
    y, x = stdscr.Maxyx()
    stdscr.Print(y-2, 1, "'q' to exit")
    stdscr.Refresh()
    
    menu.Post()
    defer menu.UnPost()
    menuwin.Refresh()
    
    for {
        Update()
        ch, _ := menuwin.GetChar()
        
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
