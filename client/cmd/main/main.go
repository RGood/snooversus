package main

import (
	"log"

	"github.com/RGood/snooverse-client/pkg/pages"
	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

type page interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type game struct {
	pages         []page
	activePage    page
	width, height int
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g *game) Update() error {
	return g.activePage.Update()
}

func (g *game) Draw(screen *ebiten.Image) {
	g.activePage.Draw(screen)
}

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetMaxTPS(144)
	//ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowTitle("Snooversus")
	ebiten.SetWindowResizable(true)
	ebiten.SetScreenClearedEveryFrame(false)

	game := game{
		pages:      []page{},
		// activePage: pages.NewLoadingPage(50, time.Second),
		activePage: pages.NewMenuPage(),
		width:      900,
		height:     800,
	}

	err := ebiten.RunGame(&game)
	if err != nil {
		log.Print(err)
	}
}
