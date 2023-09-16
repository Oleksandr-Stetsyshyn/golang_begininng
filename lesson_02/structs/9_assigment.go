package structs

import "fmt"

func MakeTwins() {
	bob := humanBeing{"Bob", "Bobenko", 30, 1993}
	brother := bob

	fmt.Printf("are they the same person? %t\n", bob == brother)
	fmt.Printf("here they are:\nbob - %+v\n brother - %+v\n", bob, brother)

	brother.Name = "Sam"

	fmt.Printf("are they the same person? %t\n", bob == brother)
	fmt.Printf("here they are:\nbob - %+v\n brother - %+v\n", bob, brother)

}
