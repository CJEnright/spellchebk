# spellchebk
Implementation of a spell checker (and corrector) using [BK-trees](https://en.wikipedia.org/wiki/BK-tree).

## Example
```
// Start the tree with an arbitrary word
tree := BKTree{Word: "cat"}

// And add words as needed
tree.Add("rat")
tree.Add("dog")

// And query the tree by giving it a string and a tolerence
res := tree.Search("eog", 1)

// res = [{Word: "dog", Distance: 1}]
```

lmao, hope I spelled everything here correctly.
