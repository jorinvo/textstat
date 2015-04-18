package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/jorin-vogel/textstat-go/lib"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	hasArg := len(arg) != 0

	var reader io.Reader
	if hasArg {
		reader = openFile(arg)
	} else {
		reader = os.Stdin
	}

	lib.Report(lib.FromReader(reader))
}

func openFile(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
