package main

import (
	gc "code.google.com/p/goncurses"

	"math/rand"
	"time"
)

const density = 0.05
const planet_density = 0.001

var ship_ascii = []string{
	` ,`,
	` |\-`,
	`>|^===0`,
	` |/-`,
	` '`,
}

func genStarfield(pl, pc int) *gc.Pad {
	pad := gc.NewPad(pl, pc)
	stars := int(float64(pc*pl) * density)
	planets := int(float64(pc*pl) * planet_density)
	for i := 0; i < stars; i++ {
		y, x := rand.Intn(pl), rand.Intn(pc)
		c := int16(rand.Intn(4) + 1)
		pad.AttrOn(gc.A_BOLD | gc.ColorPair(c))
		pad.MovePrint(y, x, ".")
		pad.AttrOff(gc.A_BOLD | gc.ColorPair(c))
	}
	for i := 0; i < planets; i++ {
		y, x := rand.Intn(pl), rand.Intn(pc)
		c := int16(rand.Intn(2) + 5)
		pad.ColorOn(c)
		if i%2 == 0 {
			pad.MoveAddChar(y, x, 'O')
		}
		pad.MoveAddChar(y, x, 'o')
		pad.ColorOff(c)
	}
	return pad
}

func handleInput(stdscr *gc.Window, ship *Ship) bool {
	lines, cols := stdscr.MaxYX()
	y, x := ship.YX()
	k := stdscr.GetChar()

	switch byte(k) {
	case 0:
		break
	case 'a':
		x--
		if x < 2 {
			x = 2
		}
	case 'd':
		x++
		if x > cols/2 {
			x = cols / 2
		}
	case 's':
		y++
		if y > lines-5 {
			y = lines - 5
		}
	case 'w':
		y--
		if y < 5 {
			y = 5
		}
	case ' ':
		w1, _ := gc.NewWindow(1, 1, y+1, x+4)
		w2, _ := gc.NewWindow(1, 1, y+3, x+4)
		w1.AttrOn(gc.A_BOLD | gc.ColorPair(4))
		w2.AttrOn(gc.A_BOLD | gc.ColorPair(4))
		w1.Print("-")
		w2.Print("-")
		objects = append(objects, &Bullet{&w1, true})
		objects = append(objects, &Bullet{&w2, true})
	default:
		return false
	}
	ship.MoveWindow(y, x)
	return true
}

type Object interface {
	Cleanup()
	Collide(int)
	Draw(*gc.Window)
	Expired(int, int) bool
	Update()
}

type Asteroid struct {
	*gc.Window
	alive  bool
	y, x   int
	sy, sx int
}

var speeds = []int{-75, -50, -25, -10, 0, 10, 25, 50, 75}

func spawnAsteroid(my, mx int) {
	var y, x, sy, sx int
	switch rand.Intn(4) {
	case 0:
		y, x = 1, rand.Intn(mx-2)+1
		sy, sx = speeds[5:][rand.Intn(4)], speeds[rand.Intn(9)]
	case 1:
		y, x = rand.Intn(my-2)+1, 1
		sy, sx = speeds[rand.Intn(9)], speeds[5:][rand.Intn(4)]
	case 2:
		y, x = rand.Intn(my-2)+1, mx-2
		sy, sx = speeds[rand.Intn(9)], speeds[rand.Intn(4)]
	case 3:
		y, x = my-2, rand.Intn(mx-2)+1
		sy, sx = speeds[rand.Intn(4)], speeds[rand.Intn(9)]
	}
	w, _ := gc.NewWindow(1, 1, y, x)
	a := &Asteroid{Window: &w, alive: true, sy: sy, sx: sx, y: y * 100,
		x: x * 100}
	a.ColorOn(2)
	a.Print("@")
	objects = append(objects, a)
}

func (a *Asteroid) Cleanup() {
	a.Delete()
}

func (a *Asteroid) Collide(i int) {
}

func (a *Asteroid) Draw(w *gc.Window) {
	if !a.alive {
		objects = append(objects, newExplosion(a.YX()))
	}
	w.Overlay(a.Window)
}

func (a *Asteroid) Expired(my, mx int) bool {
	y, x := a.YX()
	if x <= 0 || x >= mx-1 || y <= 0 || y >= my-1 || !a.alive {
		return true
	}
	return false
}

func (a *Asteroid) Update() {
	a.y += a.sy
	a.x += a.sx
	a.MoveWindow(a.y/100, a.x/100)
}

type Bullet struct {
	*gc.Window
	alive bool
}

func (b *Bullet) Cleanup() {
	b.Delete()
}

func (b *Bullet) Collide(i int) {
	for k, v := range objects {
		if k == i {
			continue
		}
		switch t := v.(type) {
		case *Asteroid:
			ay, ax := t.YX()
			by, bx := b.YX()
			if ay == by && ax == bx {
				t.alive = false
				b.alive = false
			}
		}
	}
}

func (b *Bullet) Draw(w *gc.Window) {
	w.Overlay(b.Window)
}

func (b *Bullet) Expired(my, mx int) bool {
	_, x := b.YX()
	if x >= mx-1 || !b.alive {
		return true
	}
	return false
}

func (b *Bullet) Update() {
	y, x := b.YX()
	b.MoveWindow(y, x+1)
}

type Explosion struct {
	*gc.Window
	life int
}

func newExplosion(y, x int) *Explosion {
	w, _ := gc.NewWindow(3, 3, y-1, x-1)
	w.ColorOn(4)
	w.MovePrint(0, 0, `\ /`)
	w.AttrOn(gc.A_BOLD)
	w.MovePrint(1, 0, ` X `)
	w.AttrOn(gc.A_DIM)
	w.MovePrint(2, 0, `/ \`)
	return &Explosion{&w, 5}
}

func (e *Explosion) Cleanup() {
	e.Delete()
}

func (e *Explosion) Collide(i int) {}

func (e *Explosion) Draw(w *gc.Window) {
	w.Overlay(e.Window)
}

func (e *Explosion) Expired(y, x int) bool {
	return e.life <= 0
}

func (e *Explosion) Update() {
	e.life--
}

type Ship struct {
	*gc.Window
	life int
}

func newShip(y, x int) *Ship {
	w, _ := gc.NewWindow(5, 7, y, x)
	for i := 0; i < len(ship_ascii); i++ {
		w.MovePrint(i, 0, ship_ascii[i])
	}
	return &Ship{&w, 5}
}

func (s *Ship) Cleanup() {
	s.Delete()
}

func (s *Ship) Collide(i int) {
	ty, tx := s.YX()
	by, bx := s.MaxYX()
	for _, ob := range objects {
		if a, ok := ob.(*Asteroid); ok {
			ay, ax := a.YX()
			if ay >= ty && ay+1 <= ty+by && ax >= tx && ax+1 <= tx+bx {
				a.alive = false
				s.life--
			}
		}
	}
}

func (s *Ship) Draw(w *gc.Window) {
	w.Overlay(s.Window)
}

func (s *Ship) Expired(y, x int) bool {
	return s.life <= 0
}

func (s *Ship) Update() {}

var objects = make([]Object, 0, 16)

func updateObjects(my, mx int) {
	tmp := make([]Object, 0, 16)
	for i, ob := range objects {
		ob.Update()
		ob.Collide(i)
		if ob.Expired(my, mx) {
			ob.Cleanup()
		} else {
			tmp = append(tmp, ob)
		}
	}
	objects = tmp
}

func drawObjects(s *gc.Window) {
	for _, ob := range objects {
		ob.Draw(s)
	}
}

func lifeToText(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "*"
	}
	return s
}

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	rand.Seed(time.Now().Unix())
	gc.StartColor()
	gc.Cursor(0)
	gc.Echo(false)
	gc.Raw(true)

	gc.InitPair(1, gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(2, gc.C_YELLOW, gc.C_BLACK)
	gc.InitPair(3, gc.C_MAGENTA, gc.C_BLACK)
	gc.InitPair(4, gc.C_RED, gc.C_BLACK)

	gc.InitPair(5, gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(6, gc.C_GREEN, gc.C_BLACK)

	stdscr.Timeout(0)

	lines, cols := stdscr.MaxYX()
	pl, pc := lines, cols*3

	ship := newShip(lines/2, 5)
	objects = append(objects, ship)

	field := genStarfield(pl, pc)
	text := stdscr.Duplicate()

	c := time.NewTicker(time.Second / 2)
	c2 := time.NewTicker(time.Second / 16)
	px := 0

loop:
	for {
		text.MovePrintf(0, 0, "Life: [%-5s]", lifeToText(ship.life))
		text.MovePrintf(1, 0, "Number of objects: %3d", len(objects))
		stdscr.Clear()
		stdscr.Copy(field.Window, 0, px, 0, 0, lines-1, cols-1, true)
		drawObjects(&stdscr)
		stdscr.Overlay(&text)
		stdscr.Refresh()
		select {
		case <-c.C:
			spawnAsteroid(stdscr.MaxYX())
			if px+cols >= pc {
				break loop
			}
			px++
		case <-c2.C:
			updateObjects(stdscr.MaxYX())
			drawObjects(&stdscr)
		default:
			if !handleInput(&stdscr, ship) || ship.Expired(-1, -1) {
				break loop
			}
		}
	}
}
