package game_manager

import (
	"fmt"
)

type MenuState struct {
	menuID      string
	gameObjects []GameObject
}

//*****************************************************************************
// NewMenuState
//*****************************************************************************
func NewMenuState() *MenuState {
	return &MenuState{menuID: "MENU"}
}

//*****************************************************************************
// getStateID
//*****************************************************************************
func (ms *MenuState) getStateID() string {
	return ms.menuID
}

//*****************************************************************************
// update
//*****************************************************************************
func (ms *MenuState) update(game *Game) {
	for _, gameObject := range ms.gameObjects {
		gameObject.update(game)
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (ms *MenuState) render(game *Game) {
	for _, gameObject := range ms.gameObjects {
		gameObject.draw(game)
	}
}

//*****************************************************************************
// resume
//*****************************************************************************
func (ms *MenuState) resume(game *Game) {
}

//*****************************************************************************
// onEnter
//*****************************************************************************
func (ms *MenuState) onEnter(game *Game) bool {
	if !game.textureManager.load("assets/button.png", "playbutton", game.renderer) {
		return false
	}

	if !game.textureManager.load("assets/exit.png", "exitbutton", game.renderer) {
		return false
	}

	menuToPlay := func() {
		fmt.Println("Play button clicked")
	}

	exitFromMenu := func() {
		fmt.Println("Exit button clicked")
	}

	ms.gameObjects = append(ms.gameObjects, NewMenuButton(100, 100, 400, 100, "playbutton", menuToPlay))
	ms.gameObjects = append(ms.gameObjects, NewMenuButton(100, 300, 400, 100, "exitbutton", exitFromMenu))

	return true
}

//*****************************************************************************
// onExit
//*****************************************************************************
func (ms *MenuState) onExit(game *Game) bool {
	for _, gameObject := range ms.gameObjects {
		gameObject.clean()
	}

	ms.gameObjects = ms.gameObjects[:0]

	game.textureManager.clearFromTextureMap("playbutton")
	game.textureManager.clearFromTextureMap("exitbutton")

	fmt.Println("exiting MenuState")
	return true
}
