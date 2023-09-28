package vehicleStructs

type Transport struct {
	Speed      int
	Passengers int
}

func (transport *Transport) GetPassengers() int {
	return transport.Passengers
}

func (transport *Transport) GetSpeed() int {	
	return transport.Speed
}
