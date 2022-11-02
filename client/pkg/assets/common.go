package assets

import "github.com/hajimehoshi/ebiten/v2"

type Drawable[T any] interface {
	Draw(*T, *ebiten.Image)
}
