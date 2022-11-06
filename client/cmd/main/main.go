package main

import (
	"log"

	"github.com/RGood/snooverse-client/pkg/pages"
	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

type page interface {
	Dimensions() (float64, float64)
	Update() error
	Draw(screen *ebiten.Image)
}

type game struct {
	pages         []page
	activePage    page
	screenBuffer  *ebiten.Image
	scalingBuffer *ebiten.DrawImageOptions
}

func (g *game) GotoPage(i int) {
	g.activePage = g.pages[i]
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.screenBuffer = ebiten.NewImage(outsideWidth, outsideHeight)

	g.scalingBuffer = &ebiten.DrawImageOptions{}
	px, py := g.activePage.Dimensions()
	g.scalingBuffer.GeoM.Scale(float64(outsideWidth)/px, float64(outsideHeight)/py)

	return outsideWidth, outsideHeight
}

func (g *game) Update() error {
	return g.activePage.Update()
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Clear()

	g.activePage.Draw(g.screenBuffer)

	screen.DrawImage(g.screenBuffer, g.scalingBuffer)
}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetMaxTPS(144)
	//ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowTitle("Snooversus")
	ebiten.SetWindowResizable(true)
	ebiten.SetScreenClearedEveryFrame(false)

	game := game{
		pages:      []page{},
		activePage: nil,
	}

	game.pages = []page{
		pages.MenuPage(pages.MenuPageConfig{
			StartButton: func() { game.GotoPage(1) },
		}),
		pages.LoadingPage(pages.LoadingPageConfig{
			Cancel: func() { game.GotoPage(0) },
		}),
	}

	game.GotoPage(0)

	err := ebiten.RunGame(&game)
	if err != nil {
		log.Print(err)
	}
}
