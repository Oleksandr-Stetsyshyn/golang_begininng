package observer

import "fmt"

type Player struct {
	name string
}

func (p *Player) update(player string, move string) {
	if p.name == player {
		return
	}
	fmt.Println(p.name, "notified that player", player, "made a move:", move)
}
