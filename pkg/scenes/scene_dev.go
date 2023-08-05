package scenes

import (
    "cowboy-gorl/pkg/entities"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*DevScene)(nil)

//
//  Dev Scene
//
type DevScene struct {
    background entities.BackgroundEntity
    train entities.TrainEntity
}

func (scn *DevScene) Init() {
    scn.background = entities.BackgroundEntity{}
    scn.background.SetLayers(
        []string{
            "backgrounds/parallax-mountain/parallax-mountain-bg.png",
            "backgrounds/parallax-mountain/parallax-mountain-montain-far.png",
            "backgrounds/parallax-mountain/parallax-mountain-mountains.png",
            "backgrounds/parallax-mountain/parallax-mountain-trees.png",
            "backgrounds/parallax-mountain/parallax-mountain-foreground-trees.png",
        },
        []float32{
            0.0,
            0.1,
            0.2,
            0.5,
            0.8,
        })
    scn.train = entities.TrainEntity{}
    scn.train.Init()
}

func (scn *DevScene) Deinit() {
    scn.background.Deinit()
}

func (scn *DevScene) DrawGUI() {
    // Draw the GUI for the scene
}

func (scn *DevScene) Draw() {
    scn.background.Update()
    scn.train.Update()
}
