package pages

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type LoadingPage struct {
	UnimplementedPage
	rotation float64
	size     int
	period   time.Duration
	lastTime time.Time
}

func NewLoadingPage(size int, period time.Duration) *LoadingPage {
	return &LoadingPage{
		rotation: 0,
		size:     size,
		period:   period,
		lastTime: time.Now(),
	}
}

func (lp *LoadingPage) Update() error {
	lp.rotation += (float64(2*math.Pi) * float64(time.Since(lp.lastTime)) / float64(lp.period))
	lp.lastTime = time.Now()
	return nil
}

func (lp *LoadingPage) Draw(screen *ebiten.Image) {
	screen.Clear()
	x, y := ebiten.WindowSize()

	spinner := ebiten.NewImage(lp.size, lp.size)
	spinner.Fill(color.RGBA64{0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(-lp.size/2), float64(-lp.size/2))
	opts.GeoM.Rotate(lp.rotation)
	opts.GeoM.Translate(float64(x/2), float64(y/2))

	screen.DrawImage(spinner, opts)
}
