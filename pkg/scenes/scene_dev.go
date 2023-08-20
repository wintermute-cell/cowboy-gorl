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
    rails entities.BackgroundEntity
    train entities.TrainEntity
    ui entities.GameplayUIEntity
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
            0.1 * 2.5,
            0.2 * 2.5,
            0.5 * 2.5,
            0.8 * 2.5,
            1.0 * 2.5,
        })

    scn.rails = entities.BackgroundEntity{}
    scn.background.SetLayers(
        []string{
            "sprites/rails.png",
        },
        []float32{
            10.0,
        })

    scn.train = entities.TrainEntity{}
    scn.train.Init()

    scn.ui.Init()

    logging.Info("DevScene initialized.")
}

func (scn *DevScene) Deinit() {
    scn.background.Deinit()
    scn.rails.Deinit()
    scn.train.Deinit()
    scn.ui.Deinit()
    logging.Info("DevScene de-initialized.")
}

func (scn *DevScene) DrawGUI() {
    scn.ui.Update()
}

func (scn *DevScene) Draw() {
    scn.background.Update()
    scn.rails.Update()
    scn.train.Update()
}
