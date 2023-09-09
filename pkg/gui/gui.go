package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GuiState struct {
	fonts map[string]rl.Font
}

var Gs GuiState

func Init() {
	Gs.fonts = make(map[string]rl.Font)
	Gs.fonts["default"] = rl.GetFontDefault()
}

// ----------------
//  CONFIGURATION |
// ----------------

// Set the default font to the font specified by the path argument.
func SetDefaultFont(path string, filter_mode rl.TextureFilterMode) {
    AddFont("default", path, filter_mode)
}

// Add a font that will be addressable by "name".
func AddFont(name string, path string, filter_mode rl.TextureFilterMode) {
	Gs.fonts[name] = rl.LoadFont(path)
	rl.SetTextureFilter(Gs.fonts[name].Texture, filter_mode)
}

// ----------------
//      TEXT      |
// ----------------

// Draw text with extended parameters.
func TextEx(text string, position rl.Vector2, scale float32, font_name string, color rl.Color) {
	// if font name empty or does not exist, use default font
	font, ok := Gs.fonts[font_name]
	if !ok {
		font = Gs.fonts["default"]
	}

	rl.DrawTextEx(font, text, position, float32(font.BaseSize)*scale, float32(font.BaseSize/10), color)
}

// Draw text at the specified position using the default font.
func Text(text string, position rl.Vector2, color rl.Color) {
    TextEx(text, position, 1.0, "default", color)
}

// ----------------
//     BUTTONS    |
// ----------------

type ButtonState int32
const (
    ButtonStateNone ButtonState = iota
    ButtonStateHovered
    ButtonStatePressed
    ButtonStateReleased
)

// Draw a button with extended parameters. This exposes all the available
// button settings and does therefore not respect the current GUI style.
func ButtonEx(
    text string,
    position rl.Vector2,
    size rl.Vector2,
    font_name string,
    font_scale float32,
    normal_color_button, hover_color_button, pressed_color_button rl.Color,
    normal_color_text, hover_color_text, pressed_color_text rl.Color,
) ButtonState {
	// if font name empty or does not exist, use default font
	font, ok := Gs.fonts[font_name]
	if !ok {
		font = Gs.fonts["default"]
	}

    rg.SetFont(font)
    if size == rl.Vector2Zero() {
        // NOTE: this might have to be changed if we use custom spacings
        text_size := rl.MeasureText(text, font.BaseSize*int32(font_scale))
        size = rl.NewVector2(float32(text_size), float32(font.BaseSize)*font_scale)
    }
    bounds := rl.NewRectangle(position.X, position.Y, size.X, size.Y)

    // check for click event
    return_state := ButtonStateNone
    if rl.CheckCollisionPointRec(rl.GetMousePosition(), bounds) {
        if rl.IsMouseButtonDown(rl.MouseLeftButton) {
            // mouse down, button is pressed down
            return_state = ButtonStatePressed
        } else {
            // mouse is hovering, no click occured
            return_state = ButtonStateHovered
        }
        if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
            // button released, do the button action
            return_state = ButtonStateReleased
        }
    }

    // draw the button
    btn_color := normal_color_button
    txt_color := normal_color_text
    switch return_state {
        case ButtonStateHovered:
        btn_color = hover_color_button
        txt_color = hover_color_text
        case ButtonStatePressed:
        btn_color = pressed_color_button
        txt_color = pressed_color_text
    }

    rl.DrawRectangleRec(bounds, btn_color)
    TextEx(text, position, font_scale, font_name, txt_color)

    return return_state
}
