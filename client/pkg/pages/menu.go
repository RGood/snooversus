package pages

import (
	"image/color"

	"github.com/RGood/snooverse-client/pkg/assets"
	"golang.org/x/image/font/basicfont"
)

type MenuState struct {
	Foo string
}

type MenuPageConfig struct {
	StartButton func()
}

var MenuPage = func(cfg MenuPageConfig) *CustomPage[MenuState] {

	pb := NewPageBuilder(&MenuState{
		Foo: "bar",
	}).
		SetDimensions(1024, 768).
		SetCursor(assets.NewCursor()).
		SetKeyboard(assets.NewKeyboard()).
		AddDrawable(
			assets.NewButton[MenuState](20, 100, 20, 50,
				//assets.ButtonFillColor[MenuState](color.White),
				assets.ButtonText[MenuState]("Start", color.White, 5, 15, basicfont.Face7x13),
				assets.ButtonOnHover(func(b *assets.Button[MenuState]) {
					//b.SetFill(color.Gray16{0xEFFF})
					b.SetText("Start", basicfont.Face7x13, 5, 15, color.RGBA64{0x8FFF, 0xFFFF, 0xFFFF, 0xFFFF})
				}),
				assets.ButtonOnClick(func(b *assets.Button[MenuState]) {
					cfg.StartButton()
				}),
			),
		)

	return pb.Build()
}
