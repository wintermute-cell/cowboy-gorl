package main

import (
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/render"
	"cowboy-gorl/pkg/scenes"
	"cowboy-gorl/pkg/settings"


	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Log *logging.Log = new(logging.Log)
)

func main() {

	// PRE-INIT
	settings_path := "settings.json"
	err := settings.LoadSettings(settings_path)
	if err != nil {
		settings.FallbackSettings()
		
		// FIXME: kann nicht geloggt werden, da der logging-Speicherort erst
		// aus den Settings gelesen werden muss!
		// FIXME: eigentlich soll laut issue ja logging.go sich den filepath 
		// aus den settings selbst holen, dann wuedern wir das hier umgehen,
		// zumindest wenn das so gemeint ist dass man den log_path aus der
		// settings .json holen soll
		//InfoLogger.Printf("Failed to load settings from '%s', using fallback!",
		//	settings_path)
	}

 	// FIXME: liesst nicht korrekt auf json
	Log.Init(settings.CurrentSettings().LogPath)
	Log.Info("Logges is set up") // FIXME: nur ein Beispiel

	// INITIALIZATION
	rl.InitWindow(
		int32(settings.CurrentSettings().ScreenWidth),
		int32(settings.CurrentSettings().ScreenHeight), "cowboy-gorl window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(int32(settings.CurrentSettings().TargetFps))

	render.Init(
		settings.CurrentSettings().RenderWidth,
		settings.CurrentSettings().RenderHeight)

	dev_scene := scenes.DevScene{}
	dev_scene.Init()
	defer dev_scene.Deinit()

	// GAME LOOP
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black)
		render.BeginCustomRender()
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGreen)

		// Draw GUI
		//dev_scene.DrawGUI()

		dev_scene.Draw()
		rl.DrawFPS(10, 10)
		rl.DrawGrid(10, 1.0)

		rl.EndDrawing()
		render.EndCustomRender()
	}
}
