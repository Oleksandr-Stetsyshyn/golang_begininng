package numbers

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func Main() {

	content, err := os.ReadFile("./lesson_11/numbers/numbers.txt")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("File not found")
		} else {
			panic(err)
		}
	}

	str := string(content)
	matches := regexp.MustCompile(`\(?(\d{3})\)?[ .-]?(\d{3})[ .-]?(\d{4})`).FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		fmt.Println(match[0])
	}
}
