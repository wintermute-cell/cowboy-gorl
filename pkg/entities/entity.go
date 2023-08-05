package entities

// Entity is an interface that every entity in the game should implement
type Entity interface {
    Init()
    Update()
    Deinit()
}

