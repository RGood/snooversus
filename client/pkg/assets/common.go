package assets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable[T any] interface {
	Draw(*T, *ebiten.Image)
}

type Clickable interface {
	Click()
	Hover()
	Reset()
	Bounds() image.Rectangle
}

type Updateable interface {
	Update()
}
