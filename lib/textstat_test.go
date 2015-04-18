package lib

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestTotalWords(t *testing.T) {
	s := "This sentence has exactly 10 words and no word more."
	c := New().Parse(s).TotalWords()
	assert.Equal(t, c, 10)
	c = New().Parse(s).Parse(s).TotalWords()
	assert.Equal(t, c, 20)
}

func TestUniqueWords(t *testing.T) {
	s := "A B C C C B D F A"
	c := New().Parse(s).UniqueWords()
	assert.Equal(t, c, 5)
}

func TestAverageWordLength(t *testing.T) {
	s := "one two three"
	c := New().Parse(s).AverageWordLength()
	assert.Equal(t, c, 3.66)
}

func TestMostUsedWords(t *testing.T) {
	words := []string{"the", "and", "i", "to", "of", "a", "my", "in", "you", "that"}
	stat, err := FromFile("shakespeare.txt")
	if err != nil {
		t.Error(err)
	}
}
