package main

type stateManager struct {
	worlds      map[string]*world
	activeWorld *world
}

func newStateManager() *stateManager {
	sm := &stateManager{}
	sm.worlds = make(map[string]*world)
	return sm
}

func (sm *stateManager) addWorld(w *world) {
	sm.worlds[w.name] = w
}

func (sm *stateManager) setActive(name string) {
	sm.activeWorld = sm.worlds[name]
}

func (sm *stateManager) run() {
	sm.activeWorld.run()
}
