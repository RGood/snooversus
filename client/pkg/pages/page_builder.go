package pages

import (
	"github.com/RGood/snooverse-client/pkg/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type PageBuilder[T any] struct {
	shouldReset  bool
	initialState *T
	onUpdate     func(*T)
	drawables    []assets.Drawable[T]
	cursor       *assets.Cursor
	keyboard     *assets.Keyboard
}

func NewPageBuilder[T any]() *PageBuilder[T] {
	return &PageBuilder[T]{
		shouldReset:  true,
		initialState: nil,
		onUpdate:     nil,
		drawables:    []assets.Drawable[T]{},
		cursor:       nil,
		keyboard:     nil,
	}
}

func (pb *PageBuilder[T]) AddDrawable(e assets.Drawable[T]) {
	pb.drawables = append(pb.drawables, e)
}

func (pb *PageBuilder[T]) SetInitialState(is *T) {
	pb.initialState = is
}

func (pb *PageBuilder[T]) SetOnUpdate(ou func(*T)) {
	pb.onUpdate = ou
}

func (pb *PageBuilder[T]) SetCursor(cursor *assets.Cursor) {
	pb.cursor = cursor
}

func (pb *PageBuilder[T]) SetKeyboard(kb *assets.Keyboard) {
	pb.keyboard = kb
}

func (pb *PageBuilder[T]) Build() *CustomPage[T] {
	return &CustomPage[T]{
		shouldReset: pb.shouldReset,
		state:       pb.initialState,
		onUpdate:    pb.onUpdate,
		drawables:   pb.drawables,
		cursor:      pb.cursor,
		keyboard:    pb.keyboard,
	}
}

type CustomPage[T any] struct {
	UnimplementedPage
	shouldReset bool
	state       *T
	onUpdate    func(*T)
	drawables   []assets.Drawable[T]
	cursor      *assets.Cursor
	keyboard    *assets.Keyboard
}

func (customPage *CustomPage[T]) Draw(img *ebiten.Image) {
	if customPage.shouldReset {
		img.Clear()
	}
	for _, e := range customPage.drawables {
		e.Draw(customPage.state, img)
	}
}

func (customPage *CustomPage[T]) Update() error {
	if customPage.cursor != nil {
		customPage.cursor.Update()
	}

	if customPage.keyboard != nil {
		customPage.keyboard.Update()
	}

	if customPage.onUpdate != nil {
		customPage.onUpdate(customPage.state)
	}

	return nil
}

var _ Page = (*CustomPage[any])(nil)
