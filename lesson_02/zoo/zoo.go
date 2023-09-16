package zoo

import "fmt"

type Zookeeper struct {
	Name string
}

type Animal struct {
	TypeOfAnimal string
}

type Cage struct {
	Animal Animal
}

func (z Zookeeper) CatchAnimalToCage(c *Cage, a Animal) {
	c.Animal = a
	fmt.Printf("%s caught %s and put to cage\n", z.Name, a.TypeOfAnimal)
}
