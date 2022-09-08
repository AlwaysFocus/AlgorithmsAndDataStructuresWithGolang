package main

import (
	"AlgorithmsAndDataStructures/HashTables"
	"fmt"
)

func main() {
	// Create a new hash table
	hashTable := HashTables.Init()

	// Our list of keys we'll use to test our hash table
	nameList := []string{
		"Eli",
		"Seth",
		"Liam",
		"Paten",
	}

	// Insert our keys into the hash table
	for _, n := range nameList {
		hashTable.Insert(n)
	}

	// Search for each key in the hash table - should all be true
	for _, n := range nameList {
		if hashTable.Search(n) {
			fmt.Printf("Found %s in the hash table\n", n)
		} else {
			fmt.Printf("Could not find %s in the hash table\n", n)
		}
	}

}
