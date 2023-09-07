// SceneManager provides a manager for game scenes, automating the calling
// of their Init(), Deinit(), Draw(), ... functions, 
// A SceneManager, also features enabling/disabling, and ordering of scenes
// for drawing operations.
// 
// Usage:
//    - Create a new SceneManager with `NewSceneManager`.
//    - Register scenes using `RegisterScene(name, scene)`.
//    - Control scene state with `EnableScene` and `DisableScene`.
//    - Modify draw order using `MoveSceneToFront`, `MoveSceneToBack`, and `MoveSceneBefore`.
//    - In the game loop, use `DrawScenes` and `DrawScenesGUI` to render scenes in their specified order.

package scenes

import "cowboy-gorl/pkg/logging"

type SceneManager struct {
	scenes        map[string]Scene
	enable_scenes map[string]bool
	scene_order   []string // slice to maintain order, since map is unordered
}

// Create a new SceneManager. A SceneManager will automatically take care of
// your Scenes (calling their Init(), Deinit(), Draw(), DrawGUI() functions).
func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes:        make(map[string]Scene),
		enable_scenes: make(map[string]bool),
		scene_order:   make([]string, 0),
	}
}

// Register a scene with the SceneManager for automatic control
func (sm *SceneManager) RegisterScene(name string, scene Scene) {
	if _, exists := sm.scenes[name]; exists {
		logging.Fatal("A scene with this name is already registered.")
	}
	sm.scenes[name] = scene
	sm.scene_order = append(sm.scene_order, name) // Add to the end by default
}

// MoveSceneToFront moves the scene to the front of the draw order
func (sm *SceneManager) MoveSceneToFront(name string) {
	sm.reorderScene(name, 0)
}

// MoveSceneToBack moves the scene to the end of the draw order
func (sm *SceneManager) MoveSceneToBack(name string) {
	sm.reorderScene(name, len(sm.scene_order)-1)
}

// MoveSceneBefore moves the scene right before another scene in the draw order
func (sm *SceneManager) MoveSceneBefore(sceneName, beforeSceneName string) {
	index, found := sm.getSceneOrderIndex(beforeSceneName)
	if found {
		sm.reorderScene(sceneName, index)
	}
}

func (sm *SceneManager) reorderScene(name string, index int) {
	current_idx, found := sm.getSceneOrderIndex(name)
	if !found {
		return
	}
	sm.scene_order = append(sm.scene_order[:current_idx], sm.scene_order[current_idx+1:]...)
	sm.scene_order = append(sm.scene_order[:index], append([]string{name}, sm.scene_order[index:]...)...)
}

func (sm *SceneManager) getSceneOrderIndex(name string) (int, bool) {
	for i, scene_name := range sm.scene_order {
		if scene_name == name {
			return i, true
		}
	}
	return -1, false
}

// Enable the Scene. The Scenes Init() function will be called.
func (sm *SceneManager) EnableScene(name string) {
	scene, exists := sm.scenes[name]
	if !exists {
		logging.Fatal("Scene not found.")
	}

	// Initialize the scene if it's not already enabled
	if !sm.enable_scenes[name] {
		scene.Init()
		sm.enable_scenes[name] = true
	}
}

// Disable the Scene. The Scenes Deinit() function will be called.
func (sm *SceneManager) DisableScene(name string) {
	scene, exists := sm.scenes[name]
	if !exists {
		logging.Fatal("Scene not found.")
	}

	// De-initialize the scene if it's currently enabled
	if sm.enable_scenes[name] {
		scene.Deinit()
		sm.enable_scenes[name] = false
	}
}

// Disable all Scenes that are currently enabled.
func (sm *SceneManager) DisableAllScenes() {
	for name, scene := range sm.scenes {
		if sm.enable_scenes[name] {
            scene.Deinit()
            sm.enable_scenes[name] = false
		}
	}
}

// Call the Draw() functions of all the registered Scenes in their defined order.
func (sm *SceneManager) DrawScenes() {
	for name, scene := range sm.scenes {
		if sm.enable_scenes[name] {
			scene.Draw()
		}
	}
}

// Call the DrawGUI() functions of all the registered Scenes in their defined order.
func (sm *SceneManager) DrawScenesGUI() {
	for name, scene := range sm.scenes {
		if sm.enable_scenes[name] {
			scene.DrawGUI()
		}
	}
}
