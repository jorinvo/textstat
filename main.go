package main

import (
	"flag"
	"log"
	"os"

	"github.com/jorin-vogel/textstat-go/lib"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	hasArg := len(arg) != 0

	var (
		stats lib.Textstat
		err   error
	)

	if hasArg {
		stats, err = lib.FromFile(arg)
	} else {
		stats, err = lib.FromReader(os.Stdin)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = lib.Report(stats)

	if err != nil {
		log.Fatal(err)
	}
}
