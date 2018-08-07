package loremipsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPunctuate(t *testing.T) {
	sentence := []string{"this", "is", "a", "test"}
	assert.Equal(t, "This is a test.", punctuate(sentence))
}
