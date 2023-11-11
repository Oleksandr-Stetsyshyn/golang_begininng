package observer

func Main() {
	{
		gameRoom := &GameRoom{}
		player1 := &Player{name: "Sasha"}
		player2 := &Player{name: "Pasha"}

		gameRoom.register(player1)
		gameRoom.register(player2)

		gameRoom.notifyAll(player1, "Move to position 1")
		gameRoom.notifyAll(player2, "Move to position 2")
	}
}
