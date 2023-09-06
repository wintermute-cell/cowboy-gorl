/*
-----------------------------------------------------------
| Multi-Stage Custom Renderer for                         |
-----------------------------------------------------------

Stage 1: Initial Render (TargetTex)
  - Draws sprites at a fixed resolution.
  - Uses Point Sampling for pixel-perfect scaling.

Stage 2: Intermediary Upscale (CeilTex)
  - Scales up TargetTex using integer scaling.
  - Purpose: Scale as close to output resolution while maintaining sharp edges.
  - Uses Point Sampling for crisp edges.

Stage 3: Final Render (To Screen)
  - Scales down CeilTex to fit screen size.
  - Purpose: Mask subpixel inconsistencies.
  - Uses Bilinear Sampling for smoother edges.

*/

package render

import (
	"cowboy-gorl/pkg/logging"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderState struct {
	TargetTex        rl.RenderTexture2D
	CeilTex          rl.RenderTexture2D
	RenderResolution rl.Vector2
	RenderScale      rl.Vector2
	MinScale         float32
	Crtshader        rl.Shader
    lastScreenHeight int32
}

var Rs RenderState

func Init(render_width int, render_height int) {
	Rs = RenderState{
		RenderResolution: rl.NewVector2(
			float32(render_width),
			float32(render_height)),
	}

    // this is used to detect if the screen is resized
    Rs.lastScreenHeight = int32(rl.GetScreenHeight())

    recalcScaleFactor()

	// create the primary render texture. all sprites will be drawn directly
	// to this texture.
	Rs.TargetTex = rl.LoadRenderTexture(
		int32(Rs.RenderResolution.X),
		int32(Rs.RenderResolution.Y))
	rl.SetTextureFilter(Rs.TargetTex.Texture, rl.FilterPoint)

	// create a secondary render texture, that is the next integer scaling step
	// larger than the window resolution.
	ceilX := int32(Rs.RenderResolution.X) * int32(math.Ceil(float64(Rs.MinScale)))
	ceilY := int32(Rs.RenderResolution.Y) * int32(math.Ceil(float64(Rs.MinScale)))
	Rs.CeilTex = rl.LoadRenderTexture(ceilX, ceilY)
	rl.SetTextureFilter(Rs.CeilTex.Texture, rl.FilterBilinear)
	logging.Info("Custom Rendering Environment initialized.")

	Rs.Crtshader = rl.LoadShader("", "shaders/crt-matthias.fs")
	if Rs.Crtshader.ID == 0 {
		logging.Error("Failed to load CRT shader!")
	}
}

func Deinit() {
	rl.UnloadShader(Rs.Crtshader)
}

// positioning functions
func PX(percentage_position float32) int32 {
	min := int32(Rs.RenderResolution.X)
	return int32(percentage_position * float32(min))
}

func PY(percentage_position float32) int32 {
	min := int32(Rs.RenderResolution.Y)
	return int32(percentage_position * float32(min))
}

func BeginCustomRender() {
    if Rs.lastScreenHeight != int32(rl.GetScreenHeight()) {
        recalcScaleFactor()
        Rs.lastScreenHeight = int32(rl.GetScreenHeight())
    }

	// adjust mouse coordinates so that they match with the render resolution,
	// not the screen size.
    // NOTE: We might be able to move the following into recalcScaleFactor():
	rl.SetMouseOffset(int(-(float32(rl.GetScreenWidth())-Rs.RenderResolution.X*Rs.MinScale)*0.5),
		int(-(float32(rl.GetScreenHeight())-Rs.RenderResolution.Y*Rs.MinScale)*0.5))
	rl.SetMouseScale(1/Rs.MinScale, 1/Rs.MinScale)

	// begin rendering to the primary render texture
	rl.BeginTextureMode(Rs.TargetTex)
}

func EndCustomRender() {
	rl.EndTextureMode()

	// render the contents of the primary render texture to the slightly
	// oversized (but integer scaled) intermediary render texture.
	// this texture uses bilinear sampling.
	rl.BeginTextureMode(Rs.CeilTex)
	rl.DrawTexturePro(Rs.TargetTex.Texture,
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.TargetTex.Texture.Width),
			Height: -float32(Rs.TargetTex.Texture.Height),
		},
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.CeilTex.Texture.Width),
			Height: float32(Rs.CeilTex.Texture.Height),
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White)
	rl.EndTextureMode()

	// NOTE: Not sure if the shader should also render over the GUI, or if
	// we need to separate World and GUI into individual render textures.
	rl.BeginShaderMode(Rs.Crtshader)
	// render the oversize render texture to the actual screen.
	rl.DrawTexturePro(Rs.CeilTex.Texture,
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.CeilTex.Texture.Width),
			Height: -float32(Rs.CeilTex.Texture.Height),
		},
		rl.Rectangle{
			// the position calculations for X and Y are in place, so the
			// texture is rendered in the middle of the screen, in case the
			// aspect ratio does not match
			X:      (float32(rl.GetScreenWidth()) - Rs.RenderResolution.X*Rs.MinScale) * 0.5,
			Y:      (float32(rl.GetScreenHeight()) - Rs.RenderResolution.Y*Rs.MinScale) * 0.5,
			Width:  Rs.RenderResolution.X * Rs.MinScale,
			Height: Rs.RenderResolution.Y * Rs.MinScale,
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
	rl.EndShaderMode()
}

func recalcScaleFactor() {
	Rs.RenderScale = rl.Vector2{
		X: float32(rl.GetScreenWidth()) / Rs.RenderResolution.X,
		Y: float32(rl.GetScreenHeight()) / Rs.RenderResolution.Y,
	}
	Rs.MinScale = float32(math.Min(
		float64(Rs.RenderScale.X),
		float64(Rs.RenderScale.Y)))
}
