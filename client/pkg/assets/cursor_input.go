package assets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cursor struct {
	clickables []Clickable
	hovered    map[Clickable]struct{}
	cr         image.Rectangle
}

func NewCursor() *Cursor {
	return &Cursor{
		[]Clickable{},
		map[Clickable]struct{}{},
		image.Rectangle{},
	}
}

func (c *Cursor) Overlaps() map[Clickable]struct{} {
	cx, cy := ebiten.CursorPosition()
	c.cr.Min.X = cx
	c.cr.Min.Y = cy
	c.cr.Max.X = cx + 1
	c.cr.Max.Y = cy + 1

	toReturn := map[Clickable]struct{}{}

	for _, clickable := range c.clickables {
		if c.cr.Intersect(clickable.Bounds()).Dx() > 0 {
			toReturn[clickable] = struct{}{}
		}
	}

	return toReturn
}

func (c *Cursor) IsMouseButtonPressed(btn ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(btn)
}

func (c *Cursor) RegisterClickable(clickable Clickable) {
	c.clickables = append(c.clickables, clickable)
}

func (c *Cursor) Update() {
	overlaps := c.Overlaps()

	for clickable, _ := range overlaps {
		if _, ok := c.hovered[clickable]; !ok {
			c.hovered[clickable] = struct{}{}
			clickable.Hover()
		}
	}

	for hovered, _ := range c.hovered {
		if _, ok := overlaps[hovered]; !ok {
			hovered.Reset()
		}
	}

	c.hovered = overlaps

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for clickable, _ := range overlaps {
			clickable.Click()
		}
	}
}
