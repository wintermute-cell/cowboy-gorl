package audio

import (
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type audio struct {
    // volume
    is_mute bool
    global_volume float32
    music_volume float32
    sfx_volume float32

    // tracks
    music_tracks map[string]rl.Music
    sfx_tracks map[string]rl.Sound

    // playback
    music_fade_secs float32
    fading_music_track_name string
    curr_playing_music_name string
}

var a audio

func InitAudio() {
	rl.InitAudioDevice()
    if !rl.IsAudioDeviceReady() {
        logging.Fatal("Failed to InitAudioDevice!")
    }
    a = audio{
        is_mute: false,
        global_volume: 1.0,
        music_volume: 1.0,
        sfx_volume: 1.0,
        music_tracks: make(map[string]rl.Music),
        sfx_tracks: make(map[string]rl.Sound),
        music_fade_secs: 5.0,
    }
}

func DeinitAudio() {
    // unload all audio tracks from memory
    for _, sound := range a.sfx_tracks {
        rl.UnloadSound(sound)
    }
    for _, music := range a.music_tracks {
        rl.UnloadMusicStream(music)
    }

	rl.CloseAudioDevice()
}

func Update() {
    if !a.is_mute {
        // if were currently playing music...
        if curr_mus, ok := a.music_tracks[a.curr_playing_music_name]; ok {
            rl.UpdateMusicStream(curr_mus)

            // get the remaining playtime
            time_remaining := rl.GetMusicTimeLength(curr_mus) - rl.GetMusicTimePlayed(curr_mus)

            // fade in the current music
            if rl.GetMusicTimePlayed(curr_mus) <= a.music_fade_secs {
                rl.SetMusicVolume(curr_mus, a.music_volume * a.global_volume * (rl.GetMusicTimePlayed(curr_mus) / a.music_fade_secs))
            } else {
                // apply the configured music volume
                rl.SetMusicVolume(curr_mus, a.music_volume * a.global_volume)
            }

            // determine if the current music needs to be faded out
            if time_remaining <= a.music_fade_secs {
                logging.Info("Starting to fade out music: %v", a.curr_playing_music_name)
                // set the current track to the fade var and the new one to the playing var
                a.fading_music_track_name = a.curr_playing_music_name
                new_mus := getNextMusicTrack()
                rl.SetMusicVolume(a.music_tracks[new_mus], a.music_volume * a.global_volume)
                a.curr_playing_music_name = new_mus
                // TODO: fading into the same track breaks the system. add functonality to just loop a track (but still fade the loops).
                rl.PlayMusicStream(a.music_tracks[new_mus])
                logging.Info("Starting to play music: %v", new_mus)
            }
        }

        // fade out the last music track
        if fading_mus, ok := a.music_tracks[a.fading_music_track_name]; ok {
            rl.UpdateMusicStream(fading_mus)
            if rl.IsMusicStreamPlaying(fading_mus) {
                fade_time_remaining := rl.GetMusicTimeLength(fading_mus) - rl.GetMusicTimePlayed(fading_mus)
                logging.Debug("%v", fade_time_remaining)
                rl.SetMusicVolume(fading_mus, a.music_volume * a.global_volume * util.Max((fade_time_remaining / a.music_fade_secs), 0.01))
                if fade_time_remaining <= 0.01 {
                    logging.Info("Finished fading out music: %v", a.fading_music_track_name)
                    rl.StopMusicStream(fading_mus)
                }
            }
        }
    }
}

func getNextMusicTrack() string {
    // TODO
    return "aza-tumbleweeds"
}

// ----------------
//       API      |
// ----------------


// LOADING TRACKS

// Load a music file from the given path, register it with the given name.
func RegisterMusic(name, path string) {
    m := rl.LoadMusicStream(path)
    m.Looping = false // disable looping by default, this causes issues with fading tracks and looping is done by our player anyway.
    a.music_tracks[name] = m
}

// Load a sound file from the given path, register it with the given name.
func RegisterSound(name, path string) {
    a.sfx_tracks[name] = rl.LoadSound(path)
}

// CONFIGURATION

// Mute the Audio
func Mute() {
    a.is_mute = true
}

// Unmute the Audio
func Unmute() {
    a.is_mute = true
}

// Toggle the Mute State
func ToggleMute() {
    a.is_mute = !a.is_mute
}

// Set the Global Volume to a value between 0.0 and 1.0
func SetGlobalVolume(new_volume float32) {
    new_volume = util.Clamp(new_volume, 0.0, 1.0)
    a.global_volume = new_volume
}

// Set the Music Volume to a value between 0.0 and 1.0
func SetMusicVolume(new_volume float32) {
    new_volume = util.Clamp(new_volume, 0.0, 1.0)
    a.music_volume = new_volume
}

// Set the SFX Volume to a value between 0.0 and 1.0
func SetSFXVolume(new_volume float32) {
    new_volume = util.Clamp(new_volume, 0.0, 1.0)
    a.sfx_volume = new_volume
}

// Set the Fade Time for Music Tracks in seconds
func SetMusicFade(fade_secs float32) {
    a.music_fade_secs = fade_secs
}

// PLAYBACK

// Play a sound that has been registered with "name"
func PlaySound(name string) {
    s := a.sfx_tracks[name]
    rl.SetSoundVolume(s, a.sfx_volume * a.global_volume)
    rl.PlaySound(s)
}

// PlaySound with extended parameters.
func PlaySoundEx(name string, volume, pitch, pan float32) {
    s := a.sfx_tracks[name]
    rl.SetSoundPitch(s, pitch)
    rl.SetSoundPan(s, pan)
    rl.SetSoundVolume(s, a.sfx_volume * a.global_volume * volume)

    rl.PlaySound(s)

    // reset sound properties
    rl.SetSoundVolume(s, a.sfx_volume * a.global_volume)
    rl.SetSoundPitch(s, 0)
    rl.SetSoundPan(s, 0)
}

// Instantly start playing a music track that has been registered with "name".
func PlayMusicNow(name string) {
    // stop the currently playing track
    rl.StopMusicStream(a.music_tracks[a.curr_playing_music_name])

    // set the correct volume
    rl.SetMusicVolume(a.music_tracks[name], a.music_volume * a.global_volume)
    a.curr_playing_music_name = name
    rl.PlayMusicStream(a.music_tracks[name])
}

// Start playing a music track that has been registered with "name" after
// fading out the currently playing track.
func PlayMusicNowFade(name string) {
    // TODO, set some 'fade now' flag and line up next track
    
}
