package vehicleStructs

type Car struct {
	Transport
	
}

func (car *Car) Move() bool {
	return true
}


func (car *Car) Stop() bool {
	return true
}

func (car *Car) ChangeSpeed(speed int) {
	car.Speed = speed
}

func (car *Car) BoardingPassengers(passengers int) {
	car.Passengers += passengers
}

func (car *Car) DeboardingPassenger(passengers int) {
	car.Passengers -= passengers
}
