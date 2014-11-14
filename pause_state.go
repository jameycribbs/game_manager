package game_manager

import (
	"fmt"
)

type PauseState struct {
	pauseID     string
	gameObjects []GameObject
}

//*****************************************************************************
// NewPauseState
//*****************************************************************************
func NewPauseState() *PauseState {
	return &PauseState{pauseID: "PAUSE"}
}

//*****************************************************************************
// getStateID
//*****************************************************************************
func (ps *PauseState) getStateID() string {
	return ps.pauseID
}

//*****************************************************************************
// update
//*****************************************************************************
func (ps *PauseState) update(game *Game) {
	for _, gameObject := range ps.gameObjects {
		gameObject.update(game)
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (ps *PauseState) render(game *Game) {
	for _, gameObject := range ps.gameObjects {
		gameObject.draw(game)
	}
}

//*****************************************************************************
// resume
//*****************************************************************************
func (ps *PauseState) resume(game *Game) {
}

//*****************************************************************************
// onEnter
//*****************************************************************************
func (ps *PauseState) onEnter(game *Game) bool {
	if !game.textureManager.load("assets/resume.png", "resumebutton", game.renderer) {
		return false
	}

	if !game.textureManager.load("assets/main.png", "mainbutton", game.renderer) {
		return false
	}

	pauseToMain := func() {
		game.gameStateMachine.changeState(NewMenuState())
	}

	resumePlay := func() {
		game.gameStateMachine.popState()
	}

	ps.gameObjects = append(ps.gameObjects, NewMenuButton(200, 100, 200, 80, "mainbutton", pauseToMain))
	ps.gameObjects = append(ps.gameObjects, NewMenuButton(200, 300, 200, 80, "resumebutton", resumePlay))

	fmt.Println("entering PauseState")

	return true
}

//*****************************************************************************
// onExit
//*****************************************************************************
func (ps *PauseState) onExit(game *Game) bool {
	for _, gameObject := range ps.gameObjects {
		gameObject.clean()
	}

	ps.gameObjects = ps.gameObjects[:0]

	game.textureManager.clearFromTextureMap("resumebutton")
	game.textureManager.clearFromTextureMap("mainbutton")

	game.inputHandler.reset()

	fmt.Println("exiting PauseState")
	return true
}
