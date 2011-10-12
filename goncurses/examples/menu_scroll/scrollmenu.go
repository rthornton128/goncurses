/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import . "goncurses.googlecode.com/hg/goncurses"

const (
    HEIGHT = 10
    WIDTH = 40
)

func main() {
    stdscr, _ := Init();
    defer End()
    
    StartColor()
    Raw(true)
    Echo(false)
    Cursor(0)
    stdscr.Keypad(true)
    InitPair(1, "red", "black")
    InitPair(2, "cyan", "black")
    
    // build the menu items
    menu_items := []string{
        "Choice 1", 
        "Choice 2", 
        "Choice 3", 
        "Choice 4", 
        "Choice 5", 
        "Choice 6", 
        "Choice 7", 
        "Choice 8", 
        "Choice 9", 
        "Choice 10", 
        "Exit"}
    items := make([]*MenuItem, len(menu_items))
    for i, val := range menu_items {
        items[i], _ = NewItem(val, "")
        defer items[i].Free()
    }
    
    // create the menu
    menu, _ := NewMenu(items)
    defer menu.Free()
    
    menuwin, _ := NewWindow(HEIGHT, WIDTH, 4, 14)
    menuwin.Keypad(true)
    
    menu.Window(menuwin)
    menu.SubWindow(menuwin.DerivedWindow(6, 38, 3, 1))
    menu.Format(5, 1)
    menu.Mark(" * ")
    
    // Print centered menu title
    //y, x := menuwin.Maxyx()
    title := "My Menu"
    menuwin.Box(0, 0)
    menuwin.ColorOn(1)
    menuwin.Print(1, (WIDTH/2)-(len(title)/2), title)
    menuwin.ColorOff(1)
    // FIXME: ACS_* definitions not currently available
    menuwin.HLine(2, 1, '-', WIDTH-2) 
    
    y, _ := stdscr.Maxyx()
    stdscr.ColorOn(2)
    stdscr.Print(y-3, 1, "Use up/down arrows or page up/down to navigate. 'q' to exit")
    stdscr.ColorOff(2)
    stdscr.Refresh()
    
    menu.Post()
    defer menu.UnPost()
    menuwin.Refresh()
    
    for {
        Update()
        ch, _ := menuwin.GetChar()
        
        if x := Key(ch); x == "q" {
            return
        } else {
            menu.Driver(x)
        }
    }
}
