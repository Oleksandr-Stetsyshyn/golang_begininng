package observer

type PlayerObserverInterface interface {
	joinRoom(*GameRoom)
	leaveRoom()
	makeMove(string)
	notify(string, string)
	getName() string
}
