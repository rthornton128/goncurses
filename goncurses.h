void ncurses_getyx(WINDOW *win, int *y, int *x);
void ncurses_getmaxyx(WINDOW *win, int *y, int *x);
WINDOW * ncurses_wgetparent(const WINDOW *win);
bool ncurses_is_cleared(const WINDOW *win);
bool ncurses_is_keypad(const WINDOW *win);
bool ncurses_is_pad(const WINDOW *win);
bool ncurses_is_subwin(const WINDOW *win);
bool ncurses_has_mouse(void);
