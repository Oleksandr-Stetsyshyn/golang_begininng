package main

import (
	"flag"
	"sync"
	// "flag"
	"fmt"
	fileOrganizer "golang_beginning/hw13"
	mwcRestAPI "golang_beginning/hw14"
	"golang_beginning/hw15/observer"
	fileChangedEvent "golang_beginning/hw15/pubsub"

	rabbitmq "golang_beginning/hw16/chanel"
	rabbitmqListener "golang_beginning/hw16/listener"

	zoo "golang_beginning/lesson_02"
	game "golang_beginning/lesson_03"
	arraysAndSlices "golang_beginning/lesson_04"
	ticTacToe "golang_beginning/lesson_05"
	postOffice "golang_beginning/lesson_06/post"
	vehicle "golang_beginning/lesson_06/vehicle"
	goroutines "golang_beginning/lesson_07"
	gamers "golang_beginning/lesson_08/gamers"
	shop "golang_beginning/lesson_08/shop"
	school "golang_beginning/lesson_09/school"
	todoList "golang_beginning/lesson_09/todoList"
	"golang_beginning/lesson_10/translator"
	"golang_beginning/lesson_10/weather"
	"golang_beginning/lesson_11/numbers"
	"golang_beginning/lesson_11/words"
	mainTextProcessor "golang_beginning/lesson_12"
)

func main() {
	// go run . -lesson=15 -task=2
	var runLessonNumber int
	// fmt.Print("\033[H\033[2J")
	flag.IntVar(&runLessonNumber, "lesson", 16, "Select a homework number")

	var selectTask int
	flag.IntVar(&selectTask, "task", 1, "Select a task number")

	flag.Parse()

	switch runLessonNumber {
	case 1:
		fmt.Println("Hello world!")
	case 2:
		zoo.CatchAllAnimals()
	case 3:
		game.Game()
	case 4:
		if selectTask == 1 {
			arraysAndSlices.TextSearcher()
		} else {
			arraysAndSlices.InputAndCalculateGrade()
		}
	case 5:
		ticTacToe.StartGame()
	case 6:
		switch selectTask {
		case 1:
			postOffice.Office()
		case 2:
			vehicle.GoToRoute()
		}
	case 7:
		switch selectTask {
		case 1:
			goroutines.RunCalculateAverage()
		case 2:
			goroutines.RunCalculateAMinMax()
		}
	case 8:
		switch selectTask {
		case 1:
			shop.Shop()
		case 2:
			gamers.Game()
		}

	case 9:
		switch selectTask {
		case 1:
			todoList.Main()
		case 2:
			school.Main()
		}
	case 10:
		switch selectTask {
		case 1:
			weather.Main()
		case 2:
			translator.Main()
		}
	case 11:
		switch selectTask {
		case 1:
			numbers.Main()
		case 2:
			words.Main()
		}
	case 12:
		mainTextProcessor.Main()
	case 13:
		fileOrganizer.Main()
	case 14:
		mwcRestAPI.Main()
	case 15:
		switch selectTask {
		case 1:
			fileChangedEvent.Main()
		case 2:
			observer.Main()
		}

	case 16:
		var wg sync.WaitGroup

		wg.Add(3) // We are going to run 3 goroutines

		go func() {
			defer wg.Done()
			rabbitmq.Main()
		}()

		go func() {
			defer wg.Done()
			rabbitmq.Main()
		}()

		go func() {
			defer wg.Done()
			rabbitmqListener.Main()
		}()

		wg.Wait() // Wait for all goroutines to finish
	default:
		fmt.Println("This homework is not ready yet.")
	}

}
