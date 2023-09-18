package arraysAndSlices

import "fmt"

func InputAndCalculateGrade() {
	fmt.Println("------ Homework 4, task: Average grades calculator -------")
	var grades []float32

	for {
		var grade float32
		numbToExit := 0
		fmt.Println("Please enter a grade, if you want to finish enter ", numbToExit)
		fmt.Scan(&grade)

		if grade == float32(numbToExit) {
			break
		}
		grades = append(grades, grade)

	}

	averageMark := calculateAverage(grades)
	fmt.Printf("average mark is: %.2f", averageMark)
}

func calculateAverage(grades []float32) float32 {
	var sum float32

	for _, grade := range grades {
		sum += grade
	}
	return sum / float32(len(grades))

}
