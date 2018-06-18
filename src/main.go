package bkspellcheck

type BKTree struct {
	Word string
	Children []BKTree
	Distance int
}

type SearchResult struct {
	Word string
	Distance int
}

func (t BKTree) AddWord(input string) {
	dist := Distance(t.Word, input)

	for i := range t.Children {
		if t.Children[i].Distance == dist {
			t.Children[i].AddWord(input)
			return
		}
	}

	t.Children = append(t.Children, BKTree{Word: input, Distance: dist})
}

func (t BKTree) Search(word string, maxDist int) (found []SearchResult) {
	dist := Distance(t.Word, word)

	if t.Distance >= dist - maxDist && t.Distance <= dist + maxDist {
		found = append(found, SearchResult{Word: word, Distance: dist})
	}

	for i := range t.Children {
		found = append(found, t.Children[i].Search(word, maxDist)...)
	}

	return found
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Distance(word1 string, word2 string) int {
	// True Damerau–Levenshtein distance, but figured i'd keep the func name light

	// Size of our alphabet
	var da [255]int
	d := make([][]int, len(word1) + 2)
	for i := range d {
		d[i] = make([]int, len(word2) + 2)
	}

	max := len(word1) + len(word2)

	d[0][0] = max
	for i := 0; i <= len(word1); i++ {
		d[i + 1][0] = max
		d[i + 1][1] = i
	}
	for i := 0; i <= len(word2); i++ {
		d[0][i + 1] = max
		d[1][i + 1] = i
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
			ins := d[i + 1][j] + 1
			del := d[i][j + 1] + 1
			trans := d[k][l] + (i-k-1) + 1 + (j-l-1)

			d[i + 1][j + 1] = min(sub, min(ins, min(del, trans)))
		}

		da[word1[i-1]] = i
	}

	return d[len(word1) + 1][len(word2) + 1]
}
