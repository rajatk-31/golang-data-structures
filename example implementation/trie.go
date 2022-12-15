package main

import "fmt"

// Node
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// Create new trie
func InitTrie() *Trie {
	newTrie := &Trie{root: &Node{}}
	return newTrie
}

// Insert
func (t *Trie) Insert(w string) {
	wordLenght := len(w)
	current := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a' // This will give the exact index according to alphabets
		if current.children[charIndex] == nil {
			current.children[charIndex] = &Node{}
		}
		current = current.children[charIndex]
	}
	current.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wordLenght := len(w)
	current := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a' // This will give the exact index according to alphabets

		if current.children[charIndex] == nil {
			return false
		}
		current = current.children[charIndex]
	}
	if current.isEnd == true {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie)
	testTrie.Insert("ajay")
	testTrie.Insert("kumar")
	fmt.Println(testTrie.Search("ramesh"))
	fmt.Println(testTrie.Search("kuma"))
	fmt.Println(testTrie.Search("kumar"))
}
