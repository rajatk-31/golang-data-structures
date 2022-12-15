package main

import "fmt"

const ArraySize = 7

// Hash table
type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// Hash
func hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}
	return sum % ArraySize
}

// Insert in Hash Table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Insert in bucket
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Value already exists")
	}
}

// Search in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// search in bucket
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		} else {
			currentNode = currentNode.next
		}
	}
	return false
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// delete in bucket
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	currentNode := b.head
	for currentNode.next != nil {
		if currentNode.next.key == key {
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next

	}
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)
	list := []string{
		"Test",
		"Randy",
		"Jammy",
		"Flora",
		"Kyle",
		"Messi",
		"Me",
		"Anything",
	}
	for _, val := range list {
		testHashTable.Insert(val)
	}

	fmt.Println(testHashTable.Search("Messi"))
	fmt.Println(testHashTable.Search("MR"))
	fmt.Println(testHashTable.Search("usg"))
}
