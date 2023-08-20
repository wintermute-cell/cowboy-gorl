package entities

import (
	"cowboy-gorl/pkg/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*TrainEntity)(nil)

//
//  Train Entity
//
type TrainEntity struct {
    locomotive_sound rl.Music
    locomotive_tex rl.Texture2D
}

func (ent *TrainEntity) Init() {
    ent.locomotive_sound = rl.LoadMusicStream("audio/locomotive/ES_Train Steam 2.ogg")
	rl.PlayMusicStream(ent.locomotive_sound)
	rl.SetMusicVolume(ent.locomotive_sound, 0.6)
    ent.locomotive_tex = rl.LoadTexture("sprites/locomotive.png")
}

func (ent *TrainEntity) Deinit() {
    rl.UnloadTexture(ent.locomotive_tex)
    rl.UnloadMusicStream(ent.locomotive_sound)
}

func (ent *TrainEntity) Update() {
    rl.UpdateMusicStream(ent.locomotive_sound);
    rl.DrawTextureEx(ent.locomotive_tex,
        rl.NewVector2(-60.0, render.Rs.RenderResolution.Y-(72*1.8)-86), // position
        0.0, // rotation
        1.8, // scale
        rl.White)
}

