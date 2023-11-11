package observer

type GameRoomInterface interface {
	register(observer PlayerObserverInterface)
	deregister(observer PlayerObserverInterface)
	notifyAll()
}
