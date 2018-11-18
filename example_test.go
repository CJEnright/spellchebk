package spellchebk_test

import (
	"fmt"

	"github.com/cjenright/spellchebk"
)

func ExampleTrueDLDistance() {
	fmt.Println(spellchebk.TrueDLDistance("hello!", "goodbye!"))
	// Output: 7
}

func ExampleAdd() {
	checker := spellchebk.NewSpellChecker()
	checker.Add("hello")
	checker.Add("mellow")

	result := checker.Search("yellow", 1)
	fmt.Println(result[0].Word)
	// Output: mellow
}

func ExampleSearch() {
	checker := spellchebk.NewSpellChecker()
	checker.Add("hello")
	checker.Add("mellow")

	result := checker.Search("cello", 1)
	fmt.Println(result[0].Word)
	// Output: hello
}
