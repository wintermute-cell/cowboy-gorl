package entities

import (
	"cowboy-gorl/pkg/animation"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*AnimationTestMoverEntity)(nil)

// AnimationTestMover Entity
type AnimationTestMoverEntity struct {
	position       rl.Vector2
	move_animation *animation.Animation[float32]
}

func (ent *AnimationTestMoverEntity) Init() {
	// The MoverEntity will move in a (rough) circle

	ent.move_animation = animation.CreateAnimation[float32](4.0)
	ent.position = rl.NewVector2(0.0, 0.0)

	// center and radius
	centerX := float32(64.0)
	centerY := float32(64.0)
	radius := float32(32.0)

	// keyframes for X coordinate
	ent.move_animation.AddKeyframe(&ent.position.X, 0.0, centerX+radius)
	ent.move_animation.AddKeyframe(&ent.position.X, 0.5, centerX+radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.X, 1.0, centerX)
	ent.move_animation.AddKeyframe(&ent.position.X, 1.5, centerX-radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.X, 2.0, centerX-radius)
	ent.move_animation.AddKeyframe(&ent.position.X, 2.5, centerX-radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.X, 3.0, centerX)
	ent.move_animation.AddKeyframe(&ent.position.X, 3.5, centerX+radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.X, 4.0, centerX+radius)

	// keyframes for Y coordinate
	ent.move_animation.AddKeyframe(&ent.position.Y, 0.0, centerY)
	ent.move_animation.AddKeyframe(&ent.position.Y, 0.5, centerY+radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.Y, 1.0, centerY+radius)
	ent.move_animation.AddKeyframe(&ent.position.Y, 1.5, centerY+radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.Y, 2.0, centerY)
	ent.move_animation.AddKeyframe(&ent.position.Y, 2.5, centerY-radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.Y, 3.0, centerY-radius)
	ent.move_animation.AddKeyframe(&ent.position.Y, 3.5, centerY-radius*0.7071)
	ent.move_animation.AddKeyframe(&ent.position.Y, 4.0, centerY)

	ent.move_animation.Play(true)
}

func (ent *AnimationTestMoverEntity) Deinit() {
	// De-initialization logic for the entity
}

func (ent *AnimationTestMoverEntity) Update() {
	ent.move_animation.Update()
	rl.DrawCircle(int32(ent.position.X), int32(ent.position.Y), 16, rl.Red)
}
