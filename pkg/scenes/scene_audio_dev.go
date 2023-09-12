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
		audio.SetCurrentPlaylist("main-menu", true)
	}
    btn_playlist_main_menu := gui.NewButton(
        "Playlist: main-menu",
        rl.NewVector2(16, 96),
        rl.NewVector2(15*6, 16),
        btn_cb_playlist_main_menu,
        "background-pressed:180,10,10,255")

	btn_cb_playlist_boss_music := func(s gui.ButtonState) {
		audio.SetCurrentPlaylist("boss-music", true)
	}
    btn_playlist_boss_music := gui.NewButton(
        "Playlist: boss-music",
        rl.NewVector2(16, 124),
        rl.NewVector2(15*6, 16),
        btn_cb_playlist_boss_music,
        "background-pressed:180,10,10,255")

	scroll_panel := gui.NewScrollPanel(
		rl.NewRectangle(10, 48, 620, 400),
		rl.NewRectangle(10, 48, 620, 2000),
		"debug:false|background:200,200,200,255")

	scn.g.AddWidget(scroll_panel)
	scroll_panel.AddChild(btn_playlist_main_menu)
	scroll_panel.AddChild(btn_playlist_boss_music)
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
