/*
package main

import (

	"fmt"
	"hash/fnv"

)

// Structure for each entry (key-value pair)

	type entry struct {
		key   string
		value string
		next  *entry // Points to the next entry in case of collision
	}

// Main structure for the map

	type MyMap struct {
		buckets []*entry // Array of buckets (each bucket is a linked list)
		size    int      // Size of the map (number of buckets)
	}

// Simple hash function to convert the key into an index

	func (m *MyMap) hash(key string) int {
		// Use the FNV-1a hash algorithm to generate the hash
		h := fnv.New32a()
		h.Write([]byte(key))
		// Return the index by taking the modulo of the hash with the number of buckets
		return int(h.Sum32()) % m.size
	}

// Create a new map with a specific size

	func NewMyMap(size int) *MyMap {
		return &MyMap{
			buckets: make([]*entry, size), // Initialize the buckets array
			size:    size,                 // Set the size of the map
		}
	}

// Add an entry to the map

	func (m *MyMap) Put(key, value string) {
		// Calculate the index using the hash function
		index := m.hash(key)

		// Create a new entry (key-value pair)
		newEntry := &entry{key: key, value: value}

		// Print the index and key for debugging purposes
		fmt.Printf("Key: %s, Hash Index: %d\n", key, index)

		// Check for collision in the bucket
		if m.buckets[index] == nil {
			// If the bucket is empty, add the entry
			m.buckets[index] = newEntry
		} else {
			// If the bucket already has entries, resolve collision using a linked list
			current := m.buckets[index]
			// Traverse the linked list to see if the key already exists
			for current != nil {
				if current.key == key {
					// If the key already exists, update the value
					current.value = value
					return
				}
				current = current.next
			}
			// If the key doesn't exist, add the new entry to the front of the linked list
			newEntry.next = m.buckets[index]
			m.buckets[index] = newEntry
		}
	}

// Retrieve a value based on the key

	func (m *MyMap) Get(key string) (string, bool) {
		// Calculate the index using the hash function
		index := m.hash(key)

		// Traverse the linked list in the bucket to find the key
		current := m.buckets[index]
		for current != nil {
			if current.key == key {
				// If the key is found, return the value and true
				return current.value, true
			}
			current = current.next
		}
		// If the key is not found, return an empty string and false
		return "", false
	}

	func main() {
		// Create a new map with 5 buckets (to encourage collisions)
		myMap := NewMyMap(5)

		// Add some key-value pairs (we'll see if any of the keys hash to the same index)
		myMap.Put("name", "John")
		myMap.Put("age", "30")
		myMap.Put("city", "New York")
		myMap.Put("nickname", "Johnny")
		myMap.Put("country", "USA")
		myMap.Put("john", "NY")
		myMap.Put("jane", "NY")
		myMap.Put("julia", "NY")
		myMap.Put("jim", "NY")
		myMap.Put("josh", "NY")
		myMap.Put("johnny", "NY")
		myMap.Put("jerry", "NY")

		// Retrieve values based on keys
		name, found := myMap.Get("name")
		if found {
			fmt.Println("Name:", name) // Output: Name: John
		}

		city, found := myMap.Get("city")
		if found {
			fmt.Println("City:", city) // Output: City: New York
		}

		// Attempt to retrieve a non-existent key
		_, found = myMap.Get("address")
		if !found {
			fmt.Println("Key 'address' not found")
		}
	}
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

// printMemUsage prints memory usage
func printMemUsage(msg string) uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s -> Alloc = %v KB\n", msg, m.Alloc/1024)
	return m.Alloc
}

func checkMap() {
	start := time.Now()

	m := make(map[int]int)

	prevMem := printMemUsage("Initial memory")

	// Insert elements and detect when memory allocation increases significantly
	for i := 0; i < 1000; i++ {
		m[i] = i

		// Check memory every 100 elements
		if i%100 == 0 {
			currMem := printMemUsage(fmt.Sprintf("After inserting %d elements", i))
			if currMem > prevMem+1024 { // Detect significant memory jumps (1MB+)
				fmt.Printf("⚡ Resize happened at %d elements!\n", i)
				prevMem = currMem
			}
		}
	}

	printMemUsage("Final memory usage")

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
func checkMap3() {
	m := make(map[int]int)

	prevMem := printMemUsage("Initial memory")

	// Insert elements and detect when memory allocation increases significantly
	for i := 0; i < 1000; i++ {
		m[i] = i

		// Check memory every 100 elements
		if i%100 == 0 {
			currMem := printMemUsage(fmt.Sprintf("After inserting %d elements", i))
			if currMem > prevMem+1024 { // Detect significant memory jumps (1MB+)
				fmt.Printf("⚡ Resize happened at %d elements!\n", i)
				prevMem = currMem
			}
		}
	}

	printMemUsage("Final memory usage")
}

func checkMap2() {

	start := time.Now()

	m := make(map[int]int)

	prevMem := printMemUsage("Initial memory")

	// Insert elements and detect when memory allocation increases significantly
	for i := 0; i < 300; i++ {
		m[i] = i

		// Check memory every 100 elements
		if i%100 == 0 {
			currMem := printMemUsage(fmt.Sprintf("After inserting %d elements", i))
			if currMem > prevMem+1024 { // Detect significant memory jumps (1MB+)
				fmt.Printf("⚡ Resiiiiize happened at %d elements!\n", i)
				prevMem = currMem
			}
		}
	}

	printMemUsage("Final memory usage")

	// Create a new map for remaining elements
	m2 := make(map[int]int)

	// Continue inserting remaining elements in the new map
	for i := 300; i < 1000; i++ {
		m2[i] = i

		if i%100 == 0 {
			currMem := printMemUsage(fmt.Sprintf("After inserting %d elements", i))
			if currMem > prevMem+1024 { // Detect significant memory jumps (1MB+)
				fmt.Printf("⚡ Resize happened at %d elements!\n", i)
				prevMem = currMem
			}
		}
	}
	elapsed := time.Since(start)

	fmt.Println(elapsed)
	printMemUsage("Final memory usage")
}

func main() {
	//checkMap()
	//checkMap3()
	checkMap2()
}
