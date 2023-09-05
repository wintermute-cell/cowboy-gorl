package entities

import (
	"cowboy-gorl/pkg/logging"
	"math/rand"
    "fmt"
    "strings"
    "sort"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*EventmanagerEntity)(nil)

//
//  Eventmanager Entity
//
type EventmanagerEntity struct {
    // Add fields here for any state that the entity should keep track of
    Inventory *InventoryEntity
    show_event bool
    event_timer float32
    decision_timer float32
    decision_expired bool
    current_decision_pair []Decision
}

type Decision struct {
	Text          string
	ResourceChanges map[string]int32
}

func (ent *EventmanagerEntity) Init() {
    //
}

func (ent *EventmanagerEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *EventmanagerEntity) Update() {
    if !ent.show_event {
        ent.event_timer -= rl.GetFrameTime()
    }

    if ent.event_timer <= 0 {
		// Show event and reset timer
		ent.show_event = true
        ent.current_decision_pair = getRandomDecisionPair()
		ent.decision_timer = 8
		ent.event_timer = float32(rand.Intn(8) + 40)
        ent.decision_expired = false
	}

    if ent.show_event {
		// Decrease the decision timer
		ent.decision_timer -= rl.GetFrameTime()

		// Randomly make a decision if time runs out
		if ent.decision_timer <= 0 && !ent.decision_expired {
            randomChoice := ent.current_decision_pair[rand.Intn(
                len(ent.current_decision_pair),
                )]
            // Create a new map to hold the penalized resource changes
            penalizedResourceChanges := make(map[string]int32)

            // Loop through the original resource changes and apply the penalty
            for resource, change := range randomChoice.ResourceChanges {
                penalizedResourceChanges[resource] = int32(float32(change) * 1.5)
            }

            // Construct a new Decision struct with the penalized resource changes
            penalized_decision := Decision{
                Text:            randomChoice.Text,
                ResourceChanges: penalizedResourceChanges,
            }
			changeResource(ent.Inventory, penalized_decision)
			ent.show_event = false
			ent.decision_expired = true
		}

        // Hard-coded positions for demonstration
		firstPanelX, firstPanelY, panelWidth, panelHeight := 450, 80, 160, 80
		secondPanelX, secondPanelY := firstPanelX, firstPanelY+130

        // Handle first decision
        if rg.Button(rl.NewRectangle(float32(firstPanelX), float32(firstPanelY), float32(panelWidth), float32(panelHeight)), "") {
            changeResource(ent.Inventory, ent.current_decision_pair[0])
            ent.show_event = false
        }
        text1 := ent.current_decision_pair[0].Text + "\n(" + formatResourceChanges(ent.current_decision_pair[0].ResourceChanges) + ")"
        rl.DrawText(text1, int32(firstPanelX+10), int32(firstPanelY+10), 10, rl.Black)

        // Handle second decision
        if rg.Button(rl.NewRectangle(float32(secondPanelX), float32(secondPanelY), float32(panelWidth), float32(panelHeight)), "") {
            changeResource(ent.Inventory, ent.current_decision_pair[1])
            ent.show_event = false
        }
        text2 := ent.current_decision_pair[1].Text + "\n(" + formatResourceChanges(ent.current_decision_pair[1].ResourceChanges) + ")"
        rl.DrawText(text2, int32(secondPanelX+10), int32(secondPanelY+10), 10, rl.Black)

	}
}

func getRandomDecisionPair() []Decision {
	decisionPairs := [][]Decision{
		{
			{"Repair Tracks", map[string]int32{"Iron_ore": -20, "Coal_ore": -10}},
			{"Speed Through!", map[string]int32{"Pops": -10}},
		},
		{
			{"Scavenge for Supplies", map[string]int32{"Pops": -5, "Food": 20}},
			{"Keep Moving", map[string]int32{"Coal_ore": 10, "Food": -20}},
		},
		{
			{"Fight Off Bandits", map[string]int32{"Dynamite": -20, "Pops": 10}},
			{"Pay Tribute", map[string]int32{"Food": -30, "Coal_ore": -30}},
		},
		// Add more decisions here
	}

	return decisionPairs[rand.Intn(len(decisionPairs))]
}

// Utility function to format resource changes as a string
func formatResourceChanges(changes map[string]int32) string {
	var resourceChangeTexts []string
	keys := make([]string, 0, len(changes))

	for k := range changes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, resource := range keys {
		change := changes[resource]
		sign := ""
		if change > 0 {
			sign = "+"
		}
		resourceChangeTexts = append(resourceChangeTexts, fmt.Sprintf("%s%s %d", resource, sign, change))
	}

	return strings.Join(resourceChangeTexts, ", ")
}

func changeResource(inventory *InventoryEntity, decision Decision) {
    var log_entries []string

    // NOTE: This has to be way more dynamic. Having to add an entry for every
    // type of resource here is dumb
	for resource, change := range decision.ResourceChanges {
		switch resource {
		case "Pops":
			inventory.Pops += change
		case "Coal_ore":
			inventory.Coal_ore += change
		case "Iron_ore":
			inventory.Iron_ore += change
		case "Dynamite":
			inventory.Dynamite += change
		case "Food":
			inventory.Food += change
		}
        log_entry := fmt.Sprintf("%s: %+d", resource, change)
		log_entries = append(log_entries, log_entry)
	}
    logging.Info("Resource change: %s", strings.Join(log_entries, ", "))
}
