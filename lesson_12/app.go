package mainTextProcessor

import (
	"errors"
	"flag"
	"fmt"
	"golang_beginning/lesson_12/textProcessor"
	"io/fs"
	"os"
)

func Main() {
	// go run main.go -decorator=both -algorithm=regex -save=true
	var decorator string
	var algorithm string
	var save bool

	flag.StringVar(&decorator, "decorator", "both", "Choose a decorator: html, space, both")
	flag.StringVar(&algorithm, "algorithm", "regex", "Choose a word counting algorithm: regex, split")
	flag.BoolVar(&save, "save", false, "Save the edited file")

	flag.Parse()

	content, err := os.ReadFile("./lesson_12/page.html")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("File not found")
		} else {
			panic(err)
		}
	}

	text := string(content)

	var textDecorator textProcessor.TextDecorator

	switch decorator {
	case "html":
		textDecorator = &textProcessor.HtmlTagRemover{Decorator: &textProcessor.RawText{}}
	case "space":
		textDecorator = &textProcessor.DoubleSpaceRemover{Decorator: &textProcessor.RawText{}}
	case "both":
		textDecorator = &textProcessor.DoubleSpaceRemover{Decorator: &textProcessor.HtmlTagRemover{Decorator: &textProcessor.RawText{}}}
	default:
		fmt.Println("Invalid decorator choice. Using both decorators.")
		textDecorator = &textProcessor.DoubleSpaceRemover{Decorator: &textProcessor.HtmlTagRemover{Decorator: &textProcessor.RawText{}}}
	}

	editedText := textDecorator.DecorateText(text)

	if save {
		err = os.WriteFile("./lesson_12/edited_text.txt", []byte(editedText), 0644)
		if err != nil {
			panic(err)
		}
	}

	var wordCounter *textProcessor.TextWordsCounter

	switch algorithm {
	case "regex":
		wordCounter = &textProcessor.TextWordsCounter{WordCounter: &textProcessor.RegexWordCounter{}}
	case "split":
		wordCounter = &textProcessor.TextWordsCounter{WordCounter: &textProcessor.SplitWordCounter{}}
	default:
		fmt.Println("Invalid algorithm choice. Using regex word counter.")
		wordCounter = &textProcessor.TextWordsCounter{WordCounter: &textProcessor.RegexWordCounter{}}
	}
	wordCounter.ExecuteCounter(editedText)
}
