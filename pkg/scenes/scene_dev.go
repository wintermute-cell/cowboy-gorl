package scenes

import (
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/util"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*DevScene)(nil)

//
//  Dev Scene
//
type DevScene struct {
    inventory entities.InventoryEntity
    background entities.BackgroundEntity
    rails entities.BackgroundEntity
    train entities.TrainEntity
    ui entities.GameplayUIEntity
    ore_clusters []*entities.OreClusterEntity

    // ore spawner
	spawn_timer       float32
	spawn_interval    float32
	next_spawn_type    entities.OreType
    // NOTE: having the OreType enum (entities.Coal here) scoped to entities is very much not ideal..

    // resource drain
    drain_timer float32
    drain_interval float32
    // NOTE: Maybe it would be nice to abstract timers into their own package or entity
}

func (scn *DevScene) Init() {
    scn.inventory = entities.InventoryEntity{}
    scn.inventory.Init()
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

    // NOTE: wouldn't it be better to just be able to use the Init function to
    // set initial values? or should startup logic be separate from setting
    // initial values?
    scn.ui = entities.NewGameplayUiEntity(&scn.inventory)
    scn.ui.Init()


    scn.spawn_timer = 0
	scn.spawn_interval = scn.generateSpawnInterval() // Initial spawn interval
	scn.next_spawn_type = scn.generateRandomOreType()

    scn.drain_interval = 1.0 // burn coal every second

    logging.Info("DevScene initialized.")
}

func (scn *DevScene) Deinit() {
    scn.inventory.Deinit()
    scn.background.Deinit()
    scn.rails.Deinit()
    scn.train.Deinit()
    scn.ui.Deinit()
    for _, cluster := range scn.ore_clusters {
        cluster.Deinit()
    }
    logging.Info("DevScene de-initialized.")
}

func (scn *DevScene) DrawGUI() {
    scn.ui.Update()
}

func (scn *DevScene) Draw() {
	scn.background.Update()
	scn.rails.Update()
	scn.train.Update()

    // handle regular ore cluster spawning
	scn.spawn_timer += 1.0 * rl.GetFrameTime()
	if scn.spawn_timer >= scn.spawn_interval {
		scn.spawnOreCluster()
		scn.spawn_timer = 0.0
		scn.spawn_interval = scn.generateSpawnInterval()
		scn.next_spawn_type = scn.generateRandomOreType()
	}

    // draw the current ore clusters and delete them if they're offscreen
	for i, cluster := range scn.ore_clusters {
		cluster.Update()
        if !cluster.IsOnScreen() {
            cluster.Deinit()
            scn.ore_clusters = util.DeleteFromSlice(scn.ore_clusters, i, i+1)
        }
	}

    // handle regular resource drain. At the moment, this is just the engine burning coal
    scn.drain_timer += 1.0 * rl.GetFrameTime()
    if scn.drain_timer >= scn.drain_interval {
        scn.drain_timer = 0.0
        scn.inventory.Coal_ore -= 1
    }
}


// Helper function to generate a random spawn interval
func (scn *DevScene) generateSpawnInterval() float32 {
	return 5.0 + rand.Float32()*10.0 // Random interval between 5 to 15 seconds
}

// Helper function to generate a random ore type
func (scn *DevScene) generateRandomOreType() entities.OreType {
	if rand.Intn(2) == 0 {
		return entities.Coal
	}
	return entities.Iron
}

// Function to spawn a new ore cluster
func (scn *DevScene) spawnOreCluster() {
	speed := 0.6 + (rand.Float32()*0.4 - 0.2) // Random speed between 0.4 to 0.8
	scn.ore_clusters = append(scn.ore_clusters, entities.NewOreClusterEntity(scn.next_spawn_type, 100, speed, &scn.inventory))
    scn.ore_clusters[len(scn.ore_clusters)-1].Init() // NOTE: Does this need to be? Init functions and New...Entity functions really need an overhaul
}
