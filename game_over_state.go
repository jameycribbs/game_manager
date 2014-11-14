package game_manager

import (
	"fmt"
)

type GameOverState struct {
	gameOverID  string
	gameObjects []GameObject
}

//*****************************************************************************
// NewGameOverState
//*****************************************************************************
func NewGameOverState() *GameOverState {
	return &GameOverState{gameOverID: "GAMEOVER"}
}

//*****************************************************************************
// getStateID
//*****************************************************************************
func (gos *GameOverState) getStateID() string {
	return gos.gameOverID
}

//*****************************************************************************
// update
//*****************************************************************************
func (gos *GameOverState) update(game *Game) {
	for _, gameObject := range gos.gameObjects {
		gameObject.update(game)
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (gos *GameOverState) render(game *Game) {
	for _, gameObject := range gos.gameObjects {
		gameObject.draw(game)
	}
}

//*****************************************************************************
// resume
//*****************************************************************************
func (gos *GameOverState) resume(game *Game) {
}

//*****************************************************************************
// onEnter
//*****************************************************************************
func (gos *GameOverState) onEnter(game *Game) bool {
	if !game.textureManager.load("assets/gameover.png", "gameovertext", game.renderer) {
		return false
	}

	if !game.textureManager.load("assets/main.png", "mainbutton", game.renderer) {
		return false
	}

	if !game.textureManager.load("assets/restart.png", "restartbutton", game.renderer) {
		return false
	}

	gameOverToMain := func() {
		game.gameStateMachine.changeState(NewMenuState())
	}

	restartPlay := func() {
		game.gameStateMachine.changeState(NewPlayState())
	}

	gos.gameObjects = append(gos.gameObjects, NewAnimatedGraphic(200, 100, 190, 30, "gameovertext", 2))
	gos.gameObjects = append(gos.gameObjects, NewMenuButton(200, 200, 200, 80, "mainbutton", gameOverToMain))
	gos.gameObjects = append(gos.gameObjects, NewMenuButton(200, 300, 200, 80, "restartbutton", restartPlay))

	fmt.Println("entering GameOverState")

	return true
}

//*****************************************************************************
// onExit
//*****************************************************************************
func (gos *GameOverState) onExit(game *Game) bool {
	for _, gameObject := range gos.gameObjects {
		gameObject.clean()
	}

	gos.gameObjects = gos.gameObjects[:0]

	game.textureManager.clearFromTextureMap("gameovertext")
	game.textureManager.clearFromTextureMap("restartbutton")
	game.textureManager.clearFromTextureMap("mainbutton")

	game.inputHandler.reset()

	fmt.Println("exiting GameOverState")
	return true
}
