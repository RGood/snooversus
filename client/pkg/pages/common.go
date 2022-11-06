package pages

import "github.com/hajimehoshi/ebiten/v2"

type Page interface {
	Update() error
	Draw(*ebiten.Image)
	Dimensions() (float64, float64)
	mustEmbedUnimplementedPage()
}

type UnimplementedPage struct {
}

func (up *UnimplementedPage) mustEmbedUnimplementedPage() {}

func (up *UnimplementedPage) Update() error {
	panic("update not implemented")
}

func (up *UnimplementedPage) Draw(_ *ebiten.Image) {
	panic("draw not implemented")
}
