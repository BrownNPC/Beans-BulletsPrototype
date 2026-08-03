//go:generate go run github.com/BrownNPC/things.go/cmd@latest
package t

import rl "github.com/gen2brain/raylib-go/raylib"

import cp "github.com/jakecoffman/cp/v2"

type Player struct {
	Position rl.Vector2

	// Point in world coordinates where the player is aiming.
	// (Mouse position)
	LookAt rl.Vector2

	Size float32

	Shape cp.Shape
}

// add other kinds of structs like a Zombie struct:
