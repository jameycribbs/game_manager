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
	keystates         []uint8
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
// onMouseButtonDown
//*****************************************************************************
func (ih *InputHandler) onMouseButtonDown(event sdl.Event) {
	var mouseButtonEvent = event.(*sdl.MouseButtonEvent)

	if mouseButtonEvent.Button == sdl.BUTTON_LEFT {
		ih.mouseButtonStates[LEFT] = true
	}
	if mouseButtonEvent.Button == sdl.BUTTON_MIDDLE {
		ih.mouseButtonStates[MIDDLE] = true
	}
	if mouseButtonEvent.Button == sdl.BUTTON_RIGHT {
		ih.mouseButtonStates[RIGHT] = true
	}
}

//*****************************************************************************
// onMouseButtonUp
//*****************************************************************************
func (ih *InputHandler) onMouseButtonUp(event sdl.Event) {
	var mouseButtonEvent = event.(*sdl.MouseButtonEvent)

	if mouseButtonEvent.Button == sdl.BUTTON_LEFT {
		ih.mouseButtonStates[LEFT] = false
	}
	if mouseButtonEvent.Button == sdl.BUTTON_MIDDLE {
		ih.mouseButtonStates[MIDDLE] = false
	}
	if mouseButtonEvent.Button == sdl.BUTTON_RIGHT {
		ih.mouseButtonStates[RIGHT] = false
	}
}

//*****************************************************************************
// onMouseMove
//*****************************************************************************
func (ih *InputHandler) onMouseMove(event sdl.Event) {
	var mouseMotionEvent = event.(*sdl.MouseMotionEvent)

	ih.mousePosition.X = mouseMotionEvent.X
	ih.mousePosition.Y = mouseMotionEvent.Y
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
// isKeyDown
//*****************************************************************************
func (ih *InputHandler) isKeyDown(key sdl.Scancode) bool {
	if len(ih.keystates) != 0 {
		if ih.keystates[key] == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}

//*****************************************************************************
// onKeyDown
//*****************************************************************************
func (ih *InputHandler) onKeyDown() {
	ih.keystates = sdl.GetKeyboardState()
}

//*****************************************************************************
// onKeyUp
//*****************************************************************************
func (ih *InputHandler) onKeyUp() {
	ih.keystates = sdl.GetKeyboardState()
}

//*****************************************************************************
// update
//*****************************************************************************
func (ih *InputHandler) update(game *Game) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			game.SetRunning(false)
		case *sdl.MouseMotionEvent:
			ih.onMouseMove(event)
		case *sdl.MouseButtonEvent:
			if t.State == 1 {
				ih.onMouseButtonDown(event)
			}
			if t.State == 0 {
				ih.onMouseButtonUp(event)
			}
		case *sdl.KeyDownEvent:
			ih.onKeyDown()
		case *sdl.KeyUpEvent:
			ih.onKeyUp()
		}
	}
}

//*****************************************************************************
// clean
//*****************************************************************************
func (ih *InputHandler) clean() {
}

//*****************************************************************************
// reset
//*****************************************************************************
func (ih *InputHandler) reset() {
	ih.mouseButtonStates[LEFT] = false
	ih.mouseButtonStates[RIGHT] = false
	ih.mouseButtonStates[MIDDLE] = false
}
