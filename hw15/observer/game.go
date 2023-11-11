package observer

type GameRoom struct {
	players []PlayerObserverInterface
}

func (g *GameRoom) register(observer PlayerObserverInterface) {
	g.players = append(g.players, observer)
}

func (g *GameRoom) deregister(observer PlayerObserverInterface) {
	for i, player := range g.players {
		if player == observer {
			g.players = append(g.players[:i], g.players[i+1:]...)
			break
		}
	}
}

func (g *GameRoom) notifyAll(p *Player, move string) {
	for _, player := range g.players {
		player.update(p.name, move)
	}
}
