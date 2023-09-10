package gui

import (
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Widget interface {
	SetPosition(p rl.Vector2)
	GetPosition() rl.Vector2
	SetSize(s rl.Vector2)
	GetSize() rl.Vector2
	Bounds() rl.Rectangle
}

// ----------------
//   BASE WIDGET  |
// ----------------

type BaseWidget struct {
	position rl.Vector2
	size     rl.Vector2
}

func (w *BaseWidget) SetPosition(p rl.Vector2) {
	w.position = p
}

func (w *BaseWidget) GetPosition() rl.Vector2 {
	return w.position
}

func (w *BaseWidget) SetSize(s rl.Vector2) {
	w.size = s
}

func (w *BaseWidget) GetSize() rl.Vector2 {
	return w.size
}

func (w *BaseWidget) Bounds() rl.Rectangle {
	return rl.NewRectangle(w.position.X, w.position.Y, w.size.X, w.size.Y)
}

// ----------------
//    CONTAINER   |
// ----------------

type Container struct {
	BaseWidget
	Children []Widget
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) AddChild(child Widget) {
	c.Children = append(c.Children, child)
}

func (c *Container) RemoveChild(target Widget) {
	for i, child := range c.Children {
		if child == target {
			c.Children = util.DelFromSlice(c.Children, i, i+1)
			return
		}
	}
}

// ----------------
//      LABEL     |
// ----------------

// widget definition
type Label struct {
	BaseWidget
	text       string
	state      *LabelState
	style_info string
	// we have a state pointer inside the widget itself, since each widget must
	// have the ability to update its own state, and do these updates based on
	// it's state. (see ScrollPanel for example, the scroll position is not
	// ephemeral)
}

// state info
type LabelState struct {
	// label does not need state info yet
}

// update function
func (label *Label) update_label() {
	// maybe add a hover tooltip here, or whatever else
}

// constructor
func NewLabel(text string, position rl.Vector2, style_info string) *Label {
	l := Label{text: text, style_info: style_info}
	l.position = position
	l.state = &LabelState{}
	return &l
}

// ----------------
//     BUTTON     |
// ----------------

// widget definiton
type Button struct {
	BaseWidget
	text       string
	size       rl.Vector2
	callback   func(ButtonState)
	state      ButtonState
	style_info string
}

// state info
type ButtonState int32

const (
	ButtonStateNone ButtonState = iota
	ButtonStateHovered
	ButtonStatePressed
	ButtonStateReleased
)

// update function
func (button *Button) update_button() {
	bounds := rl.NewRectangle(button.position.X, button.position.Y, button.size.X, button.size.Y)
	button.state = ButtonStateNone
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), bounds) {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			// mouse down, button is pressed down
			button.state = ButtonStatePressed
		} else {
			// mouse is hovering, no click occured
			button.state = ButtonStateHovered
		}
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			// button released, do the button action
			button.state = ButtonStateReleased
			button.callback(button.state)
		}
	}
}

// constructor
func NewButton(text string, position, size rl.Vector2, callback func(ButtonState), style_info string) *Button {
    // NOTE: The commented code below won't work out, since data only flows
    // logic -> rendering, and the font is a part of the rendering.
	//
	//if size == rl.Vector2Zero() {
	//	// NOTE: this might have to be changed if we use custom spacings (using MeasureTextEx)
	//	text_size := rl.MeasureText(text, font.BaseSize*int32(font_scale))
	//	size = rl.NewVector2(float32(text_size), float32(font.BaseSize)*font_scale)
	//}
	new_button := &Button{text: text, size: size, callback: callback, state: ButtonStateNone, style_info: style_info}
	new_button.position = position
	return new_button
}

// ----------------
//  SCROLL PANEL  |
// ----------------

// widget definiton
type ScrollPanel struct {
	BaseWidget
	visible_bounds      rl.Rectangle
	full_bounds         rl.Rectangle
	state               *ScrollPanelState
	container           *Container
	reference_positions map[Widget]rl.Vector2 // stores the original positions of children
	style_info          string
}

// state info
type ScrollPanelState struct {
	scroll_position rl.Vector2
}

// update function
func (scroll_panel *ScrollPanel) update_scroll_panel() {
	// check if the mouse overlaps the visible bounds
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), scroll_panel.visible_bounds) {
		wheel_move := rl.GetMouseWheelMoveV()

		// Compute the new scroll position
		new_scroll_position := rl.Vector2Add(
			scroll_panel.state.scroll_position,
			rl.Vector2Scale(wheel_move, rl.GetFrameTime()*1800), // scroll speed
		)

		maxXScroll := scroll_panel.visible_bounds.Width - scroll_panel.full_bounds.Width
		maxYScroll := scroll_panel.visible_bounds.Height - scroll_panel.full_bounds.Height

		// Limit the scroll_position based on full_bounds.
		if new_scroll_position.X > 0 {
			new_scroll_position.X = 0
		}
		if new_scroll_position.Y > 0 {
			new_scroll_position.Y = 0
		}
		if new_scroll_position.X < maxXScroll {
			new_scroll_position.X = maxXScroll
		}
		if new_scroll_position.Y < maxYScroll {
			new_scroll_position.Y = maxYScroll
		}

		// Apply the new scroll position
		scroll_panel.state.scroll_position = new_scroll_position

		for _, child_widget := range scroll_panel.container.Children {
			new_pos := rl.Vector2Add(
				scroll_panel.reference_positions[child_widget],
				scroll_panel.state.scroll_position)
			child_widget.SetPosition(new_pos)
		}
	}
}

// constructor
func NewScrollPanel(visible_bounds, full_bounds rl.Rectangle, style_info string) *ScrollPanel {
	return &ScrollPanel{
		visible_bounds:      visible_bounds,
		full_bounds:         full_bounds,
		state:               &ScrollPanelState{},
		container:           NewContainer(),
		reference_positions: make(map[Widget]rl.Vector2),
		style_info:          style_info,
	}
}

func (scroll_panel *ScrollPanel) AddChild(child Widget) {
	scroll_panel.container.Children = append(scroll_panel.container.Children, child)
	scroll_panel.reference_positions[child] = child.GetPosition()
}

func (scroll_panel *ScrollPanel) RemoveChild(target Widget) {
	for i, child := range scroll_panel.container.Children {
		if child == target {
			// Remove child while preserving order
			scroll_panel.container.Children = util.DelFromSlice(scroll_panel.container.Children, i, i+1)
			return
		}
	}
}

// ----------------
//       GUI      |
// ----------------

//type Gui struct {
//    // Widget is an interface, thus no need to make this a slice of pointers.
//    // (as long as we make sure we only feed pointers into it lol)
//    widgets []Widget
//}

type Gui struct {
	container Container
}

func NewGui() *Gui {
	return &Gui{}
}

func (gui *Gui) AddWidget(widget Widget) {
	gui.container.AddChild(widget)
}

func (gui *Gui) RemoveWidget(widget Widget) {
	gui.container.RemoveChild(widget)
}

func (gui *Gui) Draw() {
	doRecursiveDraw(gui.container)
}

func doRecursiveDraw(container Container) {
	for _, widget := range container.Children {
		switch w := any(widget).(type) {
		case *Label:
			w.update_label()
			backend_label(*w)
			backend_label_finalize(*w)
		case *Button:
			w.update_button()
			backend_button(*w)
			backend_button_finalize(*w)
		case *ScrollPanel:
			w.update_scroll_panel()
			backend_scroll_panel(*w)
			doRecursiveDraw(*w.container) // draw the panels children
			backend_scroll_panel_finalize(*w)
		default:
			logging.Error("Attempted to draw GUI widget type with missing draw case: %v", w)
		}
	}
}
