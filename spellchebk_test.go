package spellchebk

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"
)

func TestTrueDLDistance(t *testing.T) {
	tt := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "letters",
			word1:    "CA",
			word2:    "ABC",
			expected: 2,
		},
		{
			name:     "neat",
			word1:    "testing",
			word2:    "isneat",
			expected: 6,
		},
		{
			name:     "sports",
			word1:    "touchdown",
			word2:    "downtouch",
			expected: 8,
		},
		{
			name:     "number words",
			word1:    "seven",
			word2:    "eight",
			expected: 5,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := TrueDLDistance(tc.word1, tc.word2)
			if result != tc.expected {
				t.Errorf("expected %d but got %d for distance between %s and %s", tc.expected, result, tc.word1, tc.word2)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tree := BKtree{}
	tree.add("sail", TrueDLDistance)
	tree.add("mail", TrueDLDistance)
	tree.add("rail", TrueDLDistance)
	tree.add("snape", TrueDLDistance)

	// tree should look like this now
	//							sail
	//							/ \
	//						 1   3
	//						/			\
	//					mail 	 snape
	//					/
	//				 1
	//				/
	//			rail

	if tree.Word != "sail" {
		t.Error("expected \"sail\" to be the root")
	}
	if tree.Children[0].Word != "mail" {
		t.Error("expected \"mail\" to be the 0th child of the root")
	}
	if tree.Children[0].Children[0].Word != "rail" {
		t.Error("expected \"rail\" to be the 0th child of the 0th child of the root")
	}
	if tree.Children[1].Word != "snape" {
		t.Error("expected \"snape\" to be the 1st child of the root")
	}

	err := tree.add("snape", TrueDLDistance)
	if err != nil {
		t.Error("expected adding \"snape\", which is already in the tree, to throw an error")
	}
}

func TestSearch(t *testing.T) {
	c := NewSpellChecker()
	c.Add("sail")    // sail is the root
	c.Add("mail")    // mail has a dist of 1 (depending on distfunc)
	c.Add("rail")    // rail has a dist of 1 (depending on distfunc)
	c.Add("snape")   // snape has a dist of 3 (depending on distfunc)
	c.Add("far off") // far off has a dist of 6 (depending on distfunc)

	tt := []struct {
		query     string
		tolerance int
		expected  []SearchResult
	}{
		{
			query:     "nail",
			tolerance: 1,
			expected: []SearchResult{
				{
					Word:     "sail",
					Distance: 1,
				},
				{
					Word:     "mail",
					Distance: 1,
				},
				{
					Word:     "rail",
					Distance: 1,
				},
			},
		},
		{
			query:     "snape",
			tolerance: 0,
			expected: []SearchResult{
				{
					Word:     "snape",
					Distance: 0,
				},
			},
		},
		{
			query:     "rail",
			tolerance: 0,
			expected: []SearchResult{
				{
					Word:     "rail",
					Distance: 0,
				},
			},
		},
		{
			query:     "not here",
			tolerance: 0,
			expected:  nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.query, func(t *testing.T) {
			result := c.Search(tc.query, tc.tolerance)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("search for \"%s\" expected: %v got: %v", tc.query, tc.expected, result)
			}
		})
	}
}

func TestGob(t *testing.T) {
	encode := NewSpellChecker()
	encode.Add("test123")

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)

	err := e.Encode(encode)
	if err != nil {
		t.Errorf("failed to encode to gob: %s", err)
	}

	var decode *SpellChecker
	err = d.Decode(&decode)
	if err != nil {
		t.Errorf("failed to decode to gob: %s", err)
	}

	if !reflect.DeepEqual(encode.Tree, decode.Tree) {
		t.Errorf("decoded gob does not match original, expected: %v got: %v", encode, decode)
	}
}
