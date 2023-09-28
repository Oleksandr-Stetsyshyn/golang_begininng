package vehicleStructs

import (
	"golang_beginning/lesson_06/vehicle/vehicleInterfaces"
)

type Route struct {
    Origin          string
    Destination     string
    Distance        int
    VehiclesAtRoute map[int]vehicleInterfaces.PassengersCarryable
}

func (route *Route) AddVehicle(vehicle vehicleInterfaces.PassengersCarryable, index int) {
    route.VehiclesAtRoute[index] = vehicle
}