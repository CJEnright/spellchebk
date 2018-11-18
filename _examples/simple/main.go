package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cjenright/spellchebk"
)

func main() {
	checker := spellchebk.NewSpellChecker()

	// Assumes we're running this on a unix type system
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		checker.Add(s.Text())
	}

	fmt.Println(checker.Search("nmae", 1))
	fmt.Println(checker.Search("cronfield", 2))
}
