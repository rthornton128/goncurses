package goncurses

// #cgo LDFLAGS: -lmenu
// #include <menu.h>
import "C"

var MenuErrors = map[C.int] string {
    C.E_OK: "The routine succeeded.",
    C.E_BAD_ARGUMENT: "Routine detected an incorrect or out-of-range argument.",
    C.E_BAD_STATE: "Routine was called from an initialization or termination function.",
    C.E_NO_MATCH: "Character failed to match.",
    C.E_NO_ROOM: "Menu is too large for its window.",
    C.E_NOT_CONNECTED: "No items are connected to the menu.",
    C.E_NOT_POSTED: "The menu has not been posted.",
    C.E_NOT_SELECTABLE: "The designated item cannot be selected.",
    C.E_POSTED: "The menu is already posted.",
    C.E_REQUEST_DENIED:  "The menu driver could not process the request.",
    C.E_SYSTEM_ERROR:  "System error occurred (see errno).",
    C.E_UNKNOWN_COMMAND: "The menu driver code saw an unknown request code.",
}

type Menu C.MENU
type MenuItem C.ITEM

func NewMenu() *Menu {
    return nil
}

func (m *Menu) Free() {
    return
}
