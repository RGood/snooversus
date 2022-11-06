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
	opts       []ButtonOpt[T]
	color      color.Color
	text       string
	textColor  color.Color
	textPos    image.Point
	textFont   font.Face
	OnHover    func(*Button[T])
	OnClick    func(*Button[T])
	intercept  func(*T, ebiten.Image, ebiten.DrawImageOptions) (*ebiten.Image, *ebiten.DrawImageOptions)
	CachedImg  *ebiten.Image
	CachedOpts *ebiten.DrawImageOptions
}

type ButtonOpt[T any] interface {
	Apply(*Button[T]) *Button[T]
}

type buttonFillColor[T any] struct {
	c color.Color
}

func (bfc *buttonFillColor[T]) Apply(btn *Button[T]) *Button[T] {
	btn.CachedImg.Fill(bfc.c)
	btn.color = bfc.c
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
	btn.text = btc.text
	btn.textColor = btc.c
	btn.textPos = image.Pt(btc.x, btc.y)
	btn.textFont = btc.font
	text.Draw(btn.CachedImg, btc.text, btc.font, btc.x, btc.y, btc.c)
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

type buttonOnClick[T any] struct {
	oc func(*Button[T])
}

func (boc *buttonOnClick[T]) Apply(btn *Button[T]) *Button[T] {
	btn.OnClick = boc.oc
	return btn
}

func ButtonOnClick[T any](oc func(*Button[T])) ButtonOpt[T] {
	return &buttonOnClick[T]{oc}
}

type buttonOnHover[T any] struct {
	oh func(*Button[T])
}

func (boh *buttonOnHover[T]) Apply(btn *Button[T]) *Button[T] {
	btn.OnHover = boh.oh
	return btn
}

func ButtonOnHover[T any](oc func(*Button[T])) ButtonOpt[T] {
	return &buttonOnHover[T]{oc}
}

func NewButton[T any](x1, x2, y1, y2 int, buttonOpts ...ButtonOpt[T]) *Button[T] {
	bounds := image.Rect(x1, y1, x2, y2)
	cm := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	co := &ebiten.DrawImageOptions{}
	co.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
	btn := &Button[T]{
		bounds,
		buttonOpts,
		nil,
		"",
		nil,
		image.Point{},
		nil,
		nil,
		nil,
		nil,
		cm,
		co,
	}

	for _, btnOpt := range buttonOpts {
		btn = btnOpt.Apply(btn)
	}

	btn.Reset()

	return btn
}

func (b *Button[T]) redraw() {
	b.CachedImg.Clear()

	if b.color != nil {
		b.CachedImg.Fill(b.color)
	}

	if b.text != "" {
		text.Draw(b.CachedImg, b.text, b.textFont, b.textPos.X, b.textPos.Y, b.textColor)
	}
}

func (b *Button[T]) SetFill(c color.Color) {
	b.color = c
}

func (b *Button[T]) SetText(text string, f font.Face, x, y int, c color.Color) {
	b.text = text
	b.textFont = f
	b.textPos = image.Pt(x, y)
	b.textColor = c
}

func (b *Button[T]) Draw(state *T, img *ebiten.Image) {
	if b.intercept != nil {
		img.DrawImage(b.intercept(state, *b.CachedImg, *b.CachedOpts))
	} else {
		img.DrawImage(b.CachedImg, b.CachedOpts)
	}
}

func (b *Button[T]) Click() {
	if b.OnClick != nil {
		b.OnClick(b)
	}
}

func (b *Button[T]) Hover() {
	if b.OnHover != nil {
		b.OnHover(b)
	}
	b.redraw()
}

func (b *Button[T]) Reset() {
	b.CachedImg = ebiten.NewImage(b.Dx(), b.Dy())
	for _, opt := range b.opts {
		b = opt.Apply(b)
	}
}
