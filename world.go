package main

type world struct {
	name     string
	systems  []runner
	entities []*entity
}

func (w *world) run() {
	for _, system := range w.systems {
		system.run()
	}
}

func (w *world) addEntity(e *entity) {
	w.entities = append(w.entities, e)
}

func (w *world) addSystem(s runner) {
	w.systems = append(w.systems, s)
}

// Update all systems with the entities they will be run on
func (w *world) updateSystems() {
	hasAllComponents := true
	hasComponent := false
	for _, system := range w.systems {
		for _, entity := range w.entities {
			for _, component := range system.getComponents() {
				switch component {
				case "appearance":
					if entity.appearance != nil {
						hasComponent = true
					}
				case "collision":
					if entity.collision != nil {
						hasComponent = true
					}
				case "physics":
					if entity.physics != nil {
						hasComponent = true
					}
				case "position":
					if entity.position != nil {
						hasComponent = true
					}
				}
			}
			hasAllComponents = hasComponent
			if hasAllComponents {
				system.addEntity(entity)
			}
			hasComponent = false
			hasAllComponents = true
		}
	}
}

func newMainWorld() *world {
	w := &world{name: "main"}
	w.addSystem(newDrawSystem())
	w.addSystem(newPlayerSytem())
	w.addEntity(newPlayerEntity())
	w.updateSystems()
	return w
}
