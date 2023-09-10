package scenes

import (
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*GuiDevScene)(nil)

// GuiDev Scene
type GuiDevScene struct {
	// Required fields
	entity_manager *entities.EntityManager

	// Custom Fields
	g *gui.Gui
}

func (scn *GuiDevScene) Init() {
	// Required initialization
	scn.entity_manager = entities.NewEntityManager()

	scn.g = gui.NewGui()

	// gui elements
    label := gui.NewLabel("retained mode widget", rl.NewVector2(16, 64), "font:alagard|debug:true")
	btn := gui.NewButton("retained button", rl.NewVector2(16, 96), rl.NewVector2(15*6, 16), func(_ gui.ButtonState) {}, "")

    scroll_panel := gui.NewScrollPanel(
        rl.NewRectangle(10, 48, 620, 400),
        rl.NewRectangle(10, 48, 620, 2000),
        "debug:false|background:200,200,200,255")
	scn.g.AddWidget(scroll_panel)
	scroll_panel.AddChild(label)
	scroll_panel.AddChild(btn)
}

func (scn *GuiDevScene) Deinit() {
	// De-initialization logic for the scene
}

func (scn *GuiDevScene) DrawGUI() {

	scn.g.Draw()

	// pixel ruler
	offs := int32(7)
	rl.DrawLine(0, 480, 0, 300, rl.Red)
	rl.DrawLine(1, 480, 1, 300, rl.Green)
	rl.DrawLine(offs*1, 480, offs*1, 300, rl.Red)
	rl.DrawLine(offs*2, 480, offs*2, 300, rl.Red)
	rl.DrawLine(offs*3, 480, offs*3, 300, rl.Red)
	rl.DrawLine(offs*4, 480, offs*4, 300, rl.Red)
	rl.DrawLine(offs*5, 480, offs*5, 300, rl.Red)
}

func (scn *GuiDevScene) Draw() {
	// Draw the scene
}
