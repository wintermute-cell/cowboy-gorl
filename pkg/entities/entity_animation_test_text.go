package entities

import (
	"cowboy-gorl/pkg/animation"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*AnimationTestTextEntity)(nil)

// AnimationTestText Entity
type AnimationTestTextEntity struct {
	position       rl.Vector2
	text           string
	text_animation *animation.Animation[string]
}

func (ent *AnimationTestTextEntity) Init() {
	// Initialization logic for the entity
	ent.text_animation = animation.CreateAnimation[string](4.0)
	ent.position = rl.NewVector2(32, 127)

	ent.text_animation.AddKeyframe(&ent.text, 0.0, ent.text)
	ent.text_animation.AddKeyframe(&ent.text, 1.0, "Anim")
	ent.text_animation.AddKeyframe(&ent.text, 2.0, "Animated")
	ent.text_animation.AddKeyframe(&ent.text, 3.0, "AnimatedText")
	ent.text_animation.AddKeyframe(&ent.text, 4.0, "")

	ent.text_animation.Play(true)
}

func (ent *AnimationTestTextEntity) Deinit() {
	// De-initialization logic for the entity
}

func (ent *AnimationTestTextEntity) Update() {
	ent.text_animation.Update()
	rl.DrawText(ent.text, int32(ent.position.X), int32(ent.position.Y), 24, rl.White)
}
