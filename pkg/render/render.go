package render

import (
	"cowboy-gorl/pkg/util"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderState struct {
    TargetTex        rl.RenderTexture2D
    RenderResolution rl.Vector2
    RenderScale      rl.Vector2
    MinScale         float32
}

var Rs RenderState

func Init(render_width int, render_height int) {
	Rs = RenderState{
		RenderResolution: rl.NewVector2(
			float32(render_width),
			float32(render_height)),
	}
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
    if Rs.TargetTex.ID == 0 {
        Rs.TargetTex = rl.LoadRenderTexture(int32(Rs.RenderResolution.X), int32(Rs.RenderResolution.Y))
        rl.SetTextureFilter(Rs.TargetTex.Texture, rl.FilterPoint)
    }
    scale := rl.Vector2{
        X: float32(rl.GetScreenWidth()) / Rs.RenderResolution.X,
        Y: float32(rl.GetScreenHeight()) / Rs.RenderResolution.Y,
    }
    Rs.MinScale = float32(math.Min(float64(scale.X), float64(scale.Y)))

    rl.BeginTextureMode(Rs.TargetTex)
    Rs.RenderScale = scale

    mouse := rl.GetMousePosition()
    virtualMouse := rl.Vector2{}
    virtualMouse.X = (mouse.X - float32(rl.GetScreenWidth())-(Rs.RenderResolution.X*Rs.MinScale)*0.5) / Rs.MinScale
    virtualMouse.Y = (mouse.Y - float32(rl.GetScreenHeight())-(Rs.RenderResolution.Y*Rs.MinScale)*0.5) / Rs.MinScale
    virtualMouse = util.Vector2Clamp(virtualMouse, rl.Vector2{}, Rs.RenderResolution)

    rl.SetMouseOffset(int(-(float32(rl.GetScreenWidth()) - Rs.RenderResolution.X*Rs.MinScale)*0.5),
        int(-(float32(rl.GetScreenHeight()) - Rs.RenderResolution.Y*Rs.MinScale)*0.5))
    rl.SetMouseScale(1/Rs.MinScale, 1/Rs.MinScale)
}

func EndCustomRender() {
    rl.EndTextureMode()
    rl.DrawTexturePro(Rs.TargetTex.Texture,
        rl.Rectangle{
            X: 0.0,
            Y: 0.0,
            Width: float32(Rs.TargetTex.Texture.Width),
            Height: -float32(Rs.TargetTex.Texture.Height),
        },
        rl.Rectangle{
            X: (float32(rl.GetScreenWidth()) - Rs.RenderResolution.X*Rs.MinScale) * 0.5,
            Y: (float32(rl.GetScreenHeight()) - Rs.RenderResolution.Y*Rs.MinScale) * 0.5,
            Width: Rs.RenderResolution.X * Rs.MinScale,
            Height: Rs.RenderResolution.Y * Rs.MinScale},
        rl.Vector2{X:0, Y:0}, 0, rl.White) 
}
