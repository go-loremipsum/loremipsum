package loremipsum

import (
	"math/rand"
	"strings"
	"time"
)

// LoremIpsum is a lorem ipsum generator
type LoremIpsum struct {
	first bool
	words []string
	idx   int
}

// New returns new instance of LoremIpsum
func New() *LoremIpsum {
	rand.Seed(time.Now().Unix())

	li := new(LoremIpsum)
	li.first = true
	li.idx = 0
	li.shuffle()
	return li
}

// Word returns a single word of lorem ipsum
func (li *LoremIpsum) Word() string {
	return li.words[rand.Intn(len(li.words))]
}

// Words returns words of lorem ipsum
func (li *LoremIpsum) Words(count int) string {
	return strings.Join(li.words[:count], " ")
}

// Sentence returns full sentence of lorem ipsum
func (li *LoremIpsum) Sentence() string {
	l := int(gauss(24.46, 5.08))
	words := li.words[:l]
	return punctuate(words)
}

// Sentences returns sentences of lorem ipsum
func (li *LoremIpsum) Sentences(count int) string {
	var sentences []string
	sentences = make([]string, count)
	for idx := range sentences {
		sentences[idx] = li.Sentence()
		li.shuffle()
	}
	return strings.Join(sentences, " ")
}

// Paragraph returns full paragraph of lorem ipsum
func (li *LoremIpsum) Paragraph() string {
	return li.Sentences(int(gauss(5.8, 1.93)))
}

// Paragraphs returns paragraphs of lorem ipsum
func (li *LoremIpsum) Paragraphs(count int) string {
	var paragraphs []string
	paragraphs = make([]string, count)
	for idx := range paragraphs {
		paragraphs[idx] = li.Paragraph()
	}
	return strings.Join(paragraphs, "\\n")
}
