package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os"
)

type TextureManager struct {
	textureMap map[string]*sdl.Texture
}

//*****************************************************************************
// NewTextureManager
//*****************************************************************************
func NewTextureManager() *TextureManager {
	return &TextureManager{textureMap: make(map[string]*sdl.Texture)}
}

//*****************************************************************************
// load
//*****************************************************************************
func (tm *TextureManager) load(fileName string, id string, renderer *sdl.Renderer) bool {
	var tempSurface *sdl.Surface
	var texture *sdl.Texture

	tempSurface = img.Load(fileName)

	if tempSurface == nil {
		fmt.Fprintf(os.Stderr, "Failed to load image: %s", sdl.GetError())
		return false
	} else {
		fmt.Println("image load success")
	}

	texture = renderer.CreateTextureFromSurface(tempSurface)

	tempSurface.Free()

	if texture == nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", sdl.GetError())
		return false
	} else {
		fmt.Println("texture creation success")
	}

	tm.textureMap[id] = texture

	return true
}

//*****************************************************************************
// draw
//*****************************************************************************
func (tm *TextureManager) draw(id string, x int32, y int32, width int32, height int32, renderer *sdl.Renderer,
	flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect
	var centerPoint sdl.Point

	centerPoint.X = 0
	centerPoint.Y = 0

	srcRect.X = 0
	srcRect.Y = 0

	srcRect.W = width
	destRect.W = width

	srcRect.H = height
	destRect.H = height

	destRect.X = x
	destRect.Y = y

	renderer.CopyEx(tm.textureMap[id], &srcRect, &destRect, 0, &centerPoint, flip)
}

//*****************************************************************************
// drawFrame
//*****************************************************************************
func (tm *TextureManager) drawFrame(id string, x int32, y int32, width int32, height int32, currentRow int, currentFrame int32,
	renderer *sdl.Renderer, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect
	var centerPoint sdl.Point

	centerPoint.X = 0
	centerPoint.Y = 0

	srcRect.X = width * currentFrame
	srcRect.Y = height * int32(currentRow-1)

	srcRect.W = width
	destRect.W = width

	srcRect.H = height
	destRect.H = height

	destRect.X = x
	destRect.Y = y

	renderer.CopyEx(tm.textureMap[id], &srcRect, &destRect, 0, &centerPoint, flip)
}

//*****************************************************************************
// clearFromTextureMap
//*****************************************************************************
func (tm *TextureManager) clearFromTextureMap(id string) {
	delete(tm.textureMap, id)
}
