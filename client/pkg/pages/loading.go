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
	r, g, b  *oscilator
	adding   bool
}

type oscilator struct {
	val       uint16
	period    time.Duration
	startTime time.Time
}

func NewOscilator(period time.Duration) *oscilator {
	return &oscilator{
		val:       0,
		period:    period,
		startTime: time.Now(),
	}
}

func (o *oscilator) increment() {
	diff := time.Since(o.startTime)
	count := int(diff / o.period)
	if count%2 == 0 {
		o.val = uint16(float64(0xFFFF) * float64(diff%o.period) / float64(o.period))
	} else {
		o.val = 0xFFFF - uint16(float64(0xFFFF)*float64(diff%o.period)/float64(o.period))
	}
}

func NewLoadingPage(size int, period time.Duration) *LoadingPage {
	return &LoadingPage{
		rotation: 0,
		size:     size,
		period:   period,
		lastTime: time.Now(),
		r:        NewOscilator(time.Second * 2),
		g:        NewOscilator(time.Second * 3),
		b:        NewOscilator(time.Second * 5),
		adding:   true,
	}
}

var counter = 0

func (lp *LoadingPage) Update() error {
	counter++
	lp.rotation += (float64(2*math.Pi) * float64(time.Since(lp.lastTime)) / float64(lp.period))
	lp.lastTime = time.Now()

	lp.r.increment()
	lp.g.increment()
	lp.b.increment()

	return nil
}

func (lp *LoadingPage) Draw(screen *ebiten.Image) {
	screen.Clear()
	x, y := ebiten.WindowSize()

	spinner := ebiten.NewImage(lp.size, lp.size)
	spinner.Fill(color.RGBA64{lp.r.val, lp.g.val, lp.b.val, 0xFFFF})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(-lp.size/2), float64(-lp.size/2))
	opts.GeoM.Rotate(lp.rotation)
	opts.GeoM.Translate(float64(x/2), float64(y/2))

	screen.DrawImage(spinner, opts)
}
