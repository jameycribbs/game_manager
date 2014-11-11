package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	MOUSE_OUT  = iota
	MOUSE_OVER = iota
	CLICKED    = iota
)

type MenuButton struct {
	position     Vector2D
	velocity     Vector2D
	acceleration Vector2D
	width        int32
	height       int32
	textureID    string
	currentRow   int
	currentFrame int32
}

//*****************************************************************************
// NewMenuButton
//*****************************************************************************
func NewMenuButton(x int32, y int32, width int32, height int32, textureID string) *MenuButton {
	return &MenuButton{position: Vector2D{x, y}, velocity: Vector2D{0, 0}, width: width, height: height, textureID: textureID,
		currentRow: 1, currentFrame: 1}
}

//*****************************************************************************
// draw
//*****************************************************************************
func (mb *MenuButton) draw(game *Game) {
	game.textureManager.drawFrame(mb.textureID, mb.position.X, mb.position.Y, mb.width, mb.height, mb.currentRow,
		mb.currentFrame, game.renderer, sdl.FLIP_NONE)
}

//*****************************************************************************
// update
//*****************************************************************************
func (mb *MenuButton) update(game *Game) {
	mousePos := game.inputHandler.getMousePosition()

	if mousePos.X < (mb.position.X+mb.width) && mousePos.X > mb.position.X && mousePos.Y < (mb.position.Y+mb.height) && mousePos.Y > mb.position.Y {
		mb.currentFrame = MOUSE_OVER

		if game.inputHandler.getMouseButtonState(LEFT) {
			mb.currentFrame = CLICKED
		}
	} else {
		mb.currentFrame = MOUSE_OUT
	}
}

//*****************************************************************************
// clean
//*****************************************************************************
func (mb *MenuButton) clean() {
	fmt.Println("clean menuButton")
}
