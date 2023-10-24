package words

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func Main() {

	content, err := os.ReadFile("./lesson_11/words/text.txt")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("File not found")
		} else {
			panic(err)
		}
	}

	str := string(content)

	// re1 := regexp.MustCompile(`\b[aeiou]\w*[bcdfghjklmnpqrstvwxyz]\b`)
	re1 := regexp.MustCompile(`\b([\x{0430}\x{0435}\x{0456}\x{0457}\x{043e}\x{0443}\x{044e}\x{044f}])\p{Cyrillic}*([\x{0431}\x{0432}\x{0433}\x{0434}\x{0436}\x{0437}\x{0439}\x{043a}\x{043b}\x{043c}\x{043d}\x{043f}\x{0440}\x{0441}\x{0442}\x{0444}\x{0445}\x{0446}\x{0447}\x{0448}\x{0449}])\b`)

	// re2 := regexp.MustCompile(`\b(\w)\w?\\1\b`)

	matches1 := re1.FindAllString(str, -1)

	// matches2 := re2.FindAllString(str, -1)

	fmt.Println("Words that start with a vowel and end with a consonant:")
	for _, match := range matches1 {
		fmt.Println(match)
	}

	// fmt.Println("Words that consist of two identical letters separated by any character:")
	// for _, match := range matches2 {
	// 	fmt.Println(match)
	// }
}
