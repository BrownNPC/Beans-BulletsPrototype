package spawn

import (
	t "game/things"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlayer(things *t.Things) t.ThingRef {
	plr := t.Player{
		Position: rl.Vector2{},
		Size:     20,
	}
	
	return things.Add(&plr)
}
