package entities

import (
	"cowboy-gorl/pkg/render"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*BackgroundEntity)(nil)

//
//  Background Entity
//
type BackgroundEntity struct {
    layers []rl.Texture2D
    speeds []float32
    scroll_pos []float32
}

func (ent *BackgroundEntity) Init() {
}

func (ent *BackgroundEntity) SetLayers(layer_paths []string, layer_speeds []float32) {
    for i, path := range layer_paths {
        ent.layers = append(ent.layers, rl.LoadTexture(path))
        ent.speeds = append(ent.speeds, layer_speeds[i])
        ent.scroll_pos = append(ent.scroll_pos, 0.0)
    }
}

func (ent *BackgroundEntity) Deinit() {
    for _, tex := range ent.layers {
        rl.UnloadTexture(tex)
    }
}

func (ent *BackgroundEntity) Update() {
    for i, layer := range ent.layers {
        ent.scroll_pos[i] -= ent.speeds[i]

        // Calculate the scale factor for the texture
        scaleX := render.Rs.RenderResolution.X / float32(layer.Width)
        scaleY := render.Rs.RenderResolution.Y / float32(layer.Height)
        scale := float32(math.Max(float64(scaleX), float64(scaleY)))

        // Reset scroll if necessary
        if ent.scroll_pos[i] <= -float32(layer.Width)*scale {
            ent.scroll_pos[i] = 0
        }

        // Draw each layer twice, taking into account its speed for the parallax effect
        rl.DrawTextureEx(layer, rl.NewVector2(ent.scroll_pos[i], 0), 0.0, scale, rl.White)
        rl.DrawTextureEx(layer, rl.NewVector2(float32(layer.Width)*scale + ent.scroll_pos[i], 0), 0.0, scale, rl.White)
    }
}
