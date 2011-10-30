/* This simple example mirrors the "hello world" TLDP ncurses howto */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    Echo(false)
    CBreak(true)
    StartColor()
    stdscr.Keypad(true)
    
    InitPair(1, C_WHITE, C_BLUE)
    InitPair(2, C_YELLOW, C_BLUE)
    
    fields := make([]*Field, 2)
    fields[0], _ = NewField(1, 10, 4, 18, 0, 0)
    defer fields[0].Free()
    fields[0].Foreground(ColorPair(1))
    fields[0].Background(ColorPair(2)|A_UNDERLINE|A_BOLD)
    fields[0].Options(FO_AUTOSKIP, false)
    
    fields[1], _ = NewField(1, 10, 6, 18, 0, 0)
    defer fields[1].Free()
    fields[1].Foreground(ColorPair(1))
    fields[1].Background(A_UNDERLINE)
    fields[1].Options(FO_AUTOSKIP, false)
    fields[1].Pad('*')
    
    form, _ := NewForm(fields)
    form.Post()
    defer form.UnPost()
    defer form.Free()
    stdscr.Refresh()
    
    //stdscr.ColorOn(ColorPair(2))
    stdscr.AttrOn(ColorPair(2)|A_BOLD)
    stdscr.Print(4, 10, "Value 1:")
    stdscr.AttrOff(ColorPair(2)|A_BOLD)
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
