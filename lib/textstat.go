package lib

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Textstat ...
type Textstat struct {
	total     int
	histogram Histogram
	length    int
}

// New ...
func New() Textstat {
	return Textstat{histogram: make(map[string]int)}
}

// FromReader ...
func FromReader(reader io.Reader) (Textstat, error) {
	t := New()
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t.add(scanner.Text())
	}
	return t, scanner.Err()
}

// FromFile ...
func FromFile(path string) (Textstat, error) {
	file, err := os.Open(path)
	if err != nil {
		return New(), err
	}
	return FromReader(file)
}

// Parse ...
func (t *Textstat) Parse(text string) {
	words := strings.Fields(text)
	for _, word := range words {
		t.add(word)
	}
}

// TODO: Remove special chars from words
func (t *Textstat) add(word string) {
	t.total++
	t.histogram[strings.ToLower(word)]++
	t.length += len([]rune(word))
}

// TotalWords ...
func (t Textstat) TotalWords() int {
	return t.total
}

// UniqueWords ...
func (t Textstat) UniqueWords() int {
	return len(t.histogram)
}

// AverageWordLength ...
func (t Textstat) AverageWordLength() float64 {
	return round2(divideInt(t.length, t.total))
}

// MostUsedWords ...
func (t Textstat) MostUsedWords() []string {
	max := minInt(len(t.histogram), 10)
	words := make([]string, max)
	for i := 0; i < max; i++ {
		words[i] = t.histogram.RemoveMax()
	}
	return words
}
