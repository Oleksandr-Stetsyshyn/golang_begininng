package zoo

import (
	"fmt"
	"golang_beginning/lesson_02/zoo"
)

func CatchAllAnimals() {
	var zookeeperBob = zooStructs.Zookeeper{Name: "Bob"}
	var animal1 = zooStructs.Animal{TypeOfAnimal: "Lion"}
	var animal2 = zooStructs.Animal{TypeOfAnimal: "Tiger"}
	var animal3 = zooStructs.Animal{TypeOfAnimal: "Bear"}
	var animal4 = zooStructs.Animal{TypeOfAnimal: "Wolf"}
	var animal5 = zooStructs.Animal{TypeOfAnimal: "Fox"}

	cage1 := zooStructs.Cage{}
	cage2 := zooStructs.Cage{}
	cage3 := zooStructs.Cage{}
	cage4 := zooStructs.Cage{}
	cage5 := zooStructs.Cage{}

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
