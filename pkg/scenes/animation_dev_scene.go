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
	background entities.BackgroundEntity
	anim_mover entities.AnimationTestMoverEntity
	anim_text  entities.AnimationTestTextEntity
	anim_ore   entities.AnimationTestOreEntity
}

func (scn *AnimationDevScene) Init() {
	scn.background = entities.BackgroundEntity{}
	scn.background.SetLayers(
		[]string{
			"backgrounds/parallax-mountain/parallax-mountain-bg.png",
		},
		[]float32{
			0.0,
		})

	screen_middle := rl.NewVector2(
		float32(int32(render.Rs.RenderResolution.X/2.0))-32,
		float32(int32(render.Rs.RenderResolution.Y/2.0))-32,
	)
	scn.anim_ore = entities.NewAnimationTestOreEntity(entities.OreType_Iron, screen_middle)
	scn.anim_ore.Init()

	scn.anim_mover = entities.AnimationTestMoverEntity{}
	scn.anim_mover.Init()

	scn.anim_text = entities.AnimationTestTextEntity{}
	scn.anim_text.Init()

	logging.Info("AnimationTestScene initialized.")
}

func (scn *AnimationDevScene) Deinit() {
	scn.background.Deinit()
	scn.anim_ore.Deinit()
	scn.anim_text.Deinit()
	scn.anim_mover.Deinit()
	logging.Info("AnimationTestScene de-initialized.")
}

func (scn *AnimationDevScene) DrawGUI() {
	// Draw the GUI for the scene
}

func (scn *AnimationDevScene) Draw() {
	scn.background.Update()
	scn.anim_mover.Update()
	scn.anim_text.Update()
	scn.anim_ore.Update()
}
