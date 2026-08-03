package gfx

import rl "github.com/gen2brain/raylib-go/raylib"

// Helpers over raylib functions.

func DrawTextureTiled(
	texture rl.Texture2D,
	dest rl.Rectangle,
	scale float32,
	tint rl.Color,
) {
	if texture.ID == 0 {
		return
	}

	if scale <= 0 {
		scale = 1
	}

	tileW := float32(texture.Width) * scale
	tileH := float32(texture.Height) * scale

	// UVs larger than 1.0 cause GL_REPEAT wrapping
	u := dest.Width / tileW
	v := dest.Height / tileH

	rl.SetTexture(texture.ID)

	rl.Begin(rl.Quads)

	rl.Color4ub(tint.R, tint.G, tint.B, tint.A)
	rl.Normal3f(0, 0, 1)

	// Top-left
	rl.TexCoord2f(0, 0)
	rl.Vertex2f(dest.X, dest.Y)

	rl.TexCoord2f(0, v)
	rl.Vertex2f(dest.X, dest.Y+dest.Height)

	// Bottom-right
	rl.TexCoord2f(u, v)
	rl.Vertex2f(dest.X+dest.Width, dest.Y+dest.Height)

	// Top-right
	rl.TexCoord2f(u, 0)
	rl.Vertex2f(dest.X+dest.Width, dest.Y)

	rl.End()
	rl.SetTexture(0)
}
