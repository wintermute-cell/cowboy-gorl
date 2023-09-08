package scenes

import (
	"cowboy-gorl/pkg/entities"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Scene = (*GuiDevScene)(nil)

//
//  GuiDev Scene
//
type GuiDevScene struct {
    // Required fields
    entity_manager *entities.EntityManager

    // Custom Fields
    alagard_font rl.Font
    pixantiqua_font rl.Font
    pixelplay_font rl.Font
    romulus_font rl.Font

    scroll_offset rl.Vector2
}

func (scn *GuiDevScene) Init() {
    // Required initialization
    scn.entity_manager = entities.NewEntityManager()

    scn.alagard_font = rl.LoadFont("fonts/alagard.png")
    rl.SetTextureFilter(scn.alagard_font.Texture, rl.FilterPoint);

    scn.pixantiqua_font = rl.LoadFont("fonts/pixantiqua.png")
    rl.SetTextureFilter(scn.pixantiqua_font.Texture, rl.FilterPoint);

    scn.pixelplay_font = rl.LoadFont("fonts/pixelplay.png")
    rl.SetTextureFilter(scn.pixelplay_font.Texture, rl.FilterPoint);

    scn.romulus_font = rl.LoadFont("fonts/romulus.png")
    rl.SetTextureFilter(scn.romulus_font.Texture, rl.FilterPoint);

    scn.scroll_offset = rl.Vector2Zero()
}

func (scn *GuiDevScene) Deinit() {
    // De-initialization logic for the scene
}

func (scn *GuiDevScene) DrawGUI() {
    line_height := 48
    i := float32(2.0)

    so := scn.scroll_offset // just a shorter alias for Scroll Offset
    view := rg.ScrollPanel(rl.NewRectangle(10, 48, 620, 400), "GUI Label Examples", rl.NewRectangle(16, 48, 600, 800), &scn.scroll_offset)

    // we use the view returned by rg.ScrollPanel to scissor out the out of bounds content. 
    rl.BeginScissorMode(int32(view.X), int32(view.Y), int32(view.Width), int32(view.Height))

    // each element "inside" the scroll panel has to be offset by the current scroll offset vector
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "Gui Label Without Changes")
    i += 1

    original_font := rg.GetFont()
    rg.SetFont(original_font)
    // NOTE: notice that SetFont changes the TEXT_SIZE propery

    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "after SetFont")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 24)
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "after SetFont and setting TEXT_SIZE")
    i += 1

    rl.SetTextureFilter(original_font.Texture, rl.FilterPoint);
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "after setting Point Filtering")
    i += 1

    rg.SetFont(scn.alagard_font)
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with alagard font")
    i += 1

    rg.Label(rl.NewRectangle(so.X+17, so.Y+float32(line_height)*i, 600, 24), "with alagard font shifted +1px right")
    i += 1

    // NOTE: scaling a pixelfont with non-integer values (e.g. 1.65) is no good
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(float32(scn.alagard_font.BaseSize)*1.65))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with alagard font at BaseSize*1.65")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.alagard_font.BaseSize))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with alagard font at BaseSize")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.alagard_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with alagard font at BaseSize*2")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.alagard_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+15, so.Y+float32(line_height)*i, 600, 24), "with alagard font at BaseSize*2 shifted -1px left")
    i += 1

    orig_spacing := rg.GetStyle(rg.DEFAULT, rg.TEXT_SPACING)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SPACING, 3)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.alagard_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with alagard *2 with spacing=3")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SPACING, orig_spacing)
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SPACING, 2)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.alagard_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+15, so.Y+float32(line_height)*i, 600, 24), "alagard *2, spacing=2, -1px left")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SPACING, orig_spacing)
    i += 1

    rg.SetFont(scn.pixantiqua_font)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.pixantiqua_font.BaseSize))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with pixantiqua font at BaseSize")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.pixantiqua_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with pixantiqua font at BaseSize*2")
    i += 1

    rg.SetFont(scn.pixelplay_font)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.pixelplay_font.BaseSize))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with pixelplay font at BaseSize")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.pixelplay_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with pixelplay font at BaseSize*2")
    i += 1

    rg.SetFont(scn.romulus_font)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.romulus_font.BaseSize))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with romulus font at BaseSize")
    i += 1

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, int64(scn.romulus_font.BaseSize*2.0))
    rg.Label(rl.NewRectangle(so.X+16, so.Y+float32(line_height)*i, 600, 24), "with romulus font at BaseSize*2")
    i += 1

    rl.EndScissorMode()

    // resetting the used font for next frame
    rg.SetFont(original_font)
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 24)

    // pixel ruler // 640 480
    offs := int32(7)
    rl.DrawLine(0, 480, 0, 300, rl.Red)
    rl.DrawLine(1, 480, 1, 300, rl.Green)
    rl.DrawLine(offs*1, 480, offs*1, 300, rl.Red)
    rl.DrawLine(offs*2, 480, offs*2, 300, rl.Red)
    rl.DrawLine(offs*3, 480, offs*3, 300, rl.Red)
    rl.DrawLine(offs*4, 480, offs*4, 300, rl.Red)
    rl.DrawLine(offs*5, 480, offs*5, 300, rl.Red)
}

func (scn *GuiDevScene) Draw() {
    // Draw the scene
}
