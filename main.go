package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [GRAMMAR FILE]\n", os.Args[0])
		os.Exit(1)
	}

	grammar_file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer grammar_file.Close()

	var info os.FileInfo
	info, err = grammar_file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	l := info.Size()
	buf := make([]byte, l)
	_, err = grammar_file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Parse out the token statements in the grammar file.
	tokens := Tokenize(buf)
	PrintlnSource(tokens)
}
