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
    rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_NORMAL, 0xffffffff)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 14)

    // SCENES (this part MUST come after the system/rl initialization)
    // ( don't forget to defer the Deinit()!!! )
    scene_manager := scenes.NewSceneManager()
    scene_manager.RegisterScene("dev_scene", &scenes.DevScene{})
    scene_manager.RegisterScene("anim_dev_scene", &scenes.AnimationDevScene{})

    scene_manager.EnableScene("anim_dev_scene")

	// GAME LOOP
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black) // clearing the whole background, even behind the main rendertex

		render.BeginCustomRender()
        rl.BeginDrawing()

            rl.ClearBackground(rl.DarkGreen) // clear the main rendertex

            // Draw all registered Scenes
            scene_manager.DrawScenes()
            scene_manager.DrawScenesGUI()

        render.EndCustomRender()

        // Draw Debug Info
        rl.DrawFPS(10, 10)
        rl.DrawGrid(10, 1.0)

        rl.EndDrawing()
	}
    
    scene_manager.DisableAllScenes()
}
