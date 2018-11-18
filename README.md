# spellchebk
Implementation of a spell checker (and corrector) using [BK-trees](https://en.wikipedia.org/wiki/BK-tree).

## Example
```
checker := NewSpellChecker()

// And add words as needed
checker.Add("rat")
checker.Add("dog")

// And query the checker by giving it a string and a tolerance
res := checker.Search("eog", 1)

// res = [{Word: "dog", Distance: 1}]
```

lmao, hope I spelled everything here correctly.
