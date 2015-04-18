package lib

import (
	"html/template"
	"log"
	"os"
)

var tmpl = `
  total words: {{.TotalWords}}
  unique words: {{.UniqueWords}}
  average word length: {{.AverageWordLength}}
  most used words: {{.MostUsedWords}}
`

type report struct {
	TotalWords        int
	UniqueWords       int
	AverageWordLength float64
	MostUsedWords     []string
}

// Report statistics to stdout.
func Report(stat Textstat) {
	report{
		TotalWords:        stat.TotalWords(),
		UniqueWords:       stat.UniqueWords(),
		AverageWordLength: stat.AverageWordLength(),
		MostUsedWords:     stat.MostUsedWords(),
	}.print()
}

func (r report) print() {
	tmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
