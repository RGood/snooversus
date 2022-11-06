package pages

import (
	"image/color"
	"time"

	"github.com/RGood/snooverse-client/pkg/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
)

type LoadingPageConfig struct {
	Cancel func()
}

var LoadingPage = func(cfg LoadingPageConfig) Page {
	spinnerImg := ebiten.NewImage(50, 50)
	spinnerImg.Fill(color.White)

	pb := NewPageBuilder(&struct{}{}).
		SetDimensions(1024, 768).
		SetCursor(assets.NewCursor()).
		AddDrawable(
			assets.NewButton[struct{}](190, 300, 100, 150,
				assets.ButtonText[struct{}]("Loading...", color.White, 5, 15, basicfont.Face7x13),
			),
		).
		AddDrawable(assets.NewSpinner[struct{}](200, 180, spinnerImg, time.Second*2)).
		AddDrawable(
			assets.NewButton[struct{}](200, 250, 300, 350,
				assets.ButtonText[struct{}]("Cancel", color.White, 5, 15, basicfont.Face7x13),
				assets.ButtonOnHover(func(b *assets.Button[struct{}]) {
					b.SetText("Cancel", basicfont.Face7x13, 5, 15, color.RGBA64{0x8FFF, 0xFFFF, 0xFFFF, 0xFFFF})
				}),
				assets.ButtonOnClick(func(b *assets.Button[struct{}]) {
					cfg.Cancel()
				}),
			),
		)

	return pb.Build()
}
