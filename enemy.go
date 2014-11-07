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
func NewEnemy(x int32, y int32, width int32, height int32, textureID string) *Enemy {
	return &Enemy{position: Vector2D{x, y}, velocity: Vector2D{0, 0}, width: width, height: height, textureID: textureID,
		currentRow: 1, currentFrame: 1}
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

	enemy.velocity = enemy.velocity.Add(&enemy.acceleration)
	enemy.position = enemy.position.Add(&enemy.velocity)
}

//*****************************************************************************
// clean
//*****************************************************************************
func (enemy *Enemy) clean() {
	fmt.Println("clean enemy")
}
