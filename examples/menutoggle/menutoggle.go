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
    
    // build the menu items
    menu_items := []string{
        "Choice 1", 
        "Choice 2", 
        "Choice 3", 
        "Choice 4", 
        "Choice 5", 
        "Choice 6", 
        "Choice 7",
        "Exit"}
    items := make([]*MenuItem, len(menu_items))
    for i, val := range menu_items {
        items[i], _ = NewItem(val, "")
        defer items[i].Free()
    }
    
    // create the menu
    menu, _ := NewMenu(items)
    defer menu.Free()

    menu.Option(O_ONEVALUE, false)
    
    y, _ := stdscr.Maxyx()
    stdscr.Print(y-3, "Use up/down arrows to move, spacebar to toggle and " +
        "enter to print. 'q' to exit")
    stdscr.Refresh()
    
    menu.Post()
    defer menu.UnPost()
    
    for {
        Update()
        ch := stdscr.GetChar()

        switch ch {
        case 'q':
            return
        case ' ':
            menu.Driver(MD_TOGGLE)
        case KEY_RETURN:
            var list string
            for _, item := range menu.Items(nil) {
                if item.Value() {
                    list += "\"" + item.Name() + "\" "
                }
            }
            stdscr.Move(20, 0)
            stdscr.ClearToEOL()
            stdscr.Print(20, list)
            stdscr.Refresh()
        default:
            menu.Driver(DriverActions[ch])
        }
    }
}
