package structs

import "sync"

type humanBeing struct {
	Name     string
	Surname  string
	Age      int
	Birthday int
}

type Employee struct {
	humanBeing
	Position string
}

type ThreadSafeEmployee struct {
	Employee
	sync.Mutex
}
