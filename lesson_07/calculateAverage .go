package goroutines

import (
	"fmt"
	"math/rand"
)

func RunCalculateAverage () {
	fmt.Println("Start")

	var intCh chan int = make(chan int)
	var floatCh chan float64 = make(chan float64, 1)
	var quit chan bool = make(chan bool, 1)

	go generateRandomNumbers(intCh)
	go calculateAverage(intCh, floatCh)
	go printAverage(floatCh, quit)

	<-quit
	
	fmt.Println("The End")
}

func generateRandomNumbers(intCh chan int) {
	for i := 0; i < 3; i++ {
		randNumb := rand.Intn(1000)
		fmt.Println(randNumb)
		intCh <- randNumb
	}
	close(intCh)
}

func calculateAverage(intCh chan int, floatCh chan float64) {
	var slice []int
	for num := range intCh {
		slice = append(slice, num)
	}

	var sum int
	for _, num := range slice {
		sum += num
	}
	avg := float64(sum) / float64(len(slice))
	floatCh <- avg
}

func printAverage(floatCh chan float64, quit chan bool) {
	num := <-floatCh
	fmt.Println("Average is:", num)
	quit <- true
	close(floatCh)
	close(quit)

}
