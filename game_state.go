package game_manager

type GameState interface {
	update(*Game)
	render(*Game)

	resume(*Game)

	onEnter(*Game) bool
	onExit(*Game) bool

	getStateID() string
}
