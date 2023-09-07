package scenes

import (
	"cowboy-gorl/pkg/entities"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*ScenePickerScene)(nil)

//
//  ScenePicker Scene
//
type ScenePickerScene struct {
    // Required fields
    entity_manager *entities.EntityManager

    // Custom Fields
}

func (scn *ScenePickerScene) Init() {
    // Required initialization
    scn.entity_manager = entities.NewEntityManager()

    // Initialization logic for the scene
    // ...
}

func (scn *ScenePickerScene) Deinit() {
    // De-initialization logic for the scene
}

func (scn *ScenePickerScene) DrawGUI() {
    if rg.Button(rl.NewRectangle(100, 100, 300, 32), "Animation Dev Scene") {
        Sm.DisableScene("scene_picker")
        Sm.EnableScene("anim_dev")
    }
}

func (scn *ScenePickerScene) Draw() {
}
