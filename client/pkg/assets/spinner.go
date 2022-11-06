package assets

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Spinner[T any] struct {
	image.Rectangle
	img         *ebiten.Image
	start       time.Time
	period      time.Duration
	curRotation float64
	cachedOpts  *ebiten.DrawImageOptions
}

func rotate(img *ebiten.Image, opts *ebiten.DrawImageOptions, angle float64) *ebiten.DrawImageOptions {
	ix := float64(img.Bounds().Dx() / 2)
	iy := float64(img.Bounds().Dy() / 2)

	opts.GeoM.Translate(-ix, -iy)
	opts.GeoM.Rotate(angle)
	opts.GeoM.Translate(ix, iy)

	return opts
}

func NewSpinner[T any](x, y int, img *ebiten.Image, period time.Duration) *Spinner[T] {
	bounds := image.Rect(x, y, x+img.Bounds().Dx(), y+img.Bounds().Dy())

	return &Spinner[T]{
		bounds, img, time.Now(), period, 0, &ebiten.DrawImageOptions{},
	}
}

func (s *Spinner[T]) Update() {
	diff := time.Since(s.start)
	rotation := float64(diff%s.period) / float64(s.period)
	s.cachedOpts = rotate(s.img, &ebiten.DrawImageOptions{}, rotation*math.Pi*2)
	s.cachedOpts.GeoM.Translate(float64(s.Bounds().Min.X), float64(s.Bounds().Min.Y))
}

func (s *Spinner[T]) Draw(state *T, img *ebiten.Image) {
	img.DrawImage(s.img, s.cachedOpts)
}
