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
	re1 := regexp.MustCompile(`(?i)(^|\s)([аоуіе])\p{Cyrillic}+([бвгджзйклмнпрстфхцчш])($|\s|\.|\,)`)

	matches1 := re1.FindAllString(str, -1)

	fmt.Println("Words that start with a vowel and end with a consonant:")
	for _, match := range matches1 {
		fmt.Println(match)
	}
}
