package observer

type GameRoom struct {
	ID        uint `gorm:"primarykey"`
	LevelName string
	players   []PlayerObserverInterface `gorm:"-"`
}

func (g *GameRoom) register(observer PlayerObserverInterface) {
	g.players = append(g.players, observer)
	g.notifyAll(observer.(*Player), "joined "+g.LevelName)
}

func (g *GameRoom) deregister(observer PlayerObserverInterface) {
	for i, player := range g.players {
		if player == observer {
			g.players = append(g.players[:i], g.players[i+1:]...)
			break
		}
	}

	g.notifyAll(observer.(*Player), "lived "+g.LevelName)
}

func (g *GameRoom) notifyAll(p *Player, move string) {
	for _, player := range g.players {
		player.notify(p.Name, move)
	}
}
