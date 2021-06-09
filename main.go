package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(800, 450, "Clarity")
	r.SetWindowIcon(*r.LoadImage("sprites/window_icon.png"))
	r.SetTargetFPS(60)
	defer r.CloseWindow()

	state := newStateManager()
	mainWorld := newMainWorld()
	state.addWorld(mainWorld)
	state.activateWorld("main")

	for !r.WindowShouldClose() {
		r.BeginDrawing()
		state.run()
		r.EndDrawing()
	}
}
