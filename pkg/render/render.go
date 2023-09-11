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
	"cowboy-gorl/pkg/settings"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderState struct {
	TargetTex          rl.RenderTexture2D
	CeilTex            rl.RenderTexture2D
	Backbuffer         rl.RenderTexture2D
	Accumulationbuffer rl.RenderTexture2D
	Blurbuffer         rl.RenderTexture2D
	PingPongCounter    int32
	RenderResolution   rl.Vector2
	RenderScale        rl.Vector2
	MinScale           float32

	Crtshader           rl.Shader
	Crtshader_loc_time  int32
	Blurshader          rl.Shader
	Blurshader_loc_blur int32
	Accumulateshader    rl.Shader
	Blendshader         rl.Shader

	tex1locs map[rl.Shader]int32

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

	// create a number of helper buffers for performing multiple render passes
	// to achieve the accumulation effect of the crt shader
	Rs.Backbuffer = rl.LoadRenderTexture(ceilX, ceilY)
	rl.SetTextureFilter(Rs.Backbuffer.Texture, rl.FilterBilinear)

	Rs.Accumulationbuffer = rl.LoadRenderTexture(ceilX, ceilY)
	rl.SetTextureFilter(Rs.Accumulationbuffer.Texture, rl.FilterBilinear)

	Rs.Blurbuffer = rl.LoadRenderTexture(ceilX, ceilY)
	rl.SetTextureFilter(Rs.Blurbuffer.Texture, rl.FilterBilinear)

	// load the shaders and get their uniform locations if needed
	// crt shader
	Rs.Crtshader = rl.LoadShader("", "shaders/crt-matthias.glsl")
	if Rs.Crtshader.ID == 0 {
		logging.Error("Failed to load CRT shader!")
	}

	Rs.Crtshader_loc_time = rl.GetShaderLocation(Rs.Crtshader, "TIME")
	if Rs.Crtshader_loc_time == 0 {
		logging.Error("Failed to find shader uniform location for TIME in crt-matthias.")
	}

	// blur shader
	Rs.Blurshader = rl.LoadShader("", "shaders/crt-matthias-blur.glsl")
	if Rs.Blurshader.ID == 0 {
		logging.Error("Failed to load CRT blur shader!")
	}
	Rs.Blurshader_loc_blur = rl.GetShaderLocation(Rs.Blurshader, "BLUR")
	if Rs.Blurshader_loc_blur == 0 {
		logging.Error("Failed to find shader uniform location for BLUR in crt-matthias-blur.")
	}
	//rl.SetShaderValue(Rs.Blurshader, Rs.Blurshader_loc_blur, []float32{0.0005, 0.0006}, rl.ShaderUniformVec2)
	rl.SetShaderValue(Rs.Blurshader, Rs.Blurshader_loc_blur, []float32{0.0003, 0.0004}, rl.ShaderUniformVec2)

	// accumulate shader
	Rs.Accumulateshader = rl.LoadShader("", "shaders/crt-matthias-accumulate.glsl")
	if Rs.Accumulateshader.ID == 0 {
		logging.Error("Failed to load CRT accumulate shader!")
	}

	// blend shader
	Rs.Blendshader = rl.LoadShader("", "shaders/crt-matthias-blend.glsl")
	if Rs.Blendshader.ID == 0 {
		logging.Error("Failed to load CRT blend shader!")
	}

	// use a map to store the texture1 locations
	Rs.tex1locs = make(map[rl.Shader]int32)
	Rs.tex1locs[Rs.Accumulateshader] = rl.GetShaderLocation(Rs.Accumulateshader, "texture1")
	Rs.tex1locs[Rs.Blendshader] = rl.GetShaderLocation(Rs.Blendshader, "texture1")

	logging.Info("Custom Rendering Environment initialized.")
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
	rl.DrawTexturePro(Rs.TargetTex.Texture,
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.TargetTex.Texture.Width),
			Height: -float32(Rs.TargetTex.Texture.Height), // flip the texture upside down, OpenGL quirk
		},
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.TargetTex.Texture.Width),
			Height: float32(Rs.TargetTex.Texture.Height),
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White)

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
			Height: float32(Rs.TargetTex.Texture.Height),
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
	if settings.CurrentSettings().EnableCrtEffect {
		rl.SetShaderValue(Rs.Crtshader, Rs.Crtshader_loc_time, []float32{float32(rl.GetTime())}, rl.ShaderUniformFloat)
		// TODO: add a toggle setting between the full shader with accumulate
		// and a "lite" version without accumulation
		renderPass(&Rs.Accumulationbuffer, nil, &Rs.Blurbuffer, &Rs.Blurshader)
		renderPass(&Rs.CeilTex, &Rs.Blurbuffer, &Rs.Accumulationbuffer, &Rs.Accumulateshader)
		renderPass(&Rs.CeilTex, &Rs.Accumulationbuffer, &Rs.Backbuffer, &Rs.Blendshader)
		renderPass(&Rs.Accumulationbuffer, &Rs.Blurbuffer, &Rs.CeilTex, &Rs.Crtshader)
	}
	// render the oversize render texture to the actual screen.
	rl.DrawTexturePro(Rs.CeilTex.Texture,
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.CeilTex.Texture.Width),
			Height: float32(Rs.CeilTex.Texture.Height),
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

func renderPass(tex0 *rl.RenderTexture2D, tex1 *rl.RenderTexture2D, dest *rl.RenderTexture2D, shader *rl.Shader) {
	rl.BeginTextureMode(*dest)
	if shader != nil {
		rl.BeginShaderMode(*shader)
		if tex1 != nil {
			tex1_loc := Rs.tex1locs[*shader]
			rl.SetShaderValueTexture(*shader, tex1_loc, tex1.Texture)
		}
	}
	rl.DrawTexturePro(tex0.Texture,
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.CeilTex.Texture.Width),
			Height: -float32(Rs.CeilTex.Texture.Height),
		},
		rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(Rs.CeilTex.Texture.Width),
			Height: float32(Rs.CeilTex.Texture.Height),
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
	if shader != nil {
		rl.EndShaderMode()
	}
	rl.EndTextureMode()
}
