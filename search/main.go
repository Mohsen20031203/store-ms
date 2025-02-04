package main

import (
	"fmt"
	"math"
	"runtime"
	"sort"

	"golang.org/x/exp/rand"
)

func measureMemoryUsage() uint64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return memStats.Alloc
}

var cnt int

func LinearSearch(arr []int, target int) int {
	var count int
	for i, num := range arr {
		count++
		if num == target {
			return i // Found the target
		}
	}
	return -1 // Not found
}

func BinarySearch(arr []int, target int) int {
	var count int
	left, right := 0, len(arr)-1
	for left <= right {
		count++
		cnt++
		mid := left + (right-left)/2
		if arr[mid] == target {
			cnt = 0
			return mid // Found
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Not found
}

func JumpSearch(arr []int, target int) int {
	var count int

	add := int(math.Sqrt(float64(len(arr))))
	step := add
	prev := 0
	for arr[step-1] < target {
		count++
		prev = step
		step += add
		if prev >= len(arr) {
			return -1
		}
	}

	for i := prev; i < step; i++ {
		count++
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func InterpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		n1, n2, n3 := (target - arr[low]), (high - low), (arr[high] - arr[low])
		pos := low + (n1 * n2 / n3)

		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

func ExponentialSearch(arr []int, target int) int {
	if arr[0] == target {
		return 0
	}

	i := 1
	for i < len(arr) && arr[i] <= target {
		cnt++
		i *= 3
	}
	m := arr[:min(i, len(arr))]
	n := BinarySearch(m, target)
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FibonacciSearch(arr []int, target int) int {
	n := len(arr)

	// Find the smallest Fibonacci number greater than or equal to n
	fibM2 := 0            // (m-2)'th Fibonacci
	fibM1 := 1            // (m-1)'th Fibonacci
	fibM := fibM1 + fibM2 // m'th Fibonacci

	for fibM < n {
		fibM2, fibM1 = fibM1, fibM
		fibM = fibM1 + fibM2
	}

	offset := -1

	// Main loop for searching
	for fibM > 1 {
		i := min(offset+fibM2, n-1)

		if arr[i] < target {
			fibM, fibM1, fibM2 = fibM1, fibM2, fibM1-fibM2
			offset = i
		} else if arr[i] > target {
			fibM, fibM1, fibM2 = fibM2, fibM1-fibM2, fibM2-fibM1
		} else {
			return i
		}
	}

	if fibM1 == 1 && offset+1 < n && arr[offset+1] == target {
		return offset + 1
	}

	return -1
}

func TernarySearch(arr []int, left, right, target int) int {
	if right >= left {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			return TernarySearch(arr, left, mid1-1, target)
		} else if target > arr[mid2] {
			return TernarySearch(arr, mid2+1, right, target)
		} else {
			return TernarySearch(arr, mid1+1, mid2-1, target)
		}
	}
	return -1
}

func main() {

	arr := make([]int, 100000)
	for i := range arr {
		arr[i] = rand.Intn(1000000)
	}
	arr2 := make([]int, 100000)
	for i := range arr2 {
		arr2[i] = rand.Intn(1000000)
	}

	sort.Ints(arr2)
	ArraySort := arr2

	// Measure memory usage before executing searches
	memBefore := measureMemoryUsage()

	Binary := BinarySearch(ArraySort[:], 4620)
	if Binary != -1 {
		fmt.Println("right in BinarySort")
	}
	Line := LinearSearch(arr[:], 342725)
	if Line != -1 {
		fmt.Println("right in LinearSearch")
	}
	Jump := JumpSearch(ArraySort[:], 601202)
	if Jump != -1 {
		fmt.Println("right in JumpSearch")
	}
	Interpolation := InterpolationSearch(ArraySort[:], 601202)
	if Interpolation != -1 {
		fmt.Println("right in Interpolation")
	}
	Exponential := ExponentialSearch(ArraySort[:], 4620)
	if Exponential != -1 {
		fmt.Println("right in Exponential")
	}
	Fibonacci := FibonacciSearch(ArraySort, 499407)
	if Fibonacci != -1 {
		fmt.Println("right in Fibonacci")
	}
	Ternary := TernarySearch(ArraySort, 0, len(ArraySort), 700380)
	if Ternary != -1 {
		fmt.Println("right in Fibonacci")
	}

	// Measure memory usage after executing searches
	memAfter := measureMemoryUsage()
	fmt.Printf("Memory used: %d bytes\n", memAfter-memBefore)
}
