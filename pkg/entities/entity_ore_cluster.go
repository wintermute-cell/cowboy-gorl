package entities

import (
	"cowboy-gorl/pkg/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*OreClusterEntity)(nil)

// Ore type enum
type OreType int32
const (
    Coal OreType = iota
    Iron
)

//
//  OreCluster Entity
//
type OreClusterEntity struct {
    ore_type OreType
    ore_sprite rl.Texture2D
    max_capacity int32
    move_speed float32
    xpos float32
    target_inventory *InventoryEntity
    interaction_sound rl.Sound
}

func (ent *OreClusterEntity) Init() {
    // NOTE: Does this belong in Init or in the New...Entity function??
    ent.interaction_sound = rl.LoadSound("audio/foley/ore_interaction1.ogg")
}

func (ent *OreClusterEntity) Deinit() {
    rl.UnloadTexture(ent.ore_sprite)
}

func (ent *OreClusterEntity) Update() {
    var scalefactor float32 = 1.6
    ypos := render.Rs.RenderResolution.Y-118
    rl.DrawTextureEx(
        ent.ore_sprite,
        rl.NewVector2(float32(ent.xpos), ypos),
        0.0,
        scalefactor,
        rl.White)
    ent.xpos -= ent.move_speed * rl.GetFrameTime() * 100

    // Check for a click event
    if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
        mouseX := float32(rl.GetMouseX())
        mouseY := float32(rl.GetMouseY())
        if mouseX >= ent.xpos && mouseX <= ent.xpos+32*scalefactor && mouseY >= ypos && mouseY <= ypos+32*scalefactor {
            ent.HandleClick()
        }
    }
}

// NOTE: Maybe it would be better to separate this into a separate, generic ClickableEntity,
// and each OreClusterEntity has a sub entity of this type.
func (ent *OreClusterEntity) HandleClick() {
    // TODO: Implement this. The click somehow has to add the resource to the
    // players inventory.
    switch ent.ore_type {
    case Coal:
        ent.target_inventory.Coal_ore += 1
    case Iron:
        ent.target_inventory.Iron_ore += 1
    }
    // TODO: could make this ore type dependant if we had more sounds
    rl.PlaySound(ent.interaction_sound)

    // NOTE: it would be nice to also play a little animation here. we could do that by modifying the 
    // xpos and ypos over time. but that would be way easier, if animation was abstracted. and that again would
    // be way easier, if position was abstracted, or at least standardized. maybe we need another sub-structure below
    // entities? maybe "components"? (I think rerolling to ECS would be unneccesary, just a name coincidence)
}

// NOTE: this New...Entity function is a bit weird since it returns a pointer.
// not sure if this is th better way to handle this, but it needs to be a pointer if
// we want to use range to loop over a slice of structs returned by calls to this function.
// (since range generates copies of the slice entires. the alternative would be a fori loop with slice[i] indexing)
func NewOreClusterEntity(ore_type OreType, max_capacity int32, move_speed float32, target_inventory *InventoryEntity) (*OreClusterEntity) {
    new_cluster := OreClusterEntity{}
    new_cluster.ore_type = ore_type
    new_cluster.max_capacity = max_capacity
    new_cluster.move_speed = move_speed
    new_cluster.xpos = render.Rs.RenderResolution.X + 10
    new_cluster.target_inventory = target_inventory
    switch new_cluster.ore_type {
    case Coal:
        new_cluster.ore_sprite = rl.LoadTexture("sprites/resources/clusters/coal_cluster.png")
    case Iron:
        new_cluster.ore_sprite = rl.LoadTexture("sprites/resources/clusters/iron_cluster.png")
    }
    return &new_cluster
}

func (ent *OreClusterEntity) IsOnScreen() bool {
    return ent.xpos >= -64 // not comparing to 0 to have a little bit of a buffer outside of screen
}
