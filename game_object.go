package game_manager

type GameObject interface {
	draw(game *Game)
	update(game *Game)
	clean()
	Position() Vector2D
	Width() int32
	Height() int32
}
