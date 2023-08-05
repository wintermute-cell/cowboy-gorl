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
    locomotive_tex rl.Texture2D
}

func (ent *TrainEntity) Init() {
    ent.locomotive_tex = rl.LoadTexture("train/train.png")
}

func (ent *TrainEntity) Deinit() {
    rl.UnloadTexture(ent.locomotive_tex)
}

func (ent *TrainEntity) Update() {
    rl.DrawTextureEx(ent.locomotive_tex,
        rl.NewVector2(0.0, render.Rs.RenderResolution.Y-float32(render.PY(0.2))), // position
        0.0, // rotation
        10.0, // scale
        rl.White)
}

