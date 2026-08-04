package t

import (
	"game/gfx"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	cp "github.com/jakecoffman/cp/v2"
)

const (
	Player_StanceNeutral int = iota
	Player_StanceSliding

	Player_SlideDuration time.Duration = time.Second
)

type Player struct {
	Stance int

	StartedSliding float32

	Position rl.Vector2
	Velocity rl.Vector2

	// Point in world coordinates where the player is aiming.
	// (Mouse position)
	LookAt   rl.Vector2
	Rotation float32 // player rotation in turns

	Size float32

	Shape cp.Shape
}

// PlayerTickConfig is passed to Player.Tick
type PlayerTickConfig struct {
	// Position of the reticle on screen.
	Cursor rl.Vector2
	// Normalized movement vector.
	MoveVector rl.Vector2
}

func (plr Player) Tick(dt float32, cfg PlayerTickConfig) Player {
	plr = plr.ApplyMovement(dt, cfg.MoveVector, cfg.Cursor)
	return plr
}

func (plr Player) StartSliding(tick uint) Player {
	
	return plr
}
func (plr Player) ApplyMovement(dt float32, move rl.Vector2, cursor rl.Vector2) Player {
	if dt > 1.0/30.0 { // stop player from teleporting through walls
		dt = 1.0 / 30.0
	}
	const (
		maxSpeed = 256.0 * 1.5
		// how long it takes to reach max speed.
		// acceleration = maxSpeed * number of seconds to reach max speed.
		acceleration = maxSpeed * 10
		friction     = 14
	)

	// Pre-calculate fixed decay for the loop
	movedThisFrame := move.LengthSqr() > 0
	if movedThisFrame {
		// Apply acceleration * time
		plr.Velocity = plr.Velocity.Add(move.Scale(acceleration * dt))
	} else {
		decay := float32(math.Exp(-friction * float64(dt)))
		plr.Velocity = plr.Velocity.Scale(decay)
	}
	// Clamp velocity to maxSpeed
	if plr.Velocity.Length() > maxSpeed {
		plr.Velocity = plr.Velocity.Normalize().Scale(maxSpeed)
	}
	// Apply velocity * time to position
	plr.Position = plr.Position.Add(plr.Velocity.Scale(dt))

	plr.LookAt = cursor.Subtract(plr.Position).Normalize()
	plr.Rotation = rl.Vector2{Y: 1}.Angle(plr.LookAt) / gfx.Tau

	return plr
}
