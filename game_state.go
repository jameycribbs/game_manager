package game_manager

type GameState interface {
	update()
	render()

	onEnter() bool
	onExit() bool

	getStateID() string
}
