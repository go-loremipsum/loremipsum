package loremipsum

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoremIpsum_New(t *testing.T) {
	li := New()
	assert.Equal(t, li.idx, 0)
	assert.False(t, li.first)
	assert.Len(t, li.words, len(loremIpsumWords))
}

func TestLoremIpsum_shuffle_first(t *testing.T) {
	li := New()
	assert.Equal(t, strings.Join(beg, " "), strings.Join(li.words[:8], " "))
	for idx := range rest {
		assert.Contains(t, li.words[8:], rest[idx])
	}
}

func TestLoremIpsum_shuffle(t *testing.T) {
	li := New()
	li.shuffle()
	assert.NotEqual(t, strings.Join(beg, " "), strings.Join(li.words[:8], " "))
}

func TestLoremIpsum_Word(t *testing.T) {
	li := New()
	assert.Contains(t, loremIpsumWords, li.Word())
}

func TestLoremIpsum_Words(t *testing.T) {
	li := New()
	words := strings.Split(li.Words(20), " ")
	for idx := range words {
		assert.Contains(t, loremIpsumWords, words[idx])
	}
}

func ExampleLoremIpsum_Words() {
	li := New()
	fmt.Println(li.Words(8))
	// Output:
	// lorem ipsum dolor sit amet consectetur adipiscing elit
}

// func TestLoremIpsum_Sentence(t *testing.T) {
// 	li := New()
// 	li.Sentence()
// }

func TestLoremIpsum_Sentences(t *testing.T) {
	li := New()
	s := li.Sentences(3)
	assert.Equal(t, 3, strings.Count(s, "."))
}

// func TestLoremIpsum_Paragraph(t *testing.T) {
// 	li := New()
// 	li.Paragraph()
// }

func TestLoremIpsum_Paragraphs(t *testing.T) {
	li := New()
	s := li.Paragraphs(3)
	assert.Equal(t, 2, strings.Count(s, "\n"))
}

func TestLoremIpsum_NewWithSeed(t *testing.T) {
	testSeed(t,
		NewWithSeed(1),
		NewWithSeed(1),
	)
}

func TestLoremIpsum_NewWithSource(t *testing.T) {
	testSeed(t,
		NewWithSource(rand.New(rand.NewSource(1))),
		NewWithSource(rand.New(rand.NewSource(1))),
	)
}

func testSeed(t *testing.T, liSeeded1, liSeeded2 *LoremIpsum) {
	liRandom := New()

	var liRandomList []string
	var liSeeded1List []string
	var liSeeded2List []string
	reqCount := 5
	for i := 1; i <= reqCount; i++ {
		liRandomList = append(liRandomList, liRandom.Paragraphs(i+5))
		liSeeded1List = append(liSeeded1List, liSeeded1.Paragraphs(i+5))
		liSeeded2List = append(liSeeded2List, liSeeded2.Paragraphs(i+5))
	}
	for i := range liRandomList {
		// should differ from randomly generated
		assert.NotEqual(t, liRandomList[i], liSeeded1List[i])
		// should be same with same seed
		assert.Equal(t, liSeeded1List[i], liSeeded2List[i])
		// should not be same with next entry
		if i < len(liRandomList)-1 {
			assert.NotEqual(t, liSeeded1List[i], liSeeded1List[i+1])
		}
	}
}
