package game_manager

type GameStateMachine struct {
	gameStates []*GameState
}

//*****************************************************************************
// NewGameStateMachine
//*****************************************************************************
func NewGameStateMachine() *GameStateMachine {
	return &GameStateMachine{}
}

//*****************************************************************************
// pushState
//*****************************************************************************
func (gsm *GameStateMachine) pushState(state *GameState) {
	gsm.gameStates = append(gsm.gameStates, state)

	state.onEnter()
}

//*****************************************************************************
// popState
//*****************************************************************************
func (gsm *GameStateMachine) popState() {
	if len(gsm.gameStates) > 0 {
		if gsm.gameStates[len(gsm.gameStates)-1].onExit() {
			gsm.gameStates = gsm.gameStates[:len(gsm.gameStates)-1]
		}
	}

	gsm.gameStates[len(gsm.gameStates)-1].resume()
}

//*****************************************************************************
// changeState
//*****************************************************************************
func (gsm *GameStateMachine) changeState(state *GameState) {
	if len(gsm.gameStates) > 0 {
		if gsm.gameStates[len(gsm.gameStates)-1].getStateID() == state.getStateID() {
			return
		}

		gsm.gameStates[len(gsm.gameStates)-1].onExit()
		gsm.gameStates = gsm.gameStates[:len(gsm.gameStates)-1]
	}

	gsm.gameStates = append(gsm.gameStates, state)
	state.onEnter()
}

//*****************************************************************************
// clean
//*****************************************************************************
func (gsm *GameStateMachine) clean() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].onExit()
		gsm.gameStates = nil
	}
}

//*****************************************************************************
// update
//*****************************************************************************
func (gsm *GameStateMachine) update() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].update()
	}
}

//*****************************************************************************
// render
//*****************************************************************************
func (gsm *GameStateMachine) render() {
	if len(gsm.gameStates) > 0 {
		gsm.gameStates[len(gsm.gameStates)-1].render()
	}
}
