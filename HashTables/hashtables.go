package HashTables

import "fmt"

const ArraySize = 10

// HashTable - hash table struct - array of linked lists
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket - bucket struct - linked list
type Bucket struct {
	head *BucketNode
}

// BucketNode - bucket node struct - linked list node
type BucketNode struct {
	key  string
	next *BucketNode
}

// Search - Search for a key in the hash table. Returns true if found, false if not
func (h *HashTable) Search(k string) bool {
	// Call our hash function to generate an index to search for the key
	index := hash(k)

	return h.array[index].Search(k)
}

// Insert - Insert a key into the hash table
func (h *HashTable) Insert(k string) {
	// Call our hash function to generate an index for the new key
	index := hash(k)

	h.array[index].Insert(k)
}

// Delete - Delete a key from the hash table
func (h *HashTable) Delete(k string) {
	// Call our hash function to generate an index to delete the key
	index := hash(k)

	h.array[index].Delete(k)
}

// Search - Search for a key in the bucket
func (b *Bucket) Search(k string) bool {
	// Create a pointer to the head of the bucket
	currentNode := b.head

	// Go to each node in the linked list to see if we can find the key
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		// Grab the next node to search
		currentNode = currentNode.next
	}
	return false
}

// Insert - Insert a key into the bucket
// accepts a key, then creates a new bucket node with the key
func (b *Bucket) Insert(k string) {
	// Check to see if the key already exists in the bucket
	if !b.Search(k) {

		// Create the new bucket node with the key
		newNode := &BucketNode{key: k}
		// Set the next node to be the old head of the bucket
		newNode.next = b.head
		// Set the new node to be the head of the bucket
		b.head = newNode
	} else {
		// Key already exists in the bucket
		fmt.Println("Key: `", k, "` already exists in the bucket")
	}
}

// Delete - Delete a key from the bucket
func (b *Bucket) Delete(k string) bool {
	// Deal with the case where the key is the head of the bucket
	if b.head.key == k {
		b.head = b.head.next
		return true
	}

	// We will be deleting the current node by skipping the current node
	// and pointing the previous node to the next node.
	// We will need to store each node in the previous node and express
	// the current node as the previousNode.next node, so we can keep track of things
	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == k {
			// We found the key to delete
			// Point the previous node to the next node
			previousNode.next = previousNode.next.next
			return true
		}

		previousNode = previousNode.next
	}

	// We did not find the key to delete
	return false
}

// Hash - Hash function to generate a hash for a key by getting
// the ascii value of each char, sum it, then mod with the array size
func hash(k string) int {
	sum := 0
	// Loop through each char in the key and sum the ascii values
	for _, v := range k {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init - Initialize the hash table
func Init() *HashTable {
	// Create a new hash table
	result := &HashTable{}

	// Iterate over the array and initialize a new bucket for each index
	for i := range result.array {
		result.array[i] = &Bucket{}
	}

	return result
}
