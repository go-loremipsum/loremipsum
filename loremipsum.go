package loremipsum

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// LoremIpsum is a lorem ipsum generator
type LoremIpsum struct {
	first bool
	words []string
	idx   int
	rng   *rand.Rand
	mu    sync.Mutex
}

// New returns new instance of LoremIpsum
func New() *LoremIpsum {
	return NewWithSeed(time.Now().Unix())
}

// New returns new instance of LoremIpsum with PRNG seeded with the parameter
func NewWithSeed(seed int64) *LoremIpsum {
	return NewWithSource(rand.NewSource(seed))
}

// NewWithSource returns new instance of LoremIpsum that uses random values
// from source to generate words.
func NewWithSource(source rand.Source) *LoremIpsum {
	li := new(LoremIpsum)
	li.rng = rand.New(source)
	li.first = true
	li.idx = 0
	li.shuffle()
	return li
}

// Word returns a single word of lorem ipsum
func (li *LoremIpsum) Word() string {
	defer li.shuffle()
	li.mu.Lock()
	defer li.mu.Unlock()
	return li.words[li.rng.Intn(len(li.words))]
}

// WordList returns list of words of lorem ipsum
func (li *LoremIpsum) WordList(count int) []string {
	defer li.shuffle()
	return li.words[:count]
}

// Words returns words of lorem ipsum
func (li *LoremIpsum) Words(count int) string {
	return strings.Join(li.WordList(count), " ")
}

// Sentence returns full sentence of lorem ipsum
func (li *LoremIpsum) Sentence() string {
	defer li.shuffle()
	for {
		l := int(li.gauss(24.46, 5.08))
		if l > 0 {
			words := li.words[:l]
			return li.punctuate(words)
		}
	}
}

// SentenceList returns list of sentences of lorem ipsum
func (li *LoremIpsum) SentenceList(count int) []string {
	sentences := make([]string, count)
	for idx := range sentences {
		sentences[idx] = li.Sentence()
	}
	return sentences
}

// Sentences returns sentences of lorem ipsum
func (li *LoremIpsum) Sentences(count int) string {
	return strings.Join(li.SentenceList(count), " ")
}

// Paragraph returns full paragraph of lorem ipsum
func (li *LoremIpsum) Paragraph() string {
	for {
		count := int(li.gauss(5.8, 1.93))
		if count > 0 {
			return li.Sentences(count)
		}
	}
}

// ParagraphList returns list of paragraphs of lorem ipsum
func (li *LoremIpsum) ParagraphList(count int) []string {
	paragraphs := make([]string, count)
	for idx := range paragraphs {
		paragraphs[idx] = li.Paragraph()
	}
	return paragraphs
}

// Paragraphs returns paragraphs of lorem ipsum
func (li *LoremIpsum) Paragraphs(count int) string {
	return strings.Join(li.ParagraphList(count), "\n")
}
