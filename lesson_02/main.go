package main

import (
	"fmt"
	"structs/zoo"
)

func main() {
	var zookeeperBob = zoo.Zookeeper{Name: "Bob"}
	var animal1 = zoo.Animal{TypeOfAnimal: "Lion"}
	var animal2 = zoo.Animal{TypeOfAnimal: "Tiger"}
	var animal3 = zoo.Animal{TypeOfAnimal: "Bear"}
	var animal4 = zoo.Animal{TypeOfAnimal: "Wolf"}
	var animal5 = zoo.Animal{TypeOfAnimal: "Fox"}

	cage1 := zoo.Cage{}
	cage2 := zoo.Cage{}
	cage3 := zoo.Cage{}
	cage4 := zoo.Cage{}
	cage5 := zoo.Cage{}

	zookeeperBob.CatchAnimalToCage(&cage1, animal1)
	zookeeperBob.CatchAnimalToCage(&cage2, animal2)
	zookeeperBob.CatchAnimalToCage(&cage3, animal3)
	zookeeperBob.CatchAnimalToCage(&cage4, animal5)
	zookeeperBob.CatchAnimalToCage(&cage5, animal4)

	fmt.Println("In 1 cage", cage1.Animal.TypeOfAnimal)
	fmt.Println("In 2 cage", cage2.Animal.TypeOfAnimal)
	fmt.Println("In 3 cage", cage3.Animal.TypeOfAnimal)
	fmt.Println("In 4 cage", cage4.Animal.TypeOfAnimal)
	fmt.Println("In 5 cage", cage5.Animal.TypeOfAnimal)
}
