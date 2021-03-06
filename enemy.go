package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Enemy struct {
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
// NewEnemy
//*****************************************************************************
func NewEnemy(x float32, y float32, width int32, height int32, textureID string) *Enemy {
	return &Enemy{position: Vector2D{x, y}, velocity: Vector2D{0.001, 2}, width: width, height: height, textureID: textureID,
		currentRow: 1, currentFrame: 1}
}

//*****************************************************************************
// Position
//*****************************************************************************
func (enemy *Enemy) Position() Vector2D {
	return enemy.position
}

//*****************************************************************************
// Width
//*****************************************************************************
func (enemy *Enemy) Width() int32 {
	return enemy.width
}

//*****************************************************************************
// Height
//*****************************************************************************
func (enemy *Enemy) Height() int32 {
	return enemy.height
}

//*****************************************************************************
// draw
//*****************************************************************************
func (enemy *Enemy) draw(game *Game) {
	game.textureManager.drawFrame(enemy.textureID, enemy.position.X, enemy.position.Y, enemy.width, enemy.height, enemy.currentRow,
		enemy.currentFrame, game.renderer, sdl.FLIP_NONE)
}

//*****************************************************************************
// update
//*****************************************************************************
func (enemy *Enemy) update(game *Game) {
	enemy.currentFrame = int32(((sdl.GetTicks() / 100) % 6))

	if enemy.position.Y < 0 {
		enemy.velocity.Y = 2
	} else if enemy.position.Y > 400 {
		enemy.velocity.Y = -2
	}

	//	enemy.velocity = enemy.velocity.Add(&enemy.acceleration)
	enemy.position = enemy.position.Add(&enemy.velocity)
}

//*****************************************************************************
// clean
//*****************************************************************************
func (enemy *Enemy) clean() {
	fmt.Println("clean enemy")
}
