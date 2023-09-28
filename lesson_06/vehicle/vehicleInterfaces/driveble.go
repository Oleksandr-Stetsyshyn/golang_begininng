package vehicleInterfaces

type Drivable interface {
	Move() bool
	Stop() bool
	ChangeSpeed(int)
}

type PassengersCarryable interface {
	Drivable
	BoardingPassengers(int)
	DeboardingPassenger(int)
	GetPassengers() int
	GetSpeed() int
}
