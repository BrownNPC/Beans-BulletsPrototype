package gfx

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var virtualScreen rl.RenderTexture2D

// Initialize virtual screen resolution
func InitVirtualScreen(virtualResolutionX, virtualResolutionY int) {
	virtualScreen = rl.LoadRenderTexture(int32(virtualResolutionX), int32(virtualResolutionY))
}
func BeginDrawingVirtualScreen() {
	if virtualScreen.Texture.ID == 0 {
		panic("Virtual screen not initialized")
	}
	rl.BeginTextureMode(virtualScreen)
}

func EndDrawingVirtualScreen() {
	target := virtualScreen
	rl.EndTextureMode()

	// virtual resolution
	sW, sH := float32(target.Texture.Width), float32(target.Texture.Height)
	// screen/window size
	rW, rH := float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())

	// scaling factor based how much bigger
	// or smaller the screen is
	// compared to virtual resolution
	scale := min(rW/sW, rH/sH)

	// Position of the virtual screen within the window
	offsetX := (rW - sW*scale) / 2
	offsetY := (rH - sH*scale) / 2

	src := rl.Rectangle{Width: sW, Height: -sH}
	dst := rl.Rectangle{
		X:      offsetX,
		Y:      offsetY,
		Width:  sW * scale,
		Height: sH * scale,
	}

	rl.DrawTexturePro(target.Texture, src, dst,
		rl.Vector2{}, // rotation origin
		0,            // rotation
		rl.White)
}
func GetVirtualMousePosition() rl.Vector2 {
	target := virtualScreen

	sW, sH := float32(target.Texture.Width), float32(target.Texture.Height)

	rW, rH := float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())

	scale := min(rW/sW, rH/sH)

	// Position of the virtual screen within the window
	offsetX := (rW - sW*scale) / 2
	offsetY := (rH - sH*scale) / 2

	mouse := rl.GetMousePosition()

	position := rl.Vector2{
		X: (mouse.X - offsetX) / scale,
		Y: (mouse.Y - offsetY) / scale,
	}
	return position.Clamp(
		rl.NewVector2(0, 0),
		rl.NewVector2(sW, sH),
	)
}

func GetVirtualScreenToWorld2D(position rl.Vector2, camera rl.Camera2D) rl.Vector2 {
	target := virtualScreen

	sW, sH := float32(target.Texture.Width), float32(target.Texture.Height)
	rW, rH := float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())

	scale := min(rW/sW, rH/sH)

	virtualCamera := camera
	virtualCamera.Offset = camera.Offset.Scale(scale)
	virtualCamera.Target = camera.Target
	virtualCamera.Zoom = camera.Zoom * scale

	return rl.GetScreenToWorld2D(position, virtualCamera)
}
