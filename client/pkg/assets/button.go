package assets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button[T any] struct {
	image.Rectangle
	OnHover    func(*Button[T])
	OnClick    func(*Button[T])
	intercept  func(*T, ebiten.Image, ebiten.DrawImageOptions) (*ebiten.Image, *ebiten.DrawImageOptions)
	cachedImg  *ebiten.Image
	cachedOpts *ebiten.DrawImageOptions
}

type ButtonOpt[T any] interface {
	Apply(*Button[T]) *Button[T]
}

type buttonFillColor[T any] struct {
	c color.Color
}

func (bfc *buttonFillColor[T]) Apply(btn *Button[T]) *Button[T] {
	btn.cachedImg.Fill(bfc.c)
	return btn
}

func ButtonFillColor[T any](c color.Color) ButtonOpt[T] {
	return &buttonFillColor[T]{c}
}

type buttonText[T any] struct {
	text string
	c    color.Color
	x, y int
	font font.Face
}

func (btc *buttonText[T]) Apply(btn *Button[T]) *Button[T] {
	text.Draw(btn.cachedImg, btc.text, btc.font, btc.x, btc.y, btc.c)
	return btn
}

func ButtonText[T any](text string, c color.Color, x, y int, font font.Face) *buttonText[T] {
	return &buttonText[T]{text, c, x, y, font}
}

type buttonIntercept[T any] struct {
	intercept func(*T, ebiten.Image, ebiten.DrawImageOptions) (*ebiten.Image, *ebiten.DrawImageOptions)
}

func (bi *buttonIntercept[T]) Apply(btn *Button[T]) *Button[T] {
	btn.intercept = bi.intercept
	return btn
}

func ButtonIntercept[T any](intercept func(*T, ebiten.Image, ebiten.DrawImageOptions) (*ebiten.Image, *ebiten.DrawImageOptions)) *buttonIntercept[T] {
	return &buttonIntercept[T]{
		intercept: intercept,
	}
}

func NewButton[T any](bounds image.Rectangle, buttonOpts ...ButtonOpt[T]) *Button[T] {
	cm := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	co := &ebiten.DrawImageOptions{}
	co.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
	btn := &Button[T]{
		bounds,
		nil,
		nil,
		nil,
		cm,
		co,
	}

	for _, btnOpt := range buttonOpts {
		btn = btnOpt.Apply(btn)
	}

	return btn
}

func (b *Button[T]) Draw(state *T, img *ebiten.Image) {
	if b.intercept != nil {
		img.DrawImage(b.intercept(state, *b.cachedImg, *b.cachedOpts))
	} else {
		img.DrawImage(b.cachedImg, b.cachedOpts)
	}

}
