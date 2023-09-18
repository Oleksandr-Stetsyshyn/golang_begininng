package arraysAndSlices

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func TextSearcher() {
	fmt.Println("------ Homework 4, task: Text searcher -------")

	reader := bufio.NewReader(os.Stdin)
	var linesSlice []string
	exitWord := "exit"
	for {

		fmt.Println("Please enter a line of text, if you want to finish enter ", exitWord)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		line = strings.TrimSpace(line)
		if line == exitWord {
			break
		}

		linesSlice = append(linesSlice, line)
	}

	textWithDivider := strings.Join(linesSlice, "\n")

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(textWithDivider))
	if err != nil {
		log.Fatal(err)
	}

	var searchWord string
	fmt.Println("Enter the search word:")
	fmt.Scan(&searchWord)

	readFile, err := os.Open("output.txt") 
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, searchWord) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

}
