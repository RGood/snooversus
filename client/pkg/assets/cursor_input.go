package assets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cursor struct {
	clickCallbacks   map[ebiten.MouseButton]func(*Cursor)
	releaseCallbacks map[ebiten.MouseButton]func(*Cursor)
}

func NewCursor() *Cursor {
	return &Cursor{
		map[ebiten.MouseButton]func(*Cursor){},
		map[ebiten.MouseButton]func(*Cursor){},
	}
}

func (c *Cursor) Overlaps(r image.Rectangle) bool {
	cx, cy := ebiten.CursorPosition()
	cursorRect := &image.Rectangle{
		Min: image.Point{cx, cy},
		Max: image.Point{cx + 1, cy + 1},
	}

	overlap := cursorRect.Intersect(r)
	return overlap.Dx() > 0
}

func (c *Cursor) SetOnClick(btn ebiten.MouseButton, cb func(*Cursor)) {
	c.clickCallbacks[btn] = cb
}

func (c *Cursor) ClearClick(btn ebiten.MouseButton) {
	delete(c.clickCallbacks, btn)
}

func (c *Cursor) SetOnRelease(btn ebiten.MouseButton, cb func(*Cursor)) {
	c.releaseCallbacks[btn] = cb
}

func (c *Cursor) ClearRelease(btn ebiten.MouseButton) {
	delete(c.releaseCallbacks, btn)
}

func (c *Cursor) IsMouseButtonPressed(btn ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(btn)
}

func (c *Cursor) Update() {
	for key, cb := range c.clickCallbacks {
		if inpututil.IsMouseButtonJustPressed(key) {
			cb(c)
		}
	}

	for key, cb := range c.releaseCallbacks {
		if inpututil.IsMouseButtonJustReleased(key) {
			cb(c)
		}
	}
}
