package vehicle

import (
	"fmt"
	"golang_beginning/lesson_06/vehicle/vehicleInterfaces"
	"golang_beginning/lesson_06/vehicle/vehicleStructs"
	"math/rand"
)

func generateRandomNumber(max int) int {
	return rand.Intn(max) + 1
}

func GenerateRoute(origin string, destination string, distance int) vehicleStructs.Route {
	var route vehicleStructs.Route = vehicleStructs.Route{
		Origin:          origin,
		Destination:     destination,
		Distance:        distance,
		VehiclesAtRoute: make(map[int]vehicleInterfaces.PassengersCarryable),
	}

	var vehicle vehicleInterfaces.PassengersCarryable

	for i := 0; i < route.Distance; i++ {
		switch i % 3 {
		case 0:
			vehicle = &vehicleStructs.Car{
				Transport: vehicleStructs.Transport{
					Speed:      generateRandomNumber(100),
					Passengers: generateRandomNumber(4),
				},
			}
		case 1:
			vehicle = &vehicleStructs.Airplane{
				Transport: vehicleStructs.Transport{
					Speed:      generateRandomNumber(1000),
					Passengers: generateRandomNumber(100),
				},
			}
		case 2:
			vehicle = &vehicleStructs.Train{
				Transport: vehicleStructs.Transport{
					Speed:      generateRandomNumber(200),
					Passengers: generateRandomNumber(1000),
				},
			}
		}
		route.AddVehicle(vehicle, i)
	}

	return route
}

func GoToRoute() {
	route := GenerateRoute("Uzhhorod", "Kyiv", 1000)

	fmt.Println("Start uor route:")
	for i := 0; i <= route.Distance; i++ {
		vehicle, ok := route.VehiclesAtRoute[i]
		if ok {
			fmt.Printf("Vehicle at %v km: speed: %v passengers: %v\n", i, vehicle.GetSpeed(), vehicle.GetPassengers())
		} else {
			fmt.Printf("No vehicle at %v km\n", i)
		}
	}
	fmt.Println("End uor route")
}
