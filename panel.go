package goncurses

// #cgo LDFLAGS: -lpanel
// #include <panel.h>
// #include <ncurses.h>
import "C"

type Panel C.PANEL

func NewPanel(w *Window) *Panel {
    return (*Panel)(C.new_panel((*C.WINDOW)(w)))
}

func UpdatePanels() {
    C.update_panels()
    return
}
