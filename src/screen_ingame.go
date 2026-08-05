package main

import (
	"game/gfx"
	t "game/things"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var things t.Things

type ScreenIngameState struct {
	Cam rl.Camera2D
	Plr t.ThingRef

	Input struct {
		CurrentMouse  rl.Vector2
		MouseVelocity rl.Vector2
		TargetMouse   rl.Vector2
		MouseDelta    rl.Vector2
		MouseSpring   gfx.Spring

		Movement rl.Vector2
		accum    float32
	}

	Targets    [2]t.ThingRef
	ViewRadius float32

	initTime     time.Time
	ElapsedTicks uint
}

func (state *ScreenIngameState) Init() {
	state.Cam = rl.Camera2D{
		Target:   rl.Vector2{},
		Rotation: 0,
		Zoom:     .8,
	}
	state.Plr = things.Add(&t.Player{
		Position: rl.Vector2{},
		LookAt:   rl.Vector2{Y: 1},
		Size:     50,
	})
	state.Input.MouseSpring = gfx.NewSpring(1, .5, .2)
	state.ViewRadius = 512
	rl.DisableCursor()
}
func (state *ScreenIngameState) Update() {
	state.ProcessInput()
	state.ElapsedTicks = uint(time.Since(state.initTime).Seconds())

	plr := things.Get(state.Plr).(*t.Player)
	*plr = plr.Tick(rl.GetFrameTime(), t.PlayerTickConfig{
		Cursor:     state.Input.CurrentMouse,
		MoveVector: state.Input.Movement,
	})
	state.Cam.Zoom = (float32(min(rl.GetRenderHeight(), rl.GetRenderWidth())) / 2) / state.ViewRadius
	state.Render()
}

func (state *ScreenIngameState) Render() {
	rl.ClearBackground(rl.RayWhite)
	plr := things.Get(state.Plr).(*t.Player)
	rl.BeginMode2D(state.Cam)
	// draw map floor
	gfx.DrawTextureTiled(gfx.GetTexture(gfx.Texture_Prototype_Grid),
		rl.NewRectangle(-1024/2, -1024/2, 1024*6, 1024*6),
		1, rl.White,
	)
	rl.EndMode2D()
	rl.BeginMode2D(state.Cam)

	rl.DrawCircleV(plr.Position, plr.Size/2, rl.Red)
	rl.DrawRectanglePro(
		rl.NewRectangle(plr.Position.X, plr.Position.Y, 5, 100),
		rl.NewVector2(5.0/2, 0),
		plr.Rotation*360,
		rl.Red,
	)
	rl.DrawCircleV(plr.Position, 30/2, rl.Yellow)
	rl.DrawCircleV(state.Input.CurrentMouse, 10, rl.Black)
	state.Cam.Target = plr.Position
	state.Cam.Offset = rl.NewVector2(float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())).Scale(.5)
	rl.EndMode2D()
}
