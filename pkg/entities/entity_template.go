package entities

import ()

// This checks at compile time if the interface is implemented
var _ Entity = (*TemplateEntity)(nil)

//
//  Template Entity
//
type TemplateEntity struct {
    // Add fields here for any state that the entity should keep track of
}

func (ent *TemplateEntity) Init() {
    // Initialization logic for the entity
}

func (ent *TemplateEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *TemplateEntity) Update() {
    // Update logic for the entity
}


