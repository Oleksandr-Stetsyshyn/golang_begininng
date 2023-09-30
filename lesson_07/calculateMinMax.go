package goroutines

import (
	"fmt"
	"math/rand"
)

func RunCalculateAMinMax() {
	rng := 100

	numbChannel := make(chan int, 5)
	quit := make(chan bool)
	endProgram := make(chan bool)

	go randomNumbersFromRange(rng, numbChannel, quit, endProgram)
	go findMinMax(numbChannel, quit)

	<-endProgram
}

func randomNumbersFromRange(rng int, numbChannel chan int, quit chan bool, endProgram chan bool) {

	for i := 0; i < 5; i++ {
		randNumb := rand.Intn(rng)
		fmt.Print(randNumb, " ")
		numbChannel <- randNumb
	}

	for closed := range quit {
		if closed {
			min, max := <-numbChannel, <-numbChannel
			fmt.Println("Min:", min, "Max:", max)
			close(numbChannel)
			close(quit)
		}
	}
	endProgram <- true
}

func findMinMax(numbChannel chan int, quit chan bool) {

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

	quit <- true
	
	numbChannel <- min
	numbChannel <- max
}
