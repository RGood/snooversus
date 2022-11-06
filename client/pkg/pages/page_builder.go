package pages

import (
	"github.com/RGood/snooverse-client/pkg/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type PageBuilder[T any] struct {
	shouldReset   bool
	initialState  *T
	width, height int
	onUpdate      func(*T)
	drawables     []assets.Drawable[T]
	updateables   []assets.Updateable
	cursor        *assets.Cursor
	keyboard      *assets.Keyboard
}

func NewPageBuilder[T any](state *T) *PageBuilder[T] {
	return &PageBuilder[T]{
		shouldReset:  true,
		initialState: state,
		onUpdate:     nil,
		drawables:    []assets.Drawable[T]{},
		updateables:  []assets.Updateable{},
		cursor:       nil,
		keyboard:     nil,
	}
}

func (pb *PageBuilder[T]) AddDrawable(e assets.Drawable[T]) *PageBuilder[T] {
	pb.drawables = append(pb.drawables, e)
	if clickable, ok := e.(assets.Clickable); ok && pb.cursor != nil {
		pb.cursor.RegisterClickable(clickable)
	}

	if updateable, ok := e.(assets.Updateable); ok {
		pb.updateables = append(pb.updateables, updateable)
	}
	return pb
}

func (pb *PageBuilder[T]) SetInitialState(is *T) *PageBuilder[T] {
	pb.initialState = is
	return pb
}

func (pb *PageBuilder[T]) SetOnUpdate(ou func(*T)) *PageBuilder[T] {
	pb.onUpdate = ou
	return pb
}

func (pb *PageBuilder[T]) SetCursor(cursor *assets.Cursor) *PageBuilder[T] {
	pb.cursor = cursor
	return pb
}

func (pb *PageBuilder[T]) SetKeyboard(kb *assets.Keyboard) *PageBuilder[T] {
	pb.keyboard = kb
	return pb
}

func (pb *PageBuilder[T]) SetDimensions(width, height int) *PageBuilder[T] {
	pb.width = width
	pb.height = height
	return pb
}

func (pb *PageBuilder[T]) Build() *CustomPage[T] {
	cachedImg := ebiten.NewImage(pb.width, pb.height)

	return &CustomPage[T]{
		cachedImg:   cachedImg,
		width:       float64(pb.width),
		height:      float64(pb.height),
		shouldReset: pb.shouldReset,
		state:       pb.initialState,
		onUpdate:    pb.onUpdate,
		drawables:   pb.drawables,
		updateables: pb.updateables,
		cursor:      pb.cursor,
		keyboard:    pb.keyboard,
	}
}

type CustomPage[T any] struct {
	UnimplementedPage
	width, height float64
	cachedImg     *ebiten.Image
	shouldReset   bool
	state         *T
	onUpdate      func(*T)
	drawables     []assets.Drawable[T]
	updateables   []assets.Updateable
	cursor        *assets.Cursor
	keyboard      *assets.Keyboard
}

func (customPage *CustomPage[T]) Draw(img *ebiten.Image) {
	if customPage.shouldReset {
		customPage.cachedImg.Clear()
		img.Clear()
	}

	for _, e := range customPage.drawables {
		e.Draw(customPage.state, customPage.cachedImg)
	}

	img.DrawImage(customPage.cachedImg, nil)
}

func (customPage *CustomPage[T]) Update() error {
	if customPage.cursor != nil {
		customPage.cursor.Update()
	}

	if customPage.keyboard != nil {
		customPage.keyboard.Update()
	}

	for _, updateable := range customPage.updateables {
		updateable.Update()
	}

	if customPage.onUpdate != nil {
		customPage.onUpdate(customPage.state)
	}

	return nil
}

func (customPage *CustomPage[T]) Dimensions() (float64, float64) {
	return customPage.width, customPage.height
}

var _ Page = (*CustomPage[any])(nil)
