package structs

import "fmt"

type Engine struct {
	volume  int
	started bool
	power   int
}

type Car struct {
	brand  string
	engine Engine
}

// pointer receiver
func (car *Car) StartStop() {
	if car.engine.IsRunning() {
		car.engine.ShutDown()
	} else {
		car.engine.Ignite()
	}
}

func (engine *Engine) IsRunning() bool {
	return engine.started
}

func (engine *Engine) Ignite() {
	engine.started = true
}

func (engine *Engine) ShutDown() {
	engine.started = false
}

func LetsRide() {
	bmw := Car{brand: "BMW", engine: Engine{started: false, volume: 2000, power: 200}}

	fmt.Printf("Car state: %+v\n", bmw)

	bmw.StartStop()

	fmt.Printf("Engine is running: %t\n", bmw.engine.IsRunning())

	fmt.Printf("Car state: %+v\n", bmw)

	bmw.StartStop()

	fmt.Printf("Engine is running: %t\n", bmw.engine.IsRunning())

	fmt.Printf("Car state: %+v\n", bmw)

}
