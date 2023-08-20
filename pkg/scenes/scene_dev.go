package scenes

import (
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/logging"
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
            "backgrounds/super-parallax-mountains/sky.png",
            "backgrounds/super-parallax-mountains/far-clouds.png",
            "backgrounds/super-parallax-mountains/near-clouds.png",
            "backgrounds/super-parallax-mountains/far-mountains.png",
            "backgrounds/super-parallax-mountains/mountains.png",
            "backgrounds/super-parallax-mountains/trees.png",
        },
        []float32{
            0.0,
            0.1,
            0.2,
            0.5,
            0.8,
            1.0,
        })
    scn.train = entities.TrainEntity{}
    scn.train.Init()
    logging.Info("DevScene initialized.")
}

func (scn *DevScene) Deinit() {
    scn.background.Deinit()
    logging.Info("DevScene de-initialized.")
}

func (scn *DevScene) DrawGUI() {
    // Draw the GUI for the scene
}

func (scn *DevScene) Draw() {
    scn.background.Update()
    scn.train.Update()
}
