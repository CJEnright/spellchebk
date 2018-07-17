package spellchebk

import "testing"

func TestDistance(t *testing.T) {
	if dist := Distance("CA", "ABC"); dist != 2 {
		t.Error("Distance between \"CA\" and \"ABC\" expected", 2, "but was", dist)
	}
	if dist := Distance("testing", "isneat"); dist != 6 {
		t.Error("Distance between \"testing\" and \"isneat\" expected", 6, "but was", dist)
	}
	if dist := Distance("touchdown", "downtouch"); dist != 8 {
		t.Error("Distance between \"touchdown\" and \"downtouch\" expected", 8, "but was", dist)
	}
	if dist := Distance("seven", "eight"); dist != 5 {
		t.Error("Distance between \"seven\" and \"eight\" expected", 5, "but was", dist)
	}
}

func TestAdd(t *testing.T) {
	checker := NewSpellChecker("snail")
	checker.Add("sail")
	checker.Add("mail")
	checker.Add("snape")

	if checker.Children[0].Word != "sail" {
		t.Error("\"sail\" should be the 0th child of the root, but wasn't found")
	}
	if checker.Children[1].Word != "mail" {
		t.Error("\"mail\" should be the 1st child of the root, but wasn't found")
	}
	if checker.Children[1].Children[0].Word != "snape" {
		t.Error("\"snape\" should be the 0th child of the 1st child of the root, but wasn't found")
	}
}

func TestSearch(t *testing.T) {
	checker := NewSpellChecker("snail")
	checker.Add("sail")
	checker.Add("mail")
	checker.Add("snape")
	checker.Add("far off")

	res := checker.Search("nail", 1)
	if res[0].Word != "snail" && res[1].Word != "mail" {
		t.Error("Searching for \"nail\" did not return \"snail\" and \"mail\"")
	}

	res = checker.Search("snape", 0)
	if res[0].Word != "snape" {
		t.Error("Exact search for \"snape\" did not find \"snape\"")
	}

	res = checker.Search("not here", 0)
	if len(res) != 0 {
		t.Error("Search for non existant string yeilded a result")
	}
}
