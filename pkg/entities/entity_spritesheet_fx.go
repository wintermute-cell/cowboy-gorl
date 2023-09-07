package entities

import (
	"cowboy-gorl/pkg/animation"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*SpritesheetFxEntity)(nil)

//
//  SpritesheetFx Entity
//
type SpritesheetFxEntity struct {
    spritesheet rl.Texture2D
    frame_dimensions rl.Vector2
    num_frames int32
    current_frame int32
    ms_per_frame int32
    position rl.Vector2

    frame_animation *animation.Animation[int32]
    is_playing bool
}

func NewSpritesheetFxEntity(
    spritesheet rl.Texture2D,
    frame_dimensions rl.Vector2,
    num_frames int32,
    ms_per_frame int32,
    position rl.Vector2,
) SpritesheetFxEntity {
    new_ent := SpritesheetFxEntity{}

    new_ent.spritesheet = spritesheet
    new_ent.frame_dimensions = frame_dimensions
    new_ent.num_frames = num_frames
    new_ent.ms_per_frame = ms_per_frame
    new_ent.position = position
    new_ent.current_frame = 0
    new_ent.is_playing = false

    return new_ent
}

func (ent *SpritesheetFxEntity) Init() {
    sec_per_frame := float32(ent.ms_per_frame)/1000
    ent.frame_animation = animation.CreateAnimation[int32](
        float32(ent.num_frames)*sec_per_frame,
        )

    // one keyframe per spritesheet frame
    for i := float32(0); i < float32(ent.num_frames); i++ {
        ent.frame_animation.AddKeyframe(
            &ent.current_frame,
            sec_per_frame*i,
            int32(i),
            )
    }
}

func (ent *SpritesheetFxEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *SpritesheetFxEntity) Update() {
    if ent.is_playing {
        rl.DrawTextureRec(
            ent.spritesheet,
            rl.NewRectangle(
                float32(ent.current_frame)*ent.frame_dimensions.X, 0.0,
                ent.frame_dimensions.X, ent.frame_dimensions.Y),
            ent.position,
            rl.White,
            )
        ent.frame_animation.Update()
    }
    if ent.current_frame >= ent.num_frames-1 {
        ent.is_playing = false
        ent.current_frame = 0
    }
}

func (ent *SpritesheetFxEntity) Play(position rl.Vector2, looping bool) {
    ent.position = position
    ent.frame_animation.Play(looping)
    ent.is_playing = true
}
