package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aerth/pixelthings/pixelgif/assets"
	"github.com/aerth/spriteutil"

	px "github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	gl "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type App struct {
	cfg gl.WindowConfig
	win *gl.Window

	dt float64

	statusbar    *text.Text
	buttontxt    map[int]*text.Text
	imd          *imdraw.IMDraw             // for buttons
	sprite       *spriteutil.AnimatedSprite // chosen sprite from sprites
	sprites      []*spriteutil.AnimatedSprite
	matrixrotate px.Matrix
	spritenum    int // choosing sprite from sprites
	rot          int // 360
}

func NewApp(cfg gl.WindowConfig) *App {
	return &App{cfg: cfg, buttontxt: make(map[int]*text.Text)}
}
func MkFont(name string, size float64) *text.Text {
	if size == 0 {
		size = 18.0
	}
	if name == "" {
		name = "computer-font.ttf"
	}
	file, err := assets.Assets.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	t := spriteutil.LoadTTF(file, size, px.ZV)
	return t
}
func (a *App) run() {
	checkerr := func(err error) {
		if err != nil {
			log.Fatalln("fatal:", err)
		}
	}
	win, err := gl.NewWindow(a.cfg)
	checkerr(err)
	a.win = win
	checkerr(a.setup())
	var (
		fps    = new(int)
		dt     = new(float64)
		last   = time.Now()
		second = time.Tick(time.Second)
	)
	for !win.Closed() {
		*dt = time.Since(last).Seconds()
		last = time.Now()
		*fps++
		select {
		default:
		case <-second:
			win.SetTitle(fmt.Sprintf("dt=%2.0fms, %d FPS", 1000**dt, *fps))
			*fps = 0
			*dt = 0
			a.rot = (a.rot + 1) % 360
		}
		checkerr(a.dpad())
		checkerr(a.update(*dt))
		win.Clear(colornames.Black)
		checkerr(a.draw())
		win.Update()
	}
}
func (a *App) draw() error {
	a.sprite.Draw(a.win, a.matrixrotate)
	a.imd.Draw(a.win)
	for i, v := range a.buttontxt {
		mat := px.IM.Moved(menumain[i].R.Center().Add(px.V(6+-menumain[i].R.W()/2, 0)))
		v.Draw(a.win, mat)
	}
	a.statusbar.Draw(a.win, px.IM.Moved(px.V(10, a.win.Bounds().H()-20)))
	return nil
}

func (a *App) setup() error {
	a.imd = imdraw.New(nil)
	a.statusbar = MkFont("", 18.0)
	fmt.Fprintf(a.statusbar, "Animated GIF, Rotating")
	a.imd.Color = colornames.Limegreen
	for nam, button := range menumain {
		if button.Fn == nil {
			continue
		}
		a.buttontxt[nam] = MkFont("", 18.0)
		fmt.Fprintf(a.buttontxt[nam], "%s", buttonID(nam).String())
		a.buttontxt[nam].Color = colornames.Black
		a.imd.Push(button.R.Min)
		a.imd.Push(button.R.Max)
		a.imd.Rectangle(1)
	}
	for _, filename := range []string{"allyourbase.gif", "select1.gif", "simpsons.gif"} {

		file, err := assets.Assets.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		sprite, err := spriteutil.LoadGif(file)
		if err != nil {
			return err
		}

		a.sprites = append(a.sprites, sprite)
	}
	if len(a.sprites) == 0 {
		return fmt.Errorf("couldnt load images")
	}
	a.sprite = a.sprites[0]
	return nil
}

func (a *App) dpad() error {
	w := a.win

	if w.JustPressed(gl.MouseButtonLeft) {
		mousepos := w.MousePosition()
		winbounds := w.Bounds()
		log.Println("mouseposition:", mousepos)
		if fn := checkToolbar(mousepos, winbounds); fn != nil {
			log.Println("running func")
			if err := fn(a); err != nil {
				return err
			}
		}
	}
	if w.JustPressed(gl.KeyEscape) || (w.JustPressed(gl.KeyQ) && w.Pressed(gl.KeyLeftControl)) {
		return errors.New("user escaped")
	}
	return nil
}

func (a *App) update(dt float64) error {
	// a.dt += dt
	// if a.dt > 1 {
	// 	a.dt = 0
	// 	a.rot = (a.rot + 1) % 360
	// }
	a.matrixrotate = px.IM.Rotated(px.ZV, float64(a.rot)).Moved(a.win.Bounds().Center())
	a.sprite.Update(dt) // progress gif animation
	return nil
}
