package scenes

import (
	"cowboy-gorl/pkg/audio"
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/gui"
	"cowboy-gorl/pkg/settings"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*DevMenuScene)(nil)

// DevMenu Scene
type DevMenuScene struct {
	// Required fields
	entity_manager *entities.EntityManager

    g *gui.Gui
}

func (scn *DevMenuScene) Init() {
	// Required initialization
	scn.entity_manager = entities.NewEntityManager()
    audio.SetCurrentPlaylist("main-menu", true)

    scn.g = gui.NewGui()
    btn := gui.NewButton("Audio Dev Scene", rl.NewVector2(408, 4), rl.NewVector2(128, 32), func(s gui.ButtonState) {
		Sm.DisableAllScenesExcept([]string{"dev_menu"})
		Sm.EnableScene("audio_dev")
    }, "")
    scn.g.AddWidget(btn)

	// Initialization logic for the scene
	// ...
}

func (scn *DevMenuScene) Deinit() {
	// De-initialization logic for the scene
}

func (scn *DevMenuScene) DrawGUI() {
    scn.g.Draw()
	original_text_size := rg.GetStyle(rg.DEFAULT, rg.TEXT_SIZE)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 16)

	if rg.Button(rl.NewRectangle(4, 4, 32, 32), "CRT") {
		settings.CurrentSettings().EnableCrtEffect = !settings.CurrentSettings().EnableCrtEffect
	}

	if rg.Button(rl.NewRectangle(40, 4, 180, 32), "Animation Dev Scene") {
		Sm.DisableAllScenesExcept([]string{"dev_menu"})
		Sm.EnableScene("anim_dev")
	}

	if rg.Button(rl.NewRectangle(224, 4, 180, 32), "GUI Dev Scene") {
		Sm.DisableAllScenesExcept([]string{"dev_menu"})
		Sm.EnableScene("gui_dev")
	}

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, original_text_size)
}

func (scn *DevMenuScene) Draw() {
	// Draw the scene
}
