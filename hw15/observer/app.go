package observer

func Main() {
	{
		gameRoom := &GameRoom{levelName: "Level 777"}
		player1 := &Player{name: "Sasha"}
		player2 := &Player{name: "Pasha"}
		player3 := &Player{name: "Bob"}

		player1.joinRoom(gameRoom)
		player2.joinRoom(gameRoom)
		player3.joinRoom(gameRoom)

		player1.makeMove("moved to position 1")
		player2.makeMove("moved to position 2")

		player3.leaveRoom()

	}
}
