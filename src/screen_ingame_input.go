package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (state *ScreenIngameState) ProcessInput() {
	input := &state.Input
	input.TargetMouse = input.TargetMouse.Add(rl.GetMouseDelta())
	input.TargetMouse = input.TargetMouse.Clamp(
		rl.GetScreenToWorld2D(rl.NewVector2(0, 0), state.Cam),
		rl.GetScreenToWorld2D(rl.NewVector2(float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())), state.Cam),
	)

	dt := rl.GetFrameTime()
	input.accum += dt
	const tickrate = 60
	var mouseDelta rl.Vector2
	for input.accum >= 1.0/tickrate {
		input.accum -= 1.0 / tickrate
		target := input.TargetMouse
		x, vx := state.Input.MouseSpring.Update(state.Input.CurrentMouse.X, input.MouseVelocity.X, target.X)
		y, vy := state.Input.MouseSpring.Update(state.Input.CurrentMouse.Y, input.MouseVelocity.Y, target.Y)
		input.MouseVelocity = rl.NewVector2(float32(vx), float32(vy))
		newPos := rl.NewVector2(float32(x), float32(y))
		mouseDelta = newPos.Subtract(input.CurrentMouse)
		input.CurrentMouse = newPos
	}
	input.MouseDelta = mouseDelta
	move := rl.Vector2{}
	if rl.IsKeyDown(rl.KeyW) {
		move.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		move.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		move.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		move.X += 1
	}
	input.Movement = move.Normalize()
}
