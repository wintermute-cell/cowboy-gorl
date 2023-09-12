package main

import (
	"cowboy-gorl/pkg/audio"
	"cowboy-gorl/pkg/gui"
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

	// initialize audio
    audio.InitAudio()
    defer audio.DeinitAudio()

    // register audio tracks
    audio.RegisterMusic("aza-tumbleweeds", "audio/music/azakaela/azaFMP2_field7_Tumbleweeds.ogg")
    audio.RegisterMusic("aza-outwest", "audio/music/azakaela/azaFMP2_scene1_OutWest.ogg")
    audio.RegisterMusic("aza-frontier", "audio/music/azakaela/azaFMP2_town_Frontier.ogg")
    audio.CreatePlaylist("main-menu", []string{"aza-tumbleweeds", "aza-outwest", "aza-frontier"})

	// gui
	gui.InitBackend()

	// raygui
	rg.SetStyle(rg.DEFAULT, rg.TEXT_COLOR_NORMAL, 0x000000)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 24)

	// SCENES (this part MUST come after the system/rl initialization)
	// ( don't forget to defer the Deinit()!!! )
	scenes.Sm.RegisterScene("dev", &scenes.DevScene{})
	scenes.Sm.RegisterScene("anim_dev", &scenes.AnimationDevScene{})
	scenes.Sm.RegisterScene("gui_dev", &scenes.GuiDevScene{})
	scenes.Sm.RegisterScene("dev_menu", &scenes.DevMenuScene{})
    scenes.Sm.RegisterScene("audio_dev", &scenes.AudioDevScene{})

	scenes.Sm.EnableScene("dev_menu")

	// GAME LOOP
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black) // clearing the whole background, even behind the main rendertex

		render.BeginCustomRender()
		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGreen) // clear the main rendertex

		// Draw all registered Scenes
		scenes.Sm.DrawScenes()
		scenes.Sm.DrawScenesGUI()

		render.EndCustomRender()

		// Draw Debug Info
		rl.DrawFPS(10, 10)
		rl.DrawGrid(10, 1.0)

		rl.EndDrawing()

        audio.Update()
	}

	scenes.Sm.DisableAllScenes()
}
