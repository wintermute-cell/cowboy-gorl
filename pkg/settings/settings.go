package settings

import (
	"encoding/json"
	"os"
)

type GameSettings struct {
	// Display
	ScreenWidth  int  `json:"screenWidth"`  // 1920
	ScreenHeight int  `json:"screenHeight"` // 1080
	RenderWidth  int  `json:"renderWidth"`  // 640
	RenderHeight int  `json:"renderHeight"` // 480
	TargetFps    int  `json:"targetFps"`    // 144
	Fullscreen   bool `json:"fullscreen"`   // false
    EnableCrt    bool `json:"enableCrt"`    // false
	// Gameplay
	MouseSensitivity float32 `json:"mouseSensitivity"` // 1.0
	// Audio
	SoundVolume float32 `json:"soundVolume"` // 0.5
	// Logging
	LogPath string `json:"logPath"`			 // logs/

}

var (
	settings *GameSettings
)

// Get the current settings
func CurrentSettings() *GameSettings {
	return settings
}

func FallbackSettings() {
	settings = &GameSettings{
		ScreenWidth:  1920,
		ScreenHeight:  1080,
		RenderWidth:  640,
		RenderHeight:  480,
		TargetFps:  144,
		Fullscreen:  false,
		EnableCrt:  false,
		MouseSensitivity:  1.0,
		SoundVolume:  0.5,
		LogPath:  "logs/",
	}
}

func LoadSettings(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	settings = new(GameSettings)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(settings)
	if err != nil {
		return err
	}

	return nil
}
