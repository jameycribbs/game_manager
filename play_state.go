package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type PlayState struct {
	playID      string
	gameObjects []GameObject
}

//*****************************************************************************
// NewPlayState
//*****************************************************************************
func NewPlayState() *PlayState {
	return &PlayState{playID: "PLAY"}
}

//*****************************************************************************
// getStateID
//*****************************************************************************
func (ps *PlayState) getStateID() string {
	return ps.playID
}

//*****************************************************************************
// update
//*****************************************************************************
func (ps *PlayState) update(game *Game) {
	if game.inputHandler.isKeyDown(sdl.SCANCODE_ESCAPE) {
		game.gameStateMachine.pushState(NewPauseState())
	}

	for _, gameObject := range ps.gameObjects {
		gameObject.update(game)
	}

	if ps.checkCollision(ps.gameObjects[0], ps.gameObjects[1]) {
		game.gameStateMachine.pushState(NewGameOverState())
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (ps *PlayState) render(game *Game) {
	for _, gameObject := range ps.gameObjects {
		gameObject.draw(game)
	}
}

//*****************************************************************************
// resume
//*****************************************************************************
func (ps *PlayState) resume(game *Game) {
}

//*****************************************************************************
// onEnter
//*****************************************************************************
func (ps *PlayState) onEnter(game *Game) bool {
	if !game.textureManager.load("assets/helicopter.png", "helicopter", game.renderer) {
		return false
	}

	if !game.textureManager.load("assets/helicopter2.png", "helicopter2", game.renderer) {
		return false
	}

	ps.gameObjects = append(ps.gameObjects, NewPlayer(500, 100, 128, 55, "helicopter"))
	ps.gameObjects = append(ps.gameObjects, NewEnemy(100, 100, 128, 55, "helicopter2"))

	fmt.Println("entering PlayState")

	return true
}

//*****************************************************************************
// onExit
//*****************************************************************************
func (ps *PlayState) onExit(game *Game) bool {
	for _, gameObject := range ps.gameObjects {
		gameObject.clean()
	}

	ps.gameObjects = ps.gameObjects[:0]

	game.textureManager.clearFromTextureMap("helicopter")

	fmt.Println("exiting PlayState")
	return true
}

//*****************************************************************************
// checkCollision
//*****************************************************************************
func (ps *PlayState) checkCollision(p1 GameObject, p2 GameObject) bool {
	var leftA, leftB, rightA, rightB, topA, topB, bottomA, bottomB float32

	leftA = p1.Position().X
	rightA = p1.Position().X + float32(p1.Width())
	topA = p1.Position().Y
	bottomA = p1.Position().Y + float32(p1.Height())

	//Calculate the sides of rect B
	leftB = p2.Position().X
	rightB = p2.Position().X + float32(p2.Width())
	topB = p2.Position().Y
	bottomB = p2.Position().Y + float32(p2.Height())

	//If any of the sides from A are outside of B
	if bottomA <= topB {
		return false
	}

	if topA >= bottomB {
		return false
	}

	if rightA <= leftB {
		return false
	}

	if leftA >= rightB {
		return false
	}

	return true
}
