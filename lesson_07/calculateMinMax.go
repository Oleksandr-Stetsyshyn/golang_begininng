package goroutines

import (
	"fmt"
	"math/rand"
)

func RunCalculateAMinMax() {
	rng := 100
	numbChannel := make(chan int)
	endProgram := make(chan bool)

	go randomNumbersFromRange(rng, numbChannel, endProgram)
	go findMinMax(numbChannel)

	<-endProgram
}

func randomNumbersFromRange(rng int, numbChannel chan int, endProgram chan bool) {

	for i := 0; i < 5; i++ {
		randNumb := rand.Intn(rng)
		fmt.Print(randNumb, " ")
		numbChannel <- randNumb
	}

	min, max := <-numbChannel, <-numbChannel
	
	fmt.Println("Min:", min, "Max:", max)
	close(numbChannel)
	endProgram <- true
}

func findMinMax(numbChannel chan int) {

	var slice []int
	for i := 0; i < 5; i++ {
		num := <-numbChannel
		slice = append(slice, num)
	}

	var min, max int
	for _, num := range slice {
		if num < min || min == 0 {
			min = num
		}
		if num > max {
			max = num
		}
	}

	numbChannel <- min
	numbChannel <- max
}
