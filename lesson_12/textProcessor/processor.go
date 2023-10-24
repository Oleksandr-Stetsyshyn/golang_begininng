package textProcessor

import (
	"fmt"
	"regexp"
	"strings"
)

type TextWordCounter interface {
	CountWords(text string) int
}

type RegexWordCounter struct{}

func (c *RegexWordCounter) CountWords(text string) int {
	words := regexp.MustCompile(`\w+`).FindAllString(text, -1)
	return len(words)
}

type SplitWordCounter struct{}

func (c *SplitWordCounter) CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

type TextWordsCounter struct {
	WordCounter TextWordCounter
}

func (p *TextWordsCounter) SetCounter(counter TextWordCounter) {
	p.WordCounter = counter
}

func (p *TextWordsCounter) ExecuteCounter(text string) {
	wordCount := p.WordCounter.CountWords(text)
	typeOfCounter := fmt.Sprintf("%T", p.WordCounter)

	fmt.Printf("word count: %d\n", wordCount)
	fmt.Printf("type of counter: %s\n", typeOfCounter)
}
