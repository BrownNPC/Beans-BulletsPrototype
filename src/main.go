//go:build !cgo

package main

import (
	"game/gfx"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenX, screenY = 640 * 2, 360 * 2
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint | rl.FlagWindowHighdpi)
	rl.InitWindow(0, 0, "Prototype")
	gfx.InitVirtualScreen(screenX, screenY)

	var state ScreenIngameState
	state.Init()
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		{
			state.Update()
		}
		rl.EndDrawing()
	}
}
