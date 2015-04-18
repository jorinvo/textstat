package lib

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// Textstat ...
type Textstat struct {
	total     int
	histogram Histogram
	length    int
}

// FromReader ...
func FromReader(reader io.Reader) Textstat {
	t := new()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t.add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading:", err)
	}

	return t
}

func new() Textstat {
	return Textstat{histogram: make(map[string]int)}
}

func (t Textstat) add(word string) {
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
	return round2(divideI(t.length, t.total))
}

// MostUsedWords ...
func (t Textstat) MostUsedWords() []string {
	max := minI(len(t.histogram), 10)
	words := make([]string, max)
	for i := 0; i < max; i++ {
		words[i] = t.histogram.RemoveMax()
	}
	// fmt.Println(words)
	return words
}

func divideI(x, y int) float64 {
	return float64(x) / float64(y)
}

func round2(n float64) float64 {
	return math.Floor(n*100) / 100
}

func minI(x, y int) int {
	if x < y {
		return x
	}
	return y
}
