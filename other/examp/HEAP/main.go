package main

import "fmt"

// MinHeap struct
type MinHeap struct {
	arr []int
}

// Get the index of the parent, left child, and right child
func parent(i int) int     { return (i - 1) / 2 }
func leftChild(i int) int  { return 2*i + 1 }
func rightChild(i int) int { return 2*i + 2 }

// Insert a new element and maintain heap property
func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val) // Add to the end
	h.upHeap(len(h.arr) - 1)   // Fix heap property
}

// upHeap (Heapify Up) - Move element up to maintain the Min-Heap property
func (h *MinHeap) upHeap(i int) {
	for i > 0 && h.arr[i] < h.arr[parent(i)] {
		h.arr[i], h.arr[parent(i)] = h.arr[parent(i)], h.arr[i]
		i = parent(i)
	}
}

// ExtractMin - Removes and returns the smallest element
func (h *MinHeap) ExtractMin() int {
	if len(h.arr) == 0 {
		panic("Heap is empty")
	}

	min := h.arr[0]                // The root (smallest element)
	h.arr[0] = h.arr[len(h.arr)-1] // Move last element to root
	h.arr = h.arr[:len(h.arr)-1]   // Remove last element
	h.downHeap(0)                  // Fix heap property

	return min
}

// downHeap (Heapify Down) - Move element down to maintain the Min-Heap property
func (h *MinHeap) downHeap(i int) {
	lastIndex := len(h.arr) - 1
	for {
		l, r := leftChild(i), rightChild(i)
		smallest := i

		if l <= lastIndex && h.arr[l] < h.arr[smallest] {
			smallest = l
		}
		if r <= lastIndex && h.arr[r] < h.arr[smallest] {
			smallest = r
		}

		if smallest == i {
			break
		}

		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		i = smallest
	}
}

// Heapify function to maintain the heap property
func heapify(arr []int, n, i int) {
	largest := i     // Assume the root is the largest
	left := 2*i + 1  // Left child index
	right := 2*i + 2 // Right child index

	// If the left child is larger than the root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If the right child is larger than the largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest is not the root, swap and continue heapifying
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest) // Recursively heapify the affected subtree
	}
}

// Function to build a max heap from an array
func buildHeap(arr []int, n int) {
	// Start from the last non-leaf node and heapify each one
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func main() {
	h := &MinHeap{}

	arr := []int{1, 4, 3, 8, 6, 9, 2}

	buildHeap(arr, len(arr))

	// Insert elements
	for _, a := range arr {
		h.Insert(a)
	}

	fmt.Println("Extracted Min:", h.ExtractMin()) // Should print 1
	fmt.Println("Extracted Min:", h.ExtractMin()) // Should print 2
	fmt.Println("Heap after extraction:", h.arr)  // Remaining elements
}
