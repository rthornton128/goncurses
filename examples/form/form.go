/* This simple example mirrors the "hello world" TLDP ncurses howto */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "goncurses"//".googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    Echo(false)
    CBreak(true)
    stdscr.Keypad(true)
    
    fields := make([]*Field, 2)
    fields[0], _ = NewField(1, 10, 4, 18, 0, 0)
    defer fields[0].Free()
    fields[0].Background(A_UNDERLINE)
    fields[0].Options(FO_AUTOSKIP, false)
    
    fields[1], _ = NewField(1, 10, 6, 18, 0, 0)
    defer fields[1].Free()
    fields[1].Background(A_UNDERLINE)
    fields[1].Options(FO_AUTOSKIP, false)
    
    form, _ := NewForm(fields)
    form.Post()
    defer form.UnPost()
    defer form.Free()
    stdscr.Refresh()
    
    stdscr.Print(4, 10, "Value 1:")
    stdscr.Print(6, 10, "Value 2:")
    stdscr.Refresh()
    
    ch := stdscr.GetChar()
    for ch != 'q' {
        switch (ch) {
        case KEY_DOWN:
            form.Driver(REQ_NEXT_FIELD)
            form.Driver(REQ_END_LINE)
        case KEY_UP:
            form.Driver(REQ_PREV_FIELD)
            form.Driver(REQ_END_LINE)
        default:
            form.Driver(ch)
        }
        ch = stdscr.GetChar()
    }
}
