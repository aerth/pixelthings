package main

import (
	"errors"

	px "github.com/faiface/pixel"
)

type AppFunc func(*App) error

type Button struct {
	R  px.Rect
	P  px.Vec
	ID uint32
	Fn AppFunc
}

type buttonID uint32

func (b buttonID) Uint32() uint32 {
	return uint32(b)
}

//go:generate stringer -type buttonID -trimprefix Button .
const (
	ButtonNew buttonID = iota
	ButtonImport
	ButtonLoad
	ButtonSave
	ButtonRotate
	ButtonExit
	ButtonNextGif
	ButtonVSyncToggle
)

func nextGif(a *App) error {
	a.spritenum = (a.spritenum + 1) % len(a.sprites)
	a.sprite = a.sprites[a.spritenum]
	return nil
}

// AppFunc
func rotateMainSprite(a *App) error {
	x := a.rot
	if x == 0 {
		x = 361
	}
	a.rot = (x - 1)
	return nil
}
func leaveApp(a *App) error {

	return errors.New("user left")
}
func toggleVsync(a *App) error {
	a.win.SetVSync(!a.win.VSync())
	return nil
}

var menumain = [...]Button{
	// ButtonNew: Button{
	// 	px.R(0, 0, 100, 100), px.ZV, ButtonNew.Uint32(), Popup("new"),
	// },
	// ButtonImport: Button{
	// 	px.R(100, 0, 200, 100), px.ZV, ButtonNew.Uint32(), Popup("import"),
	// },
	ButtonVSyncToggle: Button{
		px.R(100, 0, 200, 100), px.ZV, ButtonNew.Uint32(), toggleVsync,
	},
	ButtonNextGif: Button{
		px.R(200, 0, 300, 100), px.ZV, ButtonNew.Uint32(), nextGif,
	},
	ButtonExit: Button{
		px.R(300, 0, 400, 100), px.ZV, ButtonNew.Uint32(), leaveApp,
	},
	ButtonRotate: Button{
		px.R(400, 0, 500, 100), px.ZV, ButtonNew.Uint32(), rotateMainSprite,
	},
}

func checkToolbar(mp px.Vec, wb px.Rect) AppFunc {
	for _, v := range menumain {
		if v.R.Contains(mp) {
			return v.Fn
		}
	}
	return nil
}

// dummies

// var _ AppFunc = Echo("foo")

// // Echo test AppFunc
// func Echo(s string) AppFunc {
// 	return func(a *App) error {
// 		log.Println("Got click from:", s)
// 		return nil
// 	}
// }

// // Popup message AppFunc
// func Popup(s string) AppFunc {
// 	return func(a *App) error {
// 		a.statusbar.Clear()
// 		fmt.Fprintf(a.statusbar, "%s", s)
// 		return nil
// 	}
// }
