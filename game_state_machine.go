package game_manager

import (
//	"fmt"
)

type GameStateMachine struct {
	game       *Game
	gameStates []GameState
}

//*****************************************************************************
// NewGameStateMachine
//*****************************************************************************
func NewGameStateMachine(game *Game) *GameStateMachine {
	return &GameStateMachine{game: game}
}

//*****************************************************************************
// pushState
//*****************************************************************************
func (gsm *GameStateMachine) pushState(state GameState) {
	gsm.gameStates = append(gsm.gameStates, state)

	state.onEnter(gsm.game)
}

//*****************************************************************************
// popState
//*****************************************************************************
func (gsm *GameStateMachine) popState() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].onExit(gsm.game)
		gsm.gameStates = gsm.gameStates[:len(gsm.gameStates)-1]
	}

	gsm.gameStates[len(gsm.gameStates)-1].resume(gsm.game)
}

//*****************************************************************************
// changeState
//*****************************************************************************
func (gsm *GameStateMachine) changeState(state GameState) {
	if len(gsm.gameStates) > 0 {
		if gsm.gameStates[len(gsm.gameStates)-1].getStateID() == state.getStateID() {
			return
		}

		gsm.gameStates[len(gsm.gameStates)-1].onExit(gsm.game)
		gsm.gameStates = gsm.gameStates[:len(gsm.gameStates)-1]
	}

	state.onEnter(gsm.game)
	gsm.gameStates = append(gsm.gameStates, state)
}

//*****************************************************************************
// clean
//*****************************************************************************
func (gsm *GameStateMachine) clean() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].onExit(gsm.game)
		gsm.gameStates = nil
	}
}

//*****************************************************************************
// update
//*****************************************************************************
func (gsm *GameStateMachine) update() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].update(gsm.game)
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (gsm *GameStateMachine) render() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].render(gsm.game)
	}
}
