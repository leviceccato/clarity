package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(800, 450, "Clarity")
	r.SetTargetFPS(60)
	defer r.CloseWindow()

	state := newStateManager()
	mainWorld := newMainWorld()
	state.addWorld(mainWorld)
	state.setActive("main")

	for !r.WindowShouldClose() {
		state.run()
	}
}
