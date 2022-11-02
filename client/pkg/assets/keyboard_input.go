package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	pressCallbacks   map[ebiten.Key]func()
	releaseCallbacks map[ebiten.Key]func()
	pressed          map[ebiten.Key]struct{}
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		map[ebiten.Key]func(){},
		map[ebiten.Key]func(){},
		map[ebiten.Key]struct{}{},
	}
}

func (kb *Keyboard) SetOnPress(btn ebiten.Key, cb func()) {
	kb.pressCallbacks[btn] = cb
}

func (kb *Keyboard) ClearOnPress(btn ebiten.Key) {
	delete(kb.pressCallbacks, btn)
}

func (kb *Keyboard) SetOnRelease(btn ebiten.Key, cb func()) {
	kb.releaseCallbacks[btn] = cb
}

func (kb *Keyboard) ClearOnRelease(btn ebiten.Key) {
	delete(kb.releaseCallbacks, btn)
}

func (kb *Keyboard) IsPressed(btn ebiten.Key) bool {
	return ebiten.IsKeyPressed(btn)
}

func (kb *Keyboard) Update() {
	for key, cb := range kb.pressCallbacks {
		if inpututil.IsKeyJustPressed(key) {
			kb.pressed[key] = struct{}{}
			cb()
		}
	}

	for key, cb := range kb.releaseCallbacks {
		if inpututil.IsKeyJustReleased(key) {
			delete(kb.pressed, key)
			cb()
		}
	}
}
