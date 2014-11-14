package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type AnimatedGraphic struct {
	position     Vector2D
	velocity     Vector2D
	acceleration Vector2D
	width        int32
	height       int32
	textureID    string
	currentRow   int
	currentFrame int32
	bReleased    bool
	animSpeed    int
}

//*****************************************************************************
// NewAnimatedGraphic
//*****************************************************************************
func NewAnimatedGraphic(x float32, y float32, width int32, height int32, textureID string, animSpeed int) *AnimatedGraphic {
	return &AnimatedGraphic{position: Vector2D{x, y}, velocity: Vector2D{0, 0}, width: width, height: height, textureID: textureID,
		currentRow: 1, currentFrame: 1, animSpeed: animSpeed}
}

//*****************************************************************************
// Position
//*****************************************************************************
func (ag *AnimatedGraphic) Position() Vector2D {
	return ag.position
}

//*****************************************************************************
// Width
//*****************************************************************************
func (ag *AnimatedGraphic) Width() int32 {
	return ag.width
}

//*****************************************************************************
// Height
//*****************************************************************************
func (ag *AnimatedGraphic) Height() int32 {
	return ag.height
}

//*****************************************************************************
// draw
//*****************************************************************************
func (ag *AnimatedGraphic) draw(game *Game) {
	game.textureManager.drawFrame(ag.textureID, ag.position.X, ag.position.Y, ag.width, ag.height, ag.currentRow,
		ag.currentFrame, game.renderer, sdl.FLIP_NONE)
}

//*****************************************************************************
// update
//*****************************************************************************
func (ag *AnimatedGraphic) update(game *Game) {
	mousePos := game.inputHandler.getMousePosition()

	if mousePos.X < (ag.position.X+float32(ag.width)) &&
		mousePos.X > ag.position.X &&
		mousePos.Y < (ag.position.Y+float32(ag.height)) &&
		mousePos.Y > ag.position.Y {

		if game.inputHandler.getMouseButtonState(LEFT) && ag.bReleased {
			ag.currentFrame = CLICKED
			ag.bReleased = false
		} else if !game.inputHandler.getMouseButtonState(LEFT) {
			ag.bReleased = true
			ag.currentFrame = MOUSE_OVER
		}
	} else {
		ag.currentFrame = MOUSE_OUT
	}
}

//*****************************************************************************
// clean
//*****************************************************************************
func (ag *AnimatedGraphic) clean() {
	fmt.Println("clean animatedGraphic")
}
