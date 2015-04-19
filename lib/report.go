package lib

import (
	"os"
	"strings"
	"text/template"
)

var tmpl = `
  Total words: {{.TotalWords}}
  Unique words: {{.UniqueWords}}
  Average word length: {{.AverageWordLength}}
  Most used words: {{.MostUsedWords}}
`

type report struct {
	TotalWords        int
	UniqueWords       int
	AverageWordLength float64
	MostUsedWords     string
}

// Report statistics to stdout.
func Report(stat Textstat) error {
	return print(report{
		TotalWords:        stat.TotalWords(),
		UniqueWords:       stat.UniqueWords(),
		AverageWordLength: stat.AverageWordLength(),
		MostUsedWords:     strings.Join(stat.MostUsedWords(), ", "),
	})
}

// Utilities
// NOTE: Maybe I can create a separate package just to keep helpers like this.

func print(data interface{}) error {
	tmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, data)
}
