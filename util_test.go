package loremipsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPunctuate(t *testing.T) {
	li := New()
	sentence := []string{"this", "is", "a", "test"}
	assert.Equal(t, "This is a test.", li.punctuate(sentence))
}
