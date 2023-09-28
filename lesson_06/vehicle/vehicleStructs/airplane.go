package vehicleStructs

type Airplane struct {
	Transport
}

func (airplane *Airplane) Move() bool {
	return true
}

func (airplane *Airplane) Stop() bool {
	return true
}

func (airplane *Airplane) ChangeSpeed(speed int) {
	airplane.Speed = speed
}

func (airplane *Airplane) BoardingPassengers(passengers int) {
	airplane.Passengers += passengers
}

func (airplane *Airplane) DeboardingPassenger(passengers int) {
	airplane.Passengers -= passengers
}
