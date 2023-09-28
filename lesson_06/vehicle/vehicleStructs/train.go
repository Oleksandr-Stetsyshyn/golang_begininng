package vehicleStructs

type Train struct {
	Transport
}

func (train *Train) Move() bool {
	return true
}

func (train *Train) Stop() bool {
	return true
}

func (train *Train) ChangeSpeed(speed int) {
	train.Speed = speed
}

func (train *Train) BoardingPassengers(passengers int) {
	train.Passengers += passengers
}

func (train *Train) DeboardingPassenger(passengers int) {
	train.Passengers -= passengers
}
