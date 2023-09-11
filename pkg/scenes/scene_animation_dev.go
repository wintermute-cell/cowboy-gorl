package scenes

import (
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*AnimationDevScene)(nil)

// Animation Test Scene
type AnimationDevScene struct {
	entity_manager *entities.EntityManager
}

func (scn *AnimationDevScene) Init() {
	scn.entity_manager = entities.NewEntityManager()

	// Background
	bg_ent := entities.BackgroundEntity{}
	bg_ent.SetLayers(
		[]string{
			"backgrounds/parallax-mountain/parallax-mountain-bg.png",
		},
		[]float32{
			0.0,
		})

	scn.entity_manager.RegisterEntity("background", &bg_ent, true, []string{})

	// Animated objects
	// Ore Cluster
	screen_middle := rl.NewVector2(
		float32(int32(render.Rs.RenderResolution.X/2.0))-32,
		float32(int32(render.Rs.RenderResolution.Y/2.0))-32,
	)
	anim_ore := entities.NewAnimationTestOreEntity(entities.OreType_Iron, screen_middle)
	scn.entity_manager.RegisterEntity(
		"animated-ore",
		&anim_ore,
		true,
		[]string{},
	)

	// Animated Moving Circle
	scn.entity_manager.RegisterEntity(
		"animated-mover",
		&entities.AnimationTestMoverEntity{},
		true,
		[]string{},
	)

	// Animated Text
	scn.entity_manager.RegisterEntity(
		"animated-text",
		&entities.AnimationTestTextEntity{},
		true,
		[]string{},
	)

	logging.Info("AnimationTestScene initialized.")
}

func (scn *AnimationDevScene) Deinit() {
	scn.entity_manager.DisableAllEntities()
	logging.Info("AnimationTestScene de-initialized.")
}

func (scn *AnimationDevScene) DrawGUI() {
	// Draw the GUI for the scene
}

func (scn *AnimationDevScene) Draw() {
	scn.entity_manager.UpdateEntities()
}
