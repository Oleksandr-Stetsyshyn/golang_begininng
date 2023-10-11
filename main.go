package main

import (
	"fmt"
	"golang_beginning/lesson_02"
	"golang_beginning/lesson_03"
	"golang_beginning/lesson_04"
	"golang_beginning/lesson_05"
	postOffice "golang_beginning/lesson_06/post"
	vehicle "golang_beginning/lesson_06/vehicle"
	"golang_beginning/lesson_07"
	gamers "golang_beginning/lesson_08/gamers"
	shop "golang_beginning/lesson_08/shop"
	school "golang_beginning/lesson_09/school"
	todoList "golang_beginning/lesson_09/todoList"

	"golang_beginning/lesson_10/translator"
	"golang_beginning/lesson_10/weather"
)

func main() {
	var runLessonNumber int = 10
	// fmt.Println("Please select a homework number: ")
	// fmt.Println("1. Hello world")
	// fmt.Println("2. Structs")
	// fmt.Println("3. Control Statements and Loops")
	// fmt.Println("4. Arrays and slices")
	// fmt.Println("5. Maps")
	// fmt.Println("6. Interfaces")
	// fmt.Println("7. Goroutines")
	// fmt.Scanln(&runLessonNumber)
	fmt.Print("\033[H\033[2J")
	switch runLessonNumber {
	case 1:
		fmt.Println("Hello world!")
	case 2:
		zoo.CatchAllAnimals()
	case 3:
		game.Game()
	case 4:
		{
			var selectTask int
			fmt.Println("Select task rom homework 4")
			fmt.Println("1. Text searcher")
			fmt.Println("2. Calculate grade")
			fmt.Scanln(&selectTask)
			if selectTask == 1 {
				arraysAndSlices.TextSearcher()
			} else {
				arraysAndSlices.InputAndCalculateGrade()
			}
		}
	case 5:
		ticTacToe.StartGame()
	case 6:
		var selectTask int
		fmt.Println("Select task rom homework 6")
		fmt.Println("1. Post office")
		fmt.Println("2. Vehicle")
		fmt.Scanln(&selectTask)

		switch selectTask {
		case 1:
			postOffice.Office()
		case 2:
			vehicle.GoToRoute()
		}
	case 7:
		var selectTask int
		fmt.Println("Select task rom homework 7")
		fmt.Println("1. Calculate average")
		fmt.Println("2. Calculate min and max")
		fmt.Scanln(&selectTask)
		switch selectTask {
		case 1:
			goroutines.RunCalculateAverage()
		case 2:
			goroutines.RunCalculateAMinMax()
		}
	case 8:
		var selectTask int
		fmt.Println("Select task rom homework 8")
		fmt.Println("1. Shop")
		fmt.Println("2. Gamers")
		fmt.Scanln(&selectTask)

		switch selectTask {
		case 1:
			shop.Shop()
		case 2:
			gamers.Game()
		}

	case 9:
		var selectTask int
		fmt.Println("Select task rom homework 9")
		fmt.Println("1. Todo list")
		fmt.Println("2. School")
		fmt.Scanln(&selectTask)
		switch selectTask {
		case 1:
			todoList.Main()
		case 2:
			school.Main()
		}
	case 10:
		var selectTask int
		fmt.Println("Select task rom homework 10")
		fmt.Println("1. weather")
		fmt.Println("2. translator")
		fmt.Scanln(&selectTask)

		switch selectTask {
		case 1:
			weather.Main()
		case 2:
			translator.Main()
		}

	default:
		fmt.Println("This homework is not ready yet.")
	}

}
