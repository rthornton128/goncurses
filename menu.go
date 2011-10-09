package goncurses

// #cgo LDFLAGS: -lmenu
// #include <menu.h>
// #include <stdlib.h>
import "C"

import (
    "os"
    "unsafe"
)

type Errno int

var menuerrors = map[Errno] string {
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

func (e Errno) String() string {
    return menuerrors[e]
}

var driveractions = map[string] C.int {
    "down": C.REQ_DOWN_ITEM,
    "first": C.REQ_FIRST_ITEM,
    "last": C.REQ_LAST_ITEM,
    "left": C.REQ_LEFT_ITEM,
    "next": C.REQ_NEXT_ITEM,
    "page down": C.REQ_SCR_DPAGE,
    "page up": C.REQ_SCR_UPAGE,
    "prev": C.REQ_PREV_ITEM,
    "right": C.REQ_RIGHT_ITEM,
    "scroll down": C.REQ_SCR_DLINE,
    "scroll up": C.REQ_SCR_ULINE,
    "toggle": C.REQ_TOGGLE_ITEM,
    "up": C.REQ_UP_ITEM,
}

type Menu C.MENU
type MenuItem C.ITEM

func NewMenu(items []*MenuItem) (*Menu, os.Error) {
    citems := make([]*C.ITEM, len(items))
    for index, item := range items {
        citems[index] = (*C.ITEM)(item)
    }
    menu, errno := C.new_menu((**C.ITEM)(&citems[0]))
    if menu == nil {
        return nil, errno
    }
    return (*Menu)(menu), nil
}

func (m *Menu) Driver(action string) {
    C.menu_driver((*C.MENU)(m), driveractions[action])
    return
}

func (m *Menu) Free() os.Error {
    if res := C.free_menu((*C.MENU)(m)); res != C.E_OK {
        return os.NewError(menuerrors[Errno(res)])
    }
    return nil
}

func (m *Menu) Post() os.Error {
    if res := C.post_menu((*C.MENU)(m)); res != C.E_OK {
        return os.NewError(menuerrors[Errno(res)])
    }
    return nil
}

func NewItem(name, desc string) (*MenuItem, os.Error) {
    cname := C.CString(name)    
    cdesc := C.CString(desc)
    
    item, err := C.new_item(cname, cdesc)
    if item == nil {
        return nil, err
    }
    return (*MenuItem)(item), nil
}

func (mi *MenuItem) Free() {
    C.free(unsafe.Pointer(C.item_name((*C.ITEM)(mi))))
    C.free_item((*C.ITEM)(mi))
    return
}

func (mi *MenuItem) Name() string {
    return C.GoString(C.item_name((*C.ITEM)(mi)))
}
