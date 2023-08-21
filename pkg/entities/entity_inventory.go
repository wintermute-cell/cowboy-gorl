package entities

import ()

// This checks at compile time if the interface is implemented
var _ Entity = (*InventoryEntity)(nil)

//
//  Inventory Entity
//
type InventoryEntity struct {
    Coal_ore int32
    Iron_ore int32
}

func (ent *InventoryEntity) Init() {
    ent.Coal_ore = 100
}

func (ent *InventoryEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *InventoryEntity) Update() {
    // Update logic for the entity
}


