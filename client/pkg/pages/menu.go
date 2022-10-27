package pages

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuPage struct {
	UnimplementedPage
	x, y int
	hovered bool
}

func NewMenuPage() *MenuPage {
	return &MenuPage{
		x: 50,
		y: 50,
		hovered: false,
	}
}

func (mp *MenuPage) Update() error {
	return nil
}

func (mp *MenuPage) Draw(screen *ebiten.Image) {
	screen.Clear()

	// menu box
	boxRect := image.Rectangle{
		Min: image.Point{50, 50},
		Max: image.Point{100, 100},
	}

	box := ebiten.NewImage(boxRect.Dx(), boxRect.Dy())
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(boxRect.Min.X), float64(boxRect.Min.Y))
	box.Fill(color.RGBA64{0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF})

	// cursor box
	cursorX, cursorY := ebiten.CursorPosition()

	mouseRect := image.Rectangle{
		Min: image.Point{cursorX, cursorY},
		Max: image.Point{cursorX + 1, cursorY + 1},
	}

	// for debugging purposes, the mouse box!
	// mouseBox := ebiten.NewImage(mouseRect.Dx(), mouseRect.Dy())
	// mouseBox.Fill(color.RGBA64{0x0, 0x0, 0xFFFF, 0xFFFF})
	// mouseOpts := &ebiten.DrawImageOptions{}
	// mouseOpts.GeoM.Translate(float64(mouseRect.Min.X), float64(mouseRect.Min.Y))
	// screen.DrawImage(mouseBox, mouseOpts)

	intersectRect := boxRect.Intersect(mouseRect)

	if (intersectRect.Dx() != 0) {
		box.Fill(color.RGBA64{0xFFFF, 0x0, 0x0, 0xFFFF})
	}

	if (ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)) {
		box.Fill(color.RGBA64{0x0, 0x0, 0xFFFF, 0xFFFF})
	}

	screen.DrawImage(box, opts)
}
