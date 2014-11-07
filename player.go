package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
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
// NewPlayer
//*****************************************************************************
func NewPlayer(x int32, y int32, width int32, height int32, textureID string) *Player {
	return &Player{position: Vector2D{x, y}, velocity: Vector2D{0, 0}, width: width, height: height, textureID: textureID,
		currentRow: 1, currentFrame: 1}
}

//*****************************************************************************
// draw
//*****************************************************************************
func (player *Player) draw(game *Game) {
	game.textureManager.drawFrame(player.textureID, player.position.X, player.position.Y, player.width, player.height, player.currentRow,
		player.currentFrame, game.renderer, sdl.FLIP_NONE)
}

//*****************************************************************************
// update
//*****************************************************************************
func (player *Player) update(game *Game) {
	player.currentFrame = int32(((sdl.GetTicks() / 100) % 6))

	vec := game.inputHandler.getMousePosition()
	vec = vec.Subtract(&player.position)
	vec = vec.Divide(100)

	player.velocity = vec

	player.velocity = player.velocity.Add(&player.acceleration)
	player.position = player.position.Add(&player.velocity)
}

//*****************************************************************************
// clean
//*****************************************************************************
func (player *Player) clean() {
	fmt.Println("clean player")
}
