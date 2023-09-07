// EntityManager provides a manager for game entities, automating the calling
// of their Init(), Deinit() and Update() functions, 
// A EntityManager also features enabling/disabling, and ordering of entities
// for update operations.
// 
// Usage:
//    - Create a new EntityManager with `NewEntityManager`.
//    - Register entities using `RegisterEntity(name, entity)`.
//    - Control entity state with `EnableEntity` and `DisableEntity`.
//    - Modify update order using `MoveEntityToFront`, `MoveEntityToBack`, and `MoveEntityBefore`.
//    - In the game loop, use `UpdateEntities` update entities in their specified order.

package entities

import "cowboy-gorl/pkg/logging"

type EntityManager struct {
	entities        map[string]Entity
	enable_entities map[string]bool
	entity_order   []string // slice to maintain order, since map is unordered
}

// Create a new EntityManager. A EntityManager will automatically take care of
// your Entities (calling their Init(), Deinit(), Update() functions).
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities:        make(map[string]Entity),
		enable_entities: make(map[string]bool),
		entity_order:   make([]string, 0),
	}
}

// Register an entity with the EntityManager for automatic control
func (em *EntityManager) RegisterEntity(name string, entity Entity, enable_immediately bool) {
	if _, exists := em.entities[name]; exists {
		logging.Fatal("An entity with name \"%v\" is already registered.", name)
	}
	em.entities[name] = entity
	em.entity_order = append(em.entity_order, name) // Add to the end by default


    // immediately enable the entity
    if enable_immediately {
        em.EnableEntity(name)
        logging.Info("Registered Entity with name \"%v\" and immediately enabled.", name)
    } else {
        logging.Info("Registered Entity with name \"%v\" without enabling.", name)
    }
}

// MoveEntityToFront moves the entity to the front of the update order
func (em *EntityManager) MoveEntityToFront(name string) {
	em.reorderEntity(name, 0)
}

// MoveEntityToBack moves the entity to the end of the update order
func (em *EntityManager) MoveEntityToBack(name string) {
	em.reorderEntity(name, len(em.entity_order)-1)
}

// MoveEntityBefore moves the entity right before another entity in the update order
func (em *EntityManager) MoveEntityBefore(entityName, beforeEntityName string) {
	index, found := em.getEntityOrderIndex(beforeEntityName)
	if found {
		em.reorderEntity(entityName, index)
	}
}

func (em *EntityManager) reorderEntity(name string, index int) {
	current_idx, found := em.getEntityOrderIndex(name)
	if !found {
		return
	}
	em.entity_order = append(em.entity_order[:current_idx], em.entity_order[current_idx+1:]...)
	em.entity_order = append(em.entity_order[:index], append([]string{name}, em.entity_order[index:]...)...)
}

func (em *EntityManager) getEntityOrderIndex(name string) (int, bool) {
	for i, entity_name := range em.entity_order {
		if entity_name == name {
			return i, true
		}
	}
	return -1, false
}

// Enable the Entity. The Entities Init() function will be called.
func (em *EntityManager) EnableEntity(name string) {
	entity, exists := em.entities[name]
	if !exists {
		logging.Fatal("Entity with name %v not found.", name)
	}

	// Initialize the entity if it's not already enabled
	if !em.enable_entities[name] {
		entity.Init()
		em.enable_entities[name] = true
	}
}

// Disable the Entity. The Entities Deinit() function will be called.
func (em *EntityManager) DisableEntity(name string) {
	entity, exists := em.entities[name]
	if !exists {
		logging.Fatal("Entity with name %v not found.", name)
	}

	// De-initialize the entity if it's currently enabled
	if em.enable_entities[name] {
		entity.Deinit()
		em.enable_entities[name] = false
	}
}

// Disable all Entities that are currently enabled.
func (em *EntityManager) DisableAllEntities() {
    for _, name := range em.entity_order {
		if em.enable_entities[name] {
			em.entities[name].Deinit()
            em.enable_entities[name] = false
		}
    }
}

// Call the Update() functions of all the registered Entities in their defined order.
func (em *EntityManager) UpdateEntities() {
    for _, name := range em.entity_order {
		if em.enable_entities[name] {
			em.entities[name].Update()
		}
    }
}
