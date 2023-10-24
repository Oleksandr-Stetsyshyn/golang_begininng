package textProcessor

import (
	// "fmt"
	"regexp"
)

type TextDecorator interface {
	DecorateText(text string) string
}

type RawText struct{}

func (d *RawText) DecorateText(text string) string {
	return text
}

type HtmlTagRemover struct {
	Decorator TextDecorator
}

func (d *HtmlTagRemover) DecorateText(text string) string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(d.Decorator.DecorateText(text), "")
}

type DoubleSpaceRemover struct {
	Decorator TextDecorator
}

func (d *DoubleSpaceRemover) DecorateText(text string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(d.Decorator.DecorateText(text), " ")
}
