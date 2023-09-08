package scenes

import (
	"cowboy-gorl/pkg/entities"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*GuiDevScene)(nil)

//
//  GuiDev Scene
//
type GuiDevScene struct {
    // Required fields
    entity_manager *entities.EntityManager

    // Custom Fields
    // Add fields here for any state that the scene should keep track of
    // ...
}

func (scn *GuiDevScene) Init() {
    // Required initialization
    scn.entity_manager = entities.NewEntityManager()

    // Initialization logic for the scene
    // ...
}

func (scn *GuiDevScene) Deinit() {
    // De-initialization logic for the scene
}

func (scn *GuiDevScene) DrawGUI() {
    // Draw the GUI for the scene
}

func (scn *GuiDevScene) Draw() {
    // Draw the scene
}
