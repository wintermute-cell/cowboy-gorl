package main

import (
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/render"
	"cowboy-gorl/pkg/scenes"
	"cowboy-gorl/pkg/settings"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// PRE-INIT
	settings_path := "settings.json"
	err := settings.LoadSettings(settings_path)
	if err != nil {
		settings.FallbackSettings()
	}

	logging.Init(settings.CurrentSettings().LogPath)
    logging.Info("Logging initialized")
    if err == nil {
        logging.Info("Settings loaded successfully.")
    } else {
        logging.Warning("Settings loading unsuccessful, using fallback.")
    }

	// INITIALIZATION
    // raylib window
	rl.InitWindow(
		int32(settings.CurrentSettings().ScreenWidth),
		int32(settings.CurrentSettings().ScreenHeight), "cowboy-gorl window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(int32(settings.CurrentSettings().TargetFps))

    // rendering
	render.Init(
		settings.CurrentSettings().RenderWidth,
		settings.CurrentSettings().RenderHeight)
    logging.Info("Custom rendering initialized.")

    // initialize the audio device
    rl.InitAudioDevice()
    defer rl.CloseAudioDevice()

    // raygui
    font := rl.LoadFont("fonts/alagard.png")
    rl.SetTextureFilter(font.Texture, rl.FilterPoint);
    rg.SetFont(font)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_NORMAL, 0xffffffff)
    rg.SetStyle(rg.DEFAULT, rg.BASE_COLOR_NORMAL, 0xaa7a39ff)
    rg.SetStyle(rg.DEFAULT, rg.BACKGROUND_COLOR, 0xca9a39ff)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(font.BaseSize))
    rg.SetStyle(rg.DEFAULT, rg.BORDER_WIDTH, 0)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SPACING, 2)

    // SCENES (this part MUST come after the system/rl initialization)
    // ( don't forget to defer the Deinit()!!! )
	dev_scene := scenes.DevScene{}
	dev_scene.Init()
	defer dev_scene.Deinit()

	// GAME LOOP
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black) // clearing the whole background, even behind the main rendertex

		render.BeginCustomRender()
        rl.BeginDrawing()
            rl.ClearBackground(rl.DarkGreen) // clear the main rendertex

            // Draw Content
            dev_scene.Draw()
            //rl.DrawTextEx(font, "Some Text", rl.NewVector2(0, 0), float32(font.BaseSize)*2.0, 2, rl.White)
            //rg.Label(rl.NewRectangle(0, 32, 200, 16), "Some Text");

            // Draw GUI
            dev_scene.DrawGUI()

        render.EndCustomRender()

        // Draw Debug Info
        rl.DrawFPS(10, 10)
        rl.DrawGrid(10, 1.0)

        rl.EndDrawing()

	}
}
