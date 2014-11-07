package game_manager

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Game struct {
	window         *sdl.Window
	renderer       *sdl.Renderer
	event          sdl.Event
	running        bool
	currentFrame   int32
	textureManager *TextureManager
	inputHandler   *InputHandler
	gameObjects    []GameObject
}

//*****************************************************************************
// NewGame
//*****************************************************************************
func NewGame() *Game {
	return &Game{running: false, textureManager: NewTextureManager(), inputHandler: NewInputHandler()}
}

//*****************************************************************************
// Setup
//*****************************************************************************
func (game *Game) Setup(title string, xpos int, ypos int, width int, height int, fullscreen bool) bool {
	var flags uint32

	flags = 0

	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN
	}

	if setupSDL() != true {
		return false
	}

	if game.createWindow(title, xpos, ypos, width, height, flags) != true {
		return false
	}

	if game.createRenderer() != true {
		return false
	}

	game.textureManager.load("assets/animate-alpha.png", "animate", game.renderer)

	game.gameObjects = append(game.gameObjects, NewPlayer(100, 100, 128, 82, "animate"))
	game.gameObjects = append(game.gameObjects, NewEnemy(300, 300, 128, 82, "animate"))

	game.running = true

	return true
}

//*****************************************************************************
// HandleEvents
//*****************************************************************************
func (game *Game) HandleEvents() {
	game.inputHandler.update(game)
}

//*****************************************************************************
// Update
//*****************************************************************************
func (game *Game) Update() {
	for _, gameObject := range game.gameObjects {
		gameObject.update(game)
	}
}

//*****************************************************************************
// Render
//*****************************************************************************
func (game *Game) Render() {
	game.renderer.Clear()

	for _, gameObject := range game.gameObjects {
		gameObject.draw(game)
	}

	game.renderer.Present()
}

//*****************************************************************************
// Clean
//*****************************************************************************
func (game *Game) Clean() {
	game.inputHandler.clean()

	game.renderer.Destroy()
	game.window.Destroy()

	sdl.Quit()
}

//*****************************************************************************
// SetRunning
//*****************************************************************************
func (game *Game) SetRunning(r bool) {
	game.running = r
}

//*****************************************************************************
// Running
//*****************************************************************************
func (game *Game) Running() bool {
	return game.running
}

//*****************************************************************************
// setupSDL
//*****************************************************************************
func setupSDL() bool {
	if sdl.Init(sdl.INIT_EVERYTHING) < 0 {

		fmt.Fprintf(os.Stderr, "SDL failed to initialize: %s", sdl.GetError())

		return false

	} else {

		fmt.Println("SDL init success")

		return true
	}
}

//*****************************************************************************
// createWindow
//*****************************************************************************
func (game *Game) createWindow(title string, xpos int, ypos int, width int, height int, flags uint32) bool {
	game.window = sdl.CreateWindow(title, xpos, ypos, width, height, flags)

	if game.window == nil {
		fmt.Fprintf(os.Stderr, "Window creation failed: %s", sdl.GetError())

		return false

	} else {

		fmt.Println("window creation success")

		return true
	}
}

//*****************************************************************************
// createRenderer
//*****************************************************************************
func (game *Game) createRenderer() bool {
	game.renderer = sdl.CreateRenderer(game.window, -1, 0)

	if game.renderer == nil {
		fmt.Fprintf(os.Stderr, "Renderer creation failed: %s", sdl.GetError())

		return false

	} else {

		fmt.Println("renderer creation success")

		game.renderer.SetDrawColor(255, 0, 0, 255)

		return true
	}
}
