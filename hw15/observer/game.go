package observer

type GameRoom struct {
	levelName string
	players   []PlayerObserverInterface
}

func (g *GameRoom) register(observer PlayerObserverInterface) {
	g.players = append(g.players, observer)
	g.notifyAll(observer.(*Player), "joined "+g.levelName)
}

func (g *GameRoom) deregister(observer PlayerObserverInterface) {
	for i, player := range g.players {
		if player == observer {
			g.players = append(g.players[:i], g.players[i+1:]...)
			break
		}
	}

	g.notifyAll(observer.(*Player), "lived "+g.levelName)
}

func (g *GameRoom) notifyAll(p *Player, move string) {
	for _, player := range g.players {
		player.notify(p.name, move)
	}
}
