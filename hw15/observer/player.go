package observer

import "fmt"

type Player struct {
	name string
	room *GameRoom
}

func (p *Player) joinRoom(room *GameRoom) {
	p.room = room
	room.register(p)
}

func (p *Player) leaveRoom() {
	p.room.deregister(p)
	p.room = nil
}

func (p *Player) makeMove(move string) {
	p.room.notifyAll(p, move)
}

func (p *Player) notify(player string, move string) {
	if p.name == player {
		return
	}
	fmt.Println(p.name, "notified about", player, move)
}

func (p *Player) getName() string {
	return p.name
}
