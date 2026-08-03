package main

import (
	t "game/things"

	"github.com/charmbracelet/harmonica"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// import "github.com/charmbracelet/harmonica"

var things t.Things

type ScreenIngameState struct {
	Cam rl.Camera2D
	Plr t.ThingRef

	CurrentMouse  rl.Vector2
	MouseVelocity rl.Vector2

	accum       float32
	MouseSpring harmonica.Spring
}

func (state *ScreenIngameState) Init() {
	state.Cam = rl.Camera2D{
		Target:   rl.Vector2{},
		Rotation: 0,
		Zoom:     1,
	}
	state.Plr = things.Add(&t.Player{
		Position: rl.Vector2{},
		LookAt:   rl.Vector2{},
		Size:     50,
	})
	state.CurrentMouse = rl.GetScreenToWorld2D(rl.GetMousePosition(), state.Cam)
	state.MouseSpring = harmonica.NewSpring(1.0/60, 20, .67)
}

func (state *ScreenIngameState) Update() {
	plr := things.Get(state.Plr).(*t.Player)
	dt := rl.GetFrameTime()
	state.accum += dt
	const tickrate = 60
	for state.accum >= 1.0/tickrate {
		state.accum -= 1.0 / tickrate
		target := rl.GetScreenToWorld2D(rl.GetMousePosition(), state.Cam)
		x, vx := state.MouseSpring.Update(float64(state.CurrentMouse.X), float64(state.MouseVelocity.X), float64(target.X))
		y, vy := state.MouseSpring.Update(float64(state.CurrentMouse.Y), float64(state.MouseVelocity.Y), float64(target.Y))
		state.MouseVelocity = rl.NewVector2(float32(vx), float32(vy))
		state.CurrentMouse = rl.NewVector2(float32(x), float32(y))
	}

	plr.LookAt = state.CurrentMouse

	state.Render()
}

func (state *ScreenIngameState) Render() {
	rl.ClearBackground(rl.RayWhite)
	rl.BeginMode2D(state.Cam)
	plr := things.Get(state.Plr).(*t.Player)
	rl.DrawCircleV(plr.Position, plr.Size/2, rl.Red)
	rl.DrawCircleV(plr.LookAt, 5, rl.Black)
	rl.EndMode2D()
	state.Cam.Target = plr.Position
	state.Cam.Offset = rl.NewVector2(float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())).Scale(.5)
}
