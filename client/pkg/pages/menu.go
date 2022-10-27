package pages

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font/basicfont"
)

type MenuPage struct {
	UnimplementedPage
	fillColor color.RGBA64
	box image.Rectangle
	isMouseOverBox bool
	onclick func()
}

func NewMenuPage(onclick func()) *MenuPage {
	return &MenuPage{
		fillColor: color.RGBA64{0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF},
		box: image.Rectangle{
			Min: image.Point{50, 50},
			Max: image.Point{100, 100},
		},
		isMouseOverBox: false,
		onclick: onclick,
	}
}

func (mp *MenuPage) Update() error {

	// cursor box
	cursorX, cursorY := ebiten.CursorPosition()

	mouseRect := image.Rectangle{
		Min: image.Point{cursorX, cursorY},
		Max: image.Point{cursorX + 1, cursorY + 1},
	}

	intersectRect := mp.box.Intersect(mouseRect)
	mp.isMouseOverBox = intersectRect.Dx() != 0

	if (mp.isMouseOverBox && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)) {
		mp.onclick()
	}
	return nil
}

func (mp *MenuPage) Draw(screen *ebiten.Image) {
	screen.Clear()

	box := ebiten.NewImage(mp.box.Dx(), mp.box.Dy())
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(mp.box.Min.X), float64(mp.box.Min.Y))
	// a little transparent
	box.Fill(color.RGBA64{mp.fillColor.R, mp.fillColor.G, mp.fillColor.B, mp.fillColor.A/2})

	// for debugging purposes, the mouse box!
	// mouseBox := ebiten.NewImage(mouseRect.Dx(), mouseRect.Dy())
	// mouseBox.Fill(color.RGBA64{0x0, 0x0, 0xFFFF, 0xFFFF})
	// mouseOpts := &ebiten.DrawImageOptions{}
	// mouseOpts.GeoM.Translate(float64(mouseRect.Min.X), float64(mouseRect.Min.Y))
	// screen.DrawImage(mouseBox, mouseOpts)

	if (mp.isMouseOverBox) {
		box.Fill(mp.fillColor)
	}

	if (mp.isMouseOverBox && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)) {
		box.Fill(color.RGBA64{mp.fillColor.R, mp.fillColor.G, mp.fillColor.B, mp.fillColor.A/3})
	}

	text.Draw(box, "Hello", basicfont.Face7x13, 5, 20, color.RGBA64{0x0, 0x0, 0x0, 0xFFFF})
	screen.DrawImage(box, opts)
}
