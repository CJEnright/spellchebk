/*
Package spellchebk implements a spell checker using BK-trees.

It also provides potentially useful functions like TrueDLDistance to find the True
Damerau Levenshtein Distance between two words.

Here's a little example that can read your OS's words file and query the tree
that generates.

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
*/
package spellchebk
