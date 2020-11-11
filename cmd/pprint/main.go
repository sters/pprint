package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sters/pprint/pprint"
)

func main() {
	t := flag.String("type", "", "choose type")
	flag.Parse()

	printer := pprint.ChooseType(*t)

	err := printer.ParseAndPrint(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
