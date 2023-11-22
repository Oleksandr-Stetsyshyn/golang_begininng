package observer

import "fmt"

func Main() {
	state := NewGormState(NewGormConnection())

	gameRoom := &GameRoom{LevelName: "Level	777"}
	state.AddGameRoom(gameRoom)

	player1 := &Player{Name: "Sasha"}
	player2 := &Player{Name: "Pasha"}
	player3 := &Player{Name: "Bob"}

	state.AddPlayers([]*Player{player1, player2, player3})

	player1.joinRoom(gameRoom)
	player2.joinRoom(gameRoom)
	player3.joinRoom(gameRoom)

	player1.makeMove("moved to position 1")
	player2.makeMove("moved to position 2")

	player3.leaveRoom()

	state.PrinlnAllPlayers()
	state.PrinlnAllRooms()

	state.DeleteAllUsers()
	state.DeleteAllRooms()

	fmt.Println("After delete")
	state.PrinlnAllPlayers()
	state.PrinlnAllRooms()
}
