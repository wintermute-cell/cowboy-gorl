package entities

import (
	"cowboy-gorl/pkg/render"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*GameplayUIEntity)(nil)

//
//  GameplayUI Entity
//
type GameplayUIEntity struct {
    // Add fields here for any state that the entity should keep track of
}

func (ent *GameplayUIEntity) Init() {
    // Initialization logic for the entity
}

func (ent *GameplayUIEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *GameplayUIEntity) Update() {
    // Update logic for the entity
    origin := rl.NewVector2(0.0, render.Rs.RenderResolution.Y-76)
    rl.DrawRectangle(
        int32(origin.X), int32(origin.Y),
        int32(render.Rs.RenderResolution.X),
        int32(render.Rs.RenderResolution.Y-origin.Y), rl.Brown)

    offset := float32(0)

    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Pops");
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), "44")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 14)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Repair");
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), "12")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 14)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Collect");
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), "14")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 14)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Craft");
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), "6")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 14)
}
