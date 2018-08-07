// +build go1.10

package loremipsum

import "math/rand"

// Shuffle the words
func (li *LoremIpsum) shuffle() {
	var words []string

	if !li.first {
		words = make([]string, len(loremIpsumWords))
		copy(words, loremIpsumWords[:])
	} else {
		words = make([]string, len(rest))
		copy(words, rest)
	}
	rand.Shuffle(len(words), func(i int, j int) {
		words[i], words[j] = words[j], words[i]
	})
	if li.first {
		b := make([]string, len(beg))
		copy(b, beg)
		// words, b = b, words
		// words = append(words, b...)
		words = append(b, words...)
	}
	li.words = words
	li.first = false
}
