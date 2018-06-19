package main

import (
	"github.com/CJEnright/spellchebk"
	"fmt"
)

func main() {
	tree := spellchebk.BKTree{Word: "cat"}
	// Just a bunch of random but common words
	words := [...]string{"back", "top", "people", "had", "list", "name", "just", "over", "state", "year", "day", "into", "email", "health", "world", "next", "used", "work", "last", "most", "products", "music", "buy", "data", "make", "them", "should", "product", "system", "post", "city", "policy", "number", "such", "please", "available", "copyright", "support", "message", "after", "best", "software", "then", "good", "video", "well", "where", "info", "rights", "public", "books", "high", "school", "through", "each", "links", "review", "years", "order", "very", "privacy", "book", "items"}

	for i := range words {
		tree.Add(words[i])
	}

	fmt.Println(tree.Search("nmae", 1))
	// Prints "[{name 1}]", the added word to the input

	fmt.Println(tree.Search("bakc", 2))
	// Prints "[{back 1} {make 2}]", so back is the most likely attempted spelling
	// However, "make" is not far off, and may fit better depending on context
}
