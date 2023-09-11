package entities

import (
	"cowboy-gorl/pkg/animation"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*AnimationTestOreEntity)(nil)

type OreType int32

const (
	OreType_Coal OreType = iota
	OreType_Iron
)

// AnimationTestOre Entity
type AnimationTestOreEntity struct {
	ore_type            OreType
	base_sprite         rl.Texture2D
	position            rl.Vector2
	click_sounds        []rl.Sound
	special_click_sound rl.Sound
	break_sound         rl.Sound
	click_anim          *animation.Animation[float32]
	click_particles     SpritesheetFxEntity
}

func NewAnimationTestOreEntity(
	ore_type OreType,
	initial_position rl.Vector2,
) AnimationTestOreEntity {
	new_ent := AnimationTestOreEntity{}

	// Set Ore type and corresponding texture
	new_ent.ore_type = ore_type
	switch ore_type {
	case OreType_Coal:
		// NOTE: this could be optimized such that each sprite is only loaded once
		// and referenced multiple times.
		// But that should be the task of some overlying "ore spawner" system.
		new_ent.base_sprite = rl.LoadTexture("sprites/resources/clusters/coal_cluster.png")
	case OreType_Iron:
		new_ent.base_sprite = rl.LoadTexture("sprites/resources/clusters/iron_cluster.png")
	}

	// Set Position
	new_ent.position = initial_position

	// Set interaction sounds
	new_ent.click_sounds = append(new_ent.click_sounds, rl.LoadSound(
		"audio/sounds/foley/metals-ore-lowbit/iron-strike1.ogg",
	))
	new_ent.click_sounds = append(new_ent.click_sounds, rl.LoadSound(
		"audio/sounds/foley/metals-ore-lowbit/iron-strike2.ogg",
	))
	new_ent.click_sounds = append(new_ent.click_sounds, rl.LoadSound(
		"audio/sounds/foley/metals-ore-lowbit/iron-strike3.ogg",
	))
	new_ent.special_click_sound = rl.LoadSound(
		"audio/sounds/foley/metals-ore-lowbit/iron-strike-special1.ogg",
	)
	new_ent.break_sound = rl.LoadSound(
		"audio/sounds/foley/metals-ore-lowbit/iron-strike-special1.ogg",
	)
	return new_ent
}

func (ent *AnimationTestOreEntity) Init() {
	// SHAKE ANIM ON CLICK
	ent.click_anim = animation.CreateAnimation[float32](0.25)
	// Keyframes for X coordinate
	ent.click_anim.AddKeyframe(&ent.position.X, 0.0, ent.position.X)
	ent.click_anim.AddKeyframe(&ent.position.X, 0.05, ent.position.X+1)
	ent.click_anim.AddKeyframe(&ent.position.X, 0.1, ent.position.X-1)
	ent.click_anim.AddKeyframe(&ent.position.X, 0.15, ent.position.X+1)
	ent.click_anim.AddKeyframe(&ent.position.X, 0.2, ent.position.X-1)
	ent.click_anim.AddKeyframe(&ent.position.X, 0.25, ent.position.X)

	// Keyframes for Y coordinate
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.0, ent.position.Y)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.05, ent.position.Y-1)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.1, ent.position.Y+1)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.125, ent.position.Y)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.15, ent.position.Y+1)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.2, ent.position.Y-1)
	ent.click_anim.AddKeyframe(&ent.position.Y, 0.25, ent.position.Y)

	// PARTICLE ANIM ON CLICK
	click_particles_sheet := rl.LoadTexture("sprites/resources/cluster_click_sheets/iron_click1/iron_click1_sheet.png")
	ent.click_particles = NewSpritesheetFxEntity(
		click_particles_sheet,
		rl.NewVector2(64, 64),
		10, 60,
		rl.Vector2Zero(),
	)
	ent.click_particles.Init()
}

func (ent *AnimationTestOreEntity) Deinit() {
	// De-initialization logic for the entity
}

func (ent *AnimationTestOreEntity) Update() {
	rl.DrawTextureV(ent.base_sprite, ent.position, rl.White)

	// Check for a click event
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mouseX := float32(rl.GetMouseX())
		mouseY := float32(rl.GetMouseY())
		// NOTE: entering this bounding box check by hand (especially the 64x64 resolution)
		// is not good. This should be made into its own module
		if mouseX >= ent.position.X && mouseX <= ent.position.X+64 && mouseY >= ent.position.Y && mouseY <= ent.position.Y+64 {
			ent.handleClick()
		}
	}

	ent.click_anim.Update()
	ent.click_particles.Update()
}

func (ent *AnimationTestOreEntity) handleClick() {
	// Play a random click sound
	rand_int := rand.Intn(len(ent.click_sounds))
	rl.PlaySound(ent.click_sounds[rand_int])
	ent.click_anim.Play(false)
	ent.click_particles.Play(rl.Vector2Subtract(rl.GetMousePosition(), rl.NewVector2(32, 32)), false)
}
