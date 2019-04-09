package spellchebk

import (
	"fmt"
	"sync"
)

// A BKtree is an n-ary BK-tree.
type BKtree struct {
	Word     string    `json:"word"`
	Children []*BKtree `json:"children"`

	// Distance from parent to this node
	Distance int `json:"distance"`
}

// A SearchResult is one of possibly many results found during a search.
type SearchResult struct {
	Word     string `json:"word"`
	Distance int    `json:"distance"`
}

// A DistFunc returns an integer of how different two string are.
type DistFunc func(first, second string) (distance int)

// A SpellChecker builds and queries a spell checker to find similar words.
//
// It uses a string distance function to build an internal BK-tree.
// The default distance function is the True Damerau–Levenshtein distance which
// this package implements as TrueDLDistance.
// SpellCheckers are safe to use across threads.
type SpellChecker struct {
	mu   sync.Mutex
	Tree *BKtree
	// DistanceFunc will be used to find the distance between two words.
	// For every word added or query, this will likely be run several times.
	DistanceFunc DistFunc
}

// NewSpellChecker returns a new BKtree with the initial node root.
func NewSpellChecker() *SpellChecker {
	return &SpellChecker{
		DistanceFunc: TrueDLDistance,
		Tree:         &BKtree{},
	}
}

// Add inserts a word into the spell checker.
// It is safe to be used across threads.
func (s *SpellChecker) Add(input string) (err error) {
	if input == "" {
		return fmt.Errorf("attempted to add empty string")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Tree.add(input, s.DistanceFunc)
}

// add recursively traverses a BKtree to find a suitable spot for the given input.
// When it reaches across an empty string, the input will be placed there.
// For that reason (and from a practicallity standpoint), an empty string cannot be added to the tree.
func (b *BKtree) add(input string, df DistFunc) (err error) {
	if b.Word == "" {
		b.Word = input
		return err
	}

	dist := df(b.Word, input)
	if dist == 0 {
		// Distance == 0 implies that the word is already in the tree
		return fmt.Errorf("word %s already exists in spell checker", input)
	}

	for i := range b.Children {
		if b.Children[i].Distance == dist {
			return b.Children[i].add(input, df)
		}
	}

	b.Children = append(b.Children, &BKtree{Word: input, Distance: dist})
	return err
}

// Search queries the tree and returns SearchResult structs of matching words.
func (s *SpellChecker) Search(query string, tolerance int) (found []SearchResult) {
	return s.Tree.search(query, tolerance, s.DistanceFunc)
}

// search recursively traverses a BKtree.
// It returns a slice of SearchResults that fall within the given tolerance.
func (b *BKtree) search(query string, tolerance int, df DistFunc) (found []SearchResult) {
	dist := df(b.Word, query)

	if dist <= tolerance {
		found = append(found, SearchResult{Word: b.Word, Distance: dist})
	}

	for i := range b.Children {
		if b.Children[i].Distance >= dist-tolerance && b.Children[i].Distance <= dist+tolerance {
			found = append(found, b.Children[i].search(query, tolerance, df)...)
		}
	}

	return found
}

// min returns the minimun of two ints.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// TrueDLDistance finds the True Damerau–Levenshtein distance between two words.
// It's similar to the Damerau-Levenshtein distance but it also accounts for adjacent transpositions.
// Importantly, it also needs to know the size of the input alphabet.
func TrueDLDistance(first string, second string) int {
	// Size of our alphabet
	var da [255]int

	d := make([][]int, len(first)+2)
	for i := range d {
		d[i] = make([]int, len(second)+2)
	}

	max := len(first) + len(second)

	d[0][0] = max
	for i := 0; i <= len(first); i++ {
		d[i+1][0] = max
		d[i+1][1] = i
	}
	for i := 0; i <= len(second); i++ {
		d[0][i+1] = max
		d[1][i+1] = i
	}

	for i := 1; i <= len(first); i++ {
		db := 0
		for j := 1; j <= len(second); j++ {
			k := da[second[j-1]]
			l := db
			cost := 0

			if first[i-1] == second[j-1] {
				db = j
			} else {
				cost = 1
			}

			sub := d[i][j] + cost
			ins := d[i+1][j] + 1
			del := d[i][j+1] + 1
			trans := d[k][l] + (i - k - 1) + 1 + (j - l - 1)

			d[i+1][j+1] = min(sub, min(ins, min(del, trans)))
		}

		da[first[i-1]] = i
	}

	return d[len(first)+1][len(second)+1]
}
