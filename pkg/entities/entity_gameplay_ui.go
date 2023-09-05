package entities

import (
	"cowboy-gorl/pkg/render"
	"fmt"

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
    inventory *InventoryEntity
}

func NewGameplayUiEntity(inventory *InventoryEntity) (GameplayUIEntity) {
    new_ent := GameplayUIEntity{}
    new_ent.inventory = inventory
    return new_ent
}

func (ent *GameplayUIEntity) Init() {
    // Initialization logic for the entity
}

func (ent *GameplayUIEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *GameplayUIEntity) Update() {
    // Update logic for the entity
    origin := rl.NewVector2(0.0, render.Rs.RenderResolution.Y-66)
    rl.DrawRectangle(
        int32(origin.X), int32(origin.Y),
        int32(render.Rs.RenderResolution.X),
        int32(render.Rs.RenderResolution.Y-origin.Y), rl.Brown)

    offset := float32(0)

    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Pops");
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 18)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), fmt.Sprintf("%v", ent.inventory.Pops))
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Coal");
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 30)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), fmt.Sprintf("%v", ent.inventory.Coal_ore))
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Iron");
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 30)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), fmt.Sprintf("%v", ent.inventory.Iron_ore))
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Dynmt");
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 30)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), fmt.Sprintf("%v", ent.inventory.Dynamite))
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

    offset += 70
    rg.Label(rl.NewRectangle(origin.X + 8 + offset, origin.Y + 6, 60, 24), "Food");
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 30)
    rg.Label(rl.NewRectangle(origin.X + 14 + offset, origin.Y + 28, 60, 24), fmt.Sprintf("%v", ent.inventory.Food))
//    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
}
