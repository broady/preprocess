package main

import (
	"flag"
	"log"
	"os"

	"github.com/broady/preprocess/lib/preprocess"
)

func main() {
	prefix := flag.String("prefix", "//#", "Prefix for pragma.")
	flag.Parse()

	out, err := preprocess.Process(os.Stdin, os.Args[1:], *prefix)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(out)
}
