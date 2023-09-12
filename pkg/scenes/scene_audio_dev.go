package scenes

import (
	"cowboy-gorl/pkg/audio"
	"cowboy-gorl/pkg/entities"
	"cowboy-gorl/pkg/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*AudioDevScene)(nil)

// GuiDev Scene
type AudioDevScene struct {
	// Required fields
	entity_manager *entities.EntityManager

	// Custom Fields
	g *gui.Gui
}

func (scn *AudioDevScene) Init() {
	// Required initialization
	scn.entity_manager = entities.NewEntityManager()

	scn.g = gui.NewGui()
                    
    // register audio tracks and make playlists
    audio.RegisterMusic("aza-tumbleweeds", "audio/music/azakaela/azaFMP2_field7_Tumbleweeds.ogg")
    audio.RegisterMusic("aza-outwest", "audio/music/azakaela/azaFMP2_scene1_OutWest.ogg")
    audio.RegisterMusic("aza-frontier", "audio/music/azakaela/azaFMP2_town_Frontier.ogg")
    audio.CreatePlaylist("main-menu", []string{"aza-tumbleweeds", "aza-outwest", "aza-frontier"})

    audio.RegisterMusic("aza-gunset", "audio/music/azakaela/azaFMP2_bossbattle1_Gunset.ogg")
    audio.RegisterMusic("aza-the-engine-of-destruction", "audio/music/azakaela/azaFMP2_bossbattle2_TheEngineOfDestruction.ogg")
    audio.RegisterMusic("aza-the-revenant", "audio/music/azakaela/azaFMP2_bossbattle3_TheRevenant.ogg")
    audio.CreatePlaylist("boss-music", []string{"aza-gunset", "aza-the-engine-of-destruction", "aza-the-revenant"})

	// gui elements
	btn_cb_playlist_main_menu := func(s gui.ButtonState) {
        if s == gui.ButtonStateReleased {
            audio.SetCurrentPlaylist("main-menu", true)
        }
	}
    btn_playlist_main_menu := gui.NewButton(
        "Playlist: main-menu",
        rl.NewVector2(16, 96),
        rl.NewVector2(15*8, 16),
        btn_cb_playlist_main_menu,
        "background-pressed:180,10,10,255")

	btn_cb_playlist_boss_music := func(s gui.ButtonState) {
        if s == gui.ButtonStateReleased {
            audio.SetCurrentPlaylist("boss-music", true)
        }
	}
    btn_playlist_boss_music := gui.NewButton(
        "Playlist: boss-music",
        rl.NewVector2(16, 124),
        rl.NewVector2(15*8, 16),
        btn_cb_playlist_boss_music,
        "background-pressed:180,10,10,255")

    slider_global_volume := gui.NewSlider(
        0.0, 1.0, audio.GetGlobalVolume(), 0.1,
        rl.NewVector2(200, 96),
        rl.NewVector2(200, 16),
        rl.NewVector2(16, 16),
        "")
    slider_global_volume.SetValueChangedCallback(func(new_value float32) {audio.SetGlobalVolume(new_value)})
    label_global_volume := gui.NewLabel("", rl.NewVector2(420, 96), "")
    label_global_volume.WatchFloat32(slider_global_volume.GetCurrentValuePointer(), "%.2f")

    slider_music_volume := gui.NewSlider(
        0.0, 1.0, audio.GetMusicVolume(), 0.1,
        rl.NewVector2(200, 128),
        rl.NewVector2(200, 16),
        rl.NewVector2(16, 16),
        "")
    slider_music_volume.SetValueChangedCallback(func(new_value float32) {audio.SetMusicVolume(new_value)})
    label_music_volume := gui.NewLabel("", rl.NewVector2(420, 128), "")
    label_music_volume.WatchFloat32(slider_music_volume.GetCurrentValuePointer(), "%.2f")

    slider_sfx_volume := gui.NewSlider(
        0.0, 1.0, audio.GetSFXVolume(), 0.1,
        rl.NewVector2(200, 150),
        rl.NewVector2(200, 16),
        rl.NewVector2(16, 16),
        "")
    slider_sfx_volume.SetValueChangedCallback(func(new_value float32) {audio.SetSFXVolume(new_value)})
    label_sfx_volume := gui.NewLabel("", rl.NewVector2(420, 150), "")
    label_sfx_volume.WatchFloat32(slider_sfx_volume.GetCurrentValuePointer(), "%.2f")

	scroll_panel := gui.NewScrollPanel(
		rl.NewRectangle(10, 48, 620, 400),
		rl.NewRectangle(10, 48, 620, 2000),
		"debug:false|background:200,200,200,255")

	scn.g.AddWidget(scroll_panel)
	scroll_panel.AddChild(btn_playlist_main_menu)
	scroll_panel.AddChild(btn_playlist_boss_music)
	scroll_panel.AddChild(slider_global_volume)
	scroll_panel.AddChild(slider_music_volume)
	scroll_panel.AddChild(slider_sfx_volume)
    scroll_panel.AddChild(label_global_volume)
    scroll_panel.AddChild(label_music_volume)
    scroll_panel.AddChild(label_sfx_volume)
}

func (scn *AudioDevScene) Deinit() {
	// De-initialization logic for the scene
}

func (scn *AudioDevScene) DrawGUI() {
	scn.g.Draw()
}

func (scn *AudioDevScene) Draw() {
	// Draw the scene
}
