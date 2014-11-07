package game_manager

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	LEFT   = iota
	MIDDLE = iota
	RIGHT  = iota
)

type InputHandler struct {
	mouseButtonStates [3]bool
	mousePosition     Vector2D
}

//*****************************************************************************
// NewInputHandler
//*****************************************************************************
func NewInputHandler() *InputHandler {
	var ih InputHandler

	for i := 0; i < 3; i++ {
		ih.mouseButtonStates[i] = false
	}

	return &ih
}

//*****************************************************************************
// getMouseButtonState
//*****************************************************************************
func (ih *InputHandler) getMouseButtonState(buttonNumber int) bool {
	return ih.mouseButtonStates[buttonNumber]
}

//*****************************************************************************
// getMousePosition
//*****************************************************************************
func (ih *InputHandler) getMousePosition() Vector2D {
	return ih.mousePosition
}

//*****************************************************************************
// update
//*****************************************************************************
func (ih *InputHandler) update(game *Game) {
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			game.SetRunning(false)

		case *sdl.MouseMotionEvent:
			ih.mousePosition.X = t.X
			ih.mousePosition.Y = t.Y
		case *sdl.MouseButtonEvent:
			if t.State == 1 {
				if t.Button == sdl.BUTTON_LEFT {
					ih.mouseButtonStates[LEFT] = true
				}
				if t.Button == sdl.BUTTON_MIDDLE {
					ih.mouseButtonStates[MIDDLE] = true
				}
				if t.Button == sdl.BUTTON_RIGHT {
					ih.mouseButtonStates[RIGHT] = true
				}
			}
			if t.State == 0 {
				if t.Button == sdl.BUTTON_LEFT {
					ih.mouseButtonStates[LEFT] = false
				}
				if t.Button == sdl.BUTTON_MIDDLE {
					ih.mouseButtonStates[MIDDLE] = false
				}
				if t.Button == sdl.BUTTON_RIGHT {
					ih.mouseButtonStates[RIGHT] = false
				}
			}
		}
	}
}

//*****************************************************************************
// clean
//*****************************************************************************
func (ih *InputHandler) clean() {
}
