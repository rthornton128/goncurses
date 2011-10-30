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
    InitPair(2, C_GREEN, C_BLACK)
    InitPair(3, C_MAGENTA, C_BLACK)
    
    // build the menu items
    menu_items := []string{
        "Choice 1", 
        "Choice 2", 
        "Choice 3", 
        "Choice 4", 
        "Choice 5",
        "Exit"}
    items := make([]*MenuItem, len(menu_items))
    for i, val := range menu_items {
        items[i], _ = NewItem(val, "")
        defer items[i].Free()
        
        if i == 2 || i == 4 {
            items[i].Selectable(false)
        }
    }
    
    // create the menu
    menu, _ := NewMenu(items)
    defer menu.Free()
    
    y, _ := stdscr.Maxyx()
    stdscr.Print(y-3, "Use up/down arrows to move; 'q' to exit")
    stdscr.Refresh()

    menu.Foreground(ColorPair(1)|A_REVERSE)
    menu.Background(ColorPair(2)|A_BOLD)
    menu.Grey(ColorPair(3)|A_BOLD)
    
    menu.Post()
    defer menu.UnPost()
    
    for {
        Update()
        ch := stdscr.GetChar()
        switch ch {
        case ' ':
            menu.Driver(MD_TOGGLE)
        case 'q':
            return
        case KEY_RETURN:
            stdscr.Move(20, 0)
            stdscr.ClearToEOL()
            stdscr.Print(20, 0, "Item selected is: %s", menu.Current(nil).Name())
            menu.PositionCursor()
        default:
            menu.Driver(DriverActions[ch])
        }
    }
}
