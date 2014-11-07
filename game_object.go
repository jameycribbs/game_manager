package game_manager

type GameObject interface {
	draw(game *Game)
	update(game *Game)
	clean()
}
