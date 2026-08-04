package gfx

import (
	"fmt"
	"image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AssetID int

const (
	Texture_Prototype_Grid AssetID = iota
)

var loadedTextures = map[AssetID]rl.Texture2D{}
var texturePath = map[AssetID]string{
	Texture_Prototype_Grid: "texture_13.png",
}

// Texture pack
func GetTexture(assetID AssetID) rl.Texture2D {
	if tex, ok := loadedTextures[assetID]; ok {
		return tex
	}
	if texPath, ok := texturePath[assetID]; ok {
		f, err := assetsFS.Open(texPath)
		if err != nil {
			panic(err)
		}
		img, err := png.Decode(f)
		if err != nil {
			panic(err)
		}
		tex := rl.LoadTextureFromImage(rl.NewImageFromImage(img))
		loadedTextures[assetID] = tex
		return tex
	}
	panic(fmt.Errorf("No texture path assigned for id %d", assetID))
}
