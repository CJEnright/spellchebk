// Package spellchebk implements a spell checker using BK-trees.
package spellchebk

import "fmt"

// bktree is an n-ary BK-tree.
type bktree struct {
	word     string    `json:"word"`
	children []*bktree `json:"children"`
	// Distance from parent to this node
	distance int `json:"distance"`
}

// SearchResult is returned by the search function.
type SearchResult struct {
	Word     string `json:"word"`
	Distance int    `json:"distance"`
}

type distFunc func(word1, word2 string) (distance int)

type SpellChecker struct {
	tree         *bktree
	DistanceFunc distFunc
}

// NewSpellChecker returns a new bktree with the initial node root.
func NewSpellChecker() *SpellChecker {
	return &SpellChecker{
		DistanceFunc: TrueDamerauLevenshteinDistance,
		tree:         &bktree{},
	}
}

// Add inserts a word into the spell checker.
func (s *SpellChecker) Add(input string) (err error) {
	if input == "" {
		return fmt.Errorf("attempted to add empty string")
	}
	return s.tree.add(input, s.DistanceFunc)
}

// add recursively traverses a bktree to find a suitable spot for the given input.
// When it reaches across an empty string, the input will be placed there.
// For that reason (and from a practicallity standpoint), an empty string cannot be added to the tree.
func (b *bktree) add(input string, df distFunc) (err error) {
	if b.word == "" {
		b.word = input
		return err
	}

	dist := df(b.word, input)

	for i := range b.children {
		if b.children[i].distance == dist {
			b.children[i].add(input, df)
			return err
		}
	}

	// Distance == 0 implies that the word is already in the tree
	if dist == 0 {
		return fmt.Errorf("word %s already exists in spell checker", input)
	} else {
		b.children = append(b.children, &bktree{word: input, distance: dist})
	}

	return err
}

// Search queries the tree and returns SearchResult structs of matching words.
func (s *SpellChecker) Search(query string, tolerance int) (found []SearchResult) {
	return s.tree.search(query, tolerance, s.DistanceFunc)
}

// search recursively traverses a bktree.
// It returns a slice of SearchResults that fall within the given tolerance.
func (b *bktree) search(query string, tolerance int, df distFunc) (found []SearchResult) {
	dist := df(b.word, query)

	if dist <= tolerance {
		found = append(found, SearchResult{Word: b.word, Distance: dist})
	}

	for i := range b.children {
		if b.children[i].distance >= dist-tolerance && b.children[i].distance <= dist+tolerance {
			found = append(found, b.children[i].search(query, tolerance, df)...)
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

// TrueDamerauLevenshtein finds the True Damerau–Levenshtein distance between two words.
// It's similar to DamerauLevenshteinDistance but it also accounts for adjacent transpositions.
// Importantly, it also needs to know the size of the input alphabet.
func TrueDamerauLevenshteinDistance(word1 string, word2 string) int {
	// Size of our alphabet
	var da [255]int

	d := make([][]int, len(word1)+2)
	for i := range d {
		d[i] = make([]int, len(word2)+2)
	}

	max := len(word1) + len(word2)

	d[0][0] = max
	for i := 0; i <= len(word1); i++ {
		d[i+1][0] = max
		d[i+1][1] = i
	}
	for i := 0; i <= len(word2); i++ {
		d[0][i+1] = max
		d[1][i+1] = i
	}

	for i := 1; i <= len(word1); i++ {
		db := 0
		for j := 1; j <= len(word2); j++ {
			k := da[word2[j-1]]
			l := db
			cost := 0

			if word1[i-1] == word2[j-1] {
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

		da[word1[i-1]] = i
	}

	return d[len(word1)+1][len(word2)+1]
}
