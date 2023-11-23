package observer

import "fmt"

type Player struct {
	ID     int    `gorm:"primarykey" json:"id"`
	Name string
	room *GameRoom `gorm:"-"`
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
	if p.Name == player {
		return
	}
	fmt.Println(p.Name, "notified about", player, move)
}

func (p *Player) getName() string {
	return p.Name
}
