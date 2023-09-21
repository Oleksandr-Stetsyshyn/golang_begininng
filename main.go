package main

import (
	"fmt"
	"golang_beginning/lesson_02"
	"golang_beginning/lesson_03"
	"golang_beginning/lesson_04"
	"golang_beginning/lesson_05"
)

func main() {
	var runLessonNumber int
	fmt.Println("Please select a homework number: ")
	fmt.Println("1. Hello world")
	fmt.Println("2. Structs")
	fmt.Println("3. Control Statements and Loops")
	fmt.Println("4. Arrays and slices")
	fmt.Println("5. Maps")

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
	default:
		fmt.Println("This homework is not ready yet.")
	}
}
