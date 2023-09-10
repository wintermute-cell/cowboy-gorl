package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// GUI BACKEND STATE
type GuiBackendState struct {
	fonts map[string]rl.Font
}

var Gbs GuiBackendState

// This function is automatically called on package import.
func InitBackend() {
	Gbs = GuiBackendState{}
	Gbs.fonts = make(map[string]rl.Font)
	Gbs.fonts["default"] = rl.GetFontDefault()
    Gbs.fonts["alagard"] = rl.LoadFont("fonts/alagard.png")
}

// BACKEND FUNCTIONS

func backend_label(label Label) {
    style := parseStyleDef(label.style_info)

    color := rl.Black
    if c, ok := style["color"]; ok && c != nil {
        color = c.(rl.Color)
    }

    font := Gbs.fonts["default"]
    if f, ok := style["font"]; ok && f != nil {
        font = Gbs.fonts[f.(string)]
    }


	rl.DrawTextEx(
		font,
		label.text,
		label.position,
		float32(font.BaseSize),
		float32(font.BaseSize/10),
		color,
	)
}

func backend_label_finalize(label Label) {
	// nothing to do here
}

func backend_scroll_panel(scroll_panel ScrollPanel) {
    style := parseStyleDef(scroll_panel.style_info)

    draw_debug := false
    if v, ok := style["debug"]; ok && v != nil {
        draw_debug = v.(bool)
    }

    bg_color := rl.NewColor(0,0,0,0)
    if v, ok := style["background"]; ok && v != nil {
        bg_color = v.(rl.Color)
    }

    if draw_debug {
        // this represents the full bounds shifted by the scroll position
        // (useful for visualizing the scroll concept)
        virt_fbounds := scroll_panel.full_bounds
        virt_fbounds.X += scroll_panel.state.scroll_position.X
        virt_fbounds.Y += scroll_panel.state.scroll_position.Y

        rl.DrawRectangleRec(virt_fbounds, rl.Blue)
    }

    rl.DrawRectangleRec(scroll_panel.visible_bounds, bg_color)

	rl.BeginScissorMode(
		int32(scroll_panel.visible_bounds.X),
		int32(scroll_panel.visible_bounds.Y),
		int32(scroll_panel.visible_bounds.Width),
		int32(scroll_panel.visible_bounds.Height),
	)
}

func backend_scroll_panel_finalize(scroll_panel ScrollPanel) {
	rl.EndScissorMode()
}

func backend_button(button Button) {
	// determine appropriate colors based on current interaction state
	btn_color := rl.Blue
	switch button.state {
	case ButtonStateHovered:
		btn_color = rl.SkyBlue
	case ButtonStatePressed:
		btn_color = rl.DarkBlue
	}

	bounds := rl.NewRectangle(button.position.X, button.position.Y, button.size.X, button.size.Y)
	rl.DrawRectangleRec(bounds, btn_color)
	rl.DrawText(button.text, int32(button.position.X), int32(button.position.Y), Gbs.fonts["default"].BaseSize, rl.White)
}

func backend_button_finalize(button Button) {
	// nothing to do here
}
