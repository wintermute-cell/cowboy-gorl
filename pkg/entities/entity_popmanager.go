package entities

import (
	"cowboy-gorl/pkg/logging"
	"cowboy-gorl/pkg/util"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This checks at compile time if the interface is implemented
var _ Entity = (*PopmanagerEntity)(nil)

//
//  Popmanager Entity
//
type PopmanagerEntity struct {
    Scene_ore_clusters *([]*OreClusterEntity)
    Inventory *InventoryEntity
    idlePops int32
    collectorPops int32
    crafterPops int32
    repairPops int32
    medicPops int32

    // this only exists becuase I cant add fractional values to the inventory 
    secondTimer float32
}

func (ent *PopmanagerEntity) Init() {
    // Initialization logic for the entity
    ent.idlePops = ent.Inventory.Pops
    ent.collectorPops = 0
    ent.crafterPops = 0
    ent.repairPops = 0
    ent.medicPops = 0
}

func (ent *PopmanagerEntity) Deinit() {
    // De-initialization logic for the entity
}

func (ent *PopmanagerEntity) Update() {
    if ent.Inventory.Pops < 0 {
        ent.Inventory.Pops = 0
    }
    ent.removeDeadPops()
    ent.drawPanel()
    ent.makePopsWork()
}

func (ent *PopmanagerEntity) makePopsWork() {
    // JOB: Collector
    ent.secondTimer += rl.GetFrameTime()
    logging.Info("Second Timer: %v", ent.secondTimer)
    if ent.secondTimer >= 1.0 {
        ent.secondTimer = 0.0
        for _, cluster := range *ent.Scene_ore_clusters {
            // NOTE: The fact that I have to redo the whole "what type of ore logic"
            // here makes it clear that this needs to be more though through.
            switch cluster.GetOreType() {
            case Coal:
                // NOTE: I wish i could add fractions here, but since it's all done
                // as int32, im shit outta luck. Need to differentiate fractional and
                // full values in the future.
                ent.Inventory.Coal_ore += 1 * ent.collectorPops
            case Iron:
                ent.Inventory.Iron_ore += 1 * ent.collectorPops
            }
        }
    }
}

func (ent *PopmanagerEntity) drawPanel() {
    panel_origin := rl.NewVector2(120, 20)
    rg.Panel(rl.NewRectangle(panel_origin.X, panel_origin.Y, 320, 220), "Population Management")

    btnsize := float32(16.0)
    yoffset := float32(40.0)

    // Update logic for the entity
    rg.Label(rl.NewRectangle(panel_origin.X+10, panel_origin.Y+yoffset, 200, 10),
        fmt.Sprintf("Pops: %v", ent.Inventory.Pops))

    rg.Label(rl.NewRectangle(panel_origin.X+100, panel_origin.Y+yoffset, 200, 10),
        fmt.Sprintf("Idle: %v", ent.idlePops))

    yoffset += 42
    rg.Label(rl.NewRectangle(panel_origin.X+10, yoffset, 200, 10), "-- Jobs --")
    yoffset += 20
    // Pop jobs...
    // JOB: Collectors
    if rg.Button(rl.NewRectangle(panel_origin.X+10, yoffset, btnsize, btnsize), "+") {
        ent.assignPops(&ent.collectorPops, 1)
    }
    if rg.Button(rl.NewRectangle(panel_origin.X+10+btnsize, yoffset, btnsize, btnsize), "-") {
        ent.assignPops(&ent.collectorPops, -1)
    }
    rg.Label(rl.NewRectangle(panel_origin.X+14+(2*btnsize), yoffset, 200, 10),
        fmt.Sprintf("Collectors: %v", ent.collectorPops))
    yoffset += 26

    // JOB: Crafters
    if rg.Button(rl.NewRectangle(panel_origin.X+10, yoffset, btnsize, btnsize), "+") {
        ent.assignPops(&ent.crafterPops, 1)
    }
    if rg.Button(rl.NewRectangle(panel_origin.X+10+btnsize, yoffset, btnsize, btnsize), "-") {
        ent.assignPops(&ent.crafterPops, -1)
    }
    rg.Label(rl.NewRectangle(panel_origin.X+14+(2*btnsize), yoffset, 200, 10),
        fmt.Sprintf("Crafters: %v", ent.crafterPops))
    yoffset += 26

    // JOB: Repair
    if rg.Button(rl.NewRectangle(panel_origin.X+10, yoffset, btnsize, btnsize), "+") {
        ent.assignPops(&ent.repairPops, 1)
    }
    if rg.Button(rl.NewRectangle(panel_origin.X+10+btnsize, yoffset, btnsize, btnsize), "-") {
        ent.assignPops(&ent.repairPops, -1)
    }
    rg.Label(rl.NewRectangle(panel_origin.X+14+(2*btnsize), yoffset, 200, 10),
        fmt.Sprintf("Repairs: %v", ent.repairPops))
    yoffset += 26

    // JOB: Medic
    if rg.Button(rl.NewRectangle(panel_origin.X+10, yoffset, btnsize, btnsize), "+") {
        ent.assignPops(&ent.medicPops, 1)
    }
    if rg.Button(rl.NewRectangle(panel_origin.X+10+btnsize, yoffset, btnsize, btnsize), "-") {
        ent.assignPops(&ent.medicPops, -1)
    }
    rg.Label(rl.NewRectangle(panel_origin.X+14+(2*btnsize), yoffset, 200, 10),
        fmt.Sprintf("Medics: %v", ent.medicPops))
    yoffset += 26
}

func (ent *PopmanagerEntity) removeDeadPops() {
    prevPops := (ent.idlePops + ent.collectorPops + ent.crafterPops + ent.repairPops + ent.medicPops)
    deficite := prevPops - ent.Inventory.Pops
    if deficite > 0 {
        if deficite <= ent.idlePops {
            ent.idlePops -= deficite
            deficite = 0
        } else {
            deficite -= ent.idlePops
            ent.idlePops = 0
            distribution := util.DistributeInteger(deficite, []int32{ent.collectorPops, ent.crafterPops, ent.repairPops, ent.medicPops})
            ent.collectorPops -= distribution[0]
            ent.crafterPops -= distribution[1]
            ent.repairPops -= distribution[2]
            ent.medicPops -= distribution[3]
        }
    }
}



func (ent *PopmanagerEntity) assignPops(job *int32, amount int32) {
    if amount > 0 {
        if ent.idlePops - amount >= 0 {
            ent.idlePops -= amount
            *job += amount
        }
    } else {
        if *job + amount >= 0 {
            *job += amount
            ent.idlePops -= amount
        }
    }
}
