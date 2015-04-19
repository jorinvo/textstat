package lib

import (
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

// TODO: Test with smaller file

var text = "This sentence has exactly 10 words and no word more."

func TestFromFile(t *testing.T) {
	s, err := FromFile("texts/shakespeare.txt")
	assert.Equal(t, nil, err)
	assert.Equal(t, 904061, s.TotalWords())

	s, err = FromFile("non/existing/path")
	assert.Equal(t, err.Error(), "open non/existing/path: no such file or directory")
}

func TestFromReader(t *testing.T) {
	s, err := FromReader(strings.NewReader(text))
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, s.TotalWords())
}

func TestParse(t *testing.T) {
	s := New()
	s.Parse(text)
	assert.Equal(t, 10, s.TotalWords())
	s.Parse(text)
	assert.Equal(t, 20, s.TotalWords())
}

func TestTotalWords(t *testing.T) {
	s := New()
	s.Parse(text)
	assert.Equal(t, 10, s.TotalWords())
}

func TestUniqueWords(t *testing.T) {
	s := New()
	s.Parse("A B C C C B D F A")
	assert.Equal(t, 5, s.UniqueWords())
}

func TestAverageWordLength(t *testing.T) {
	s := New()
	s.Parse("one two three")
	assert.Equal(t, 3.66, s.AverageWordLength())
}

func TestMostUsedWords(t *testing.T) {
	words := []string{"the", "and", "i", "to", "of", "a", "my", "in", "you", "that"}
	s, err := FromFile("../texts/shakespeare.txt")
	assert.Equal(t, nil, err)
	assert.Equal(t, words, s.MostUsedWords())
}
