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
	checker.DistanceFunc = myDist

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

	fmt.Println(checker.Search("fishery", 1))
	fmt.Println(checker.Search("cronfield", 2))
}

// dist will find the distance between two words.
// It does a really bad job of this, and there are many better ways to do this.
// For example, spellchebk.TrueDLDistance
func myDist(first, second string) (distance int) {
	// Find the shortest word
	shortestlen := len(first)
	distance = len(second) - shortestlen
	if len(second) < len(first) {
		shortestlen = len(second)
		distance = len(first) - shortestlen
	}

	// Distance increases for each differing letter
	for i := 0; i < shortestlen; i++ {
		if first[i] != second[i] {
			distance++
		}
	}

	return distance
}
