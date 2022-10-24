package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"

	"image/color"
	_ "image/png"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
)

type pageContainer struct {
	widget    widget.PreferredSizeLocateableWidget
	titleText *widget.Text
	flipBook  *widget.FlipBook
}

type game struct {
	ui *ebitenui.UI
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *game) Update() error {
	g.ui.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowTitle("Ebiten UI Demo")
	ebiten.SetWindowResizable(true)
	ebiten.SetScreenClearedEveryFrame(false)

	ui, err := createUI()

	game := game{
		ui: ui,
	}

	err = ebiten.RunGame(&game)
	if err != nil {
		log.Print(err)
	}
}

func headerContainer() widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(15))),
	)

	c2 := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Left:  25,
				Right: 25,
			}),
		)),
	)
	c.AddChild(c2)

	c2.AddChild(widget.NewText(widget.TextOpts.Text("hello world!", basicfont.Face7x13, color.RGBA64{0x7777, 0x0000, 0xffff, 0xffff})))
	c2.AddChild(
		widget.NewText(
			widget.TextOpts.Text("This program is a showcase of Ebiten UI widgets and layouts.", basicfont.Face7x13, color.White),
		),
	)

	return c
}

func createUI() (*ebitenui.UI, error) {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20))))

	rootContainer.AddChild(headerContainer())

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui, nil
}
