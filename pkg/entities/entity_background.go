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
        ent.scroll_pos[i] -= ent.speeds[i] * rl.GetFrameTime() * 100

        // Calculate the scale factor for the texture based on height
        // This ensures the layer is not stretched vertically
        scale := render.Rs.RenderResolution.Y / float32(layer.Height)

        // Determine how many times the texture needs to be drawn to cover the screen
        repeatCount := int(math.Ceil(float64(render.Rs.RenderResolution.X) / (float64(layer.Width) * float64(scale))))

        // Ensure that textures are always in the right position to fill the screen
        for ent.scroll_pos[i] <= -float32(layer.Width)*scale {
            ent.scroll_pos[i] += float32(layer.Width) * scale
        }

        // Draw each layer, taking into account its speed for the parallax effect
        for j := 0; j < repeatCount+2; j++ {  // Draw one more layer to ensure no gaps
            xPosition := ent.scroll_pos[i] + float32(j)*float32(layer.Width)*scale
            rl.DrawTextureEx(layer, rl.NewVector2(xPosition, 0), 0.0, scale, rl.White)
        }
    }
}
