package main

import (
	"fmt"
	"golang_beginning/lesson_02"
	"golang_beginning/lesson_03"
	"golang_beginning/lesson_04"
	"golang_beginning/lesson_05"
	postOffice "golang_beginning/lesson_06/post"
	vehicle "golang_beginning/lesson_06/vehicle"


)

func main() {
	var runLessonNumber int
	fmt.Println("Please select a homework number: ")
	fmt.Println("1. Hello world")
	fmt.Println("2. Structs")
	fmt.Println("3. Control Statements and Loops")
	fmt.Println("4. Arrays and slices")
	fmt.Println("5. Maps")
	fmt.Println("6. Interfaces")

	fmt.Scanln(&runLessonNumber)

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

	default:
		fmt.Println("This homework is not ready yet.")
	}
}