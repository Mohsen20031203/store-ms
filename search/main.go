package main

import (
	"fmt"
	"math"
)

func LinearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // Found the target
		}
	}
	return -1 // Not found
}

func BinarySearch(arr []int, target int) int {
	var cnt int
	left, right := 0, len(arr)-1
	for left <= right {
		cnt++
		mid := left + (right-left)/2
		if arr[mid] == target {
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
		i *= 2
	}
	m := arr[:min(i, len(arr))]
	n := BinarySearch(m, target)
	return n
}

func main() {

	ArrayUnsort := [100]int{23, 89, 45, 12, 77, 90, 33, 56, 98, 34,
		67, 21, 5, 92, 41, 38, 74, 60, 83, 15,
		11, 95, 6, 50, 27, 36, 69, 82, 49, 18,
		64, 29, 93, 10, 97, 87, 72, 4, 58, 46,
		70, 25, 55, 30, 16, 79, 88, 8, 39, 47,
		62, 59, 78, 2, 20, 91, 43, 85, 32, 31,
		53, 66, 13, 40, 86, 35, 99, 37, 48, 81,
		22, 3, 80, 9, 52, 54, 28, 1, 63, 57,
		68, 17, 61, 7, 24, 75, 76, 44, 71, 26,
		42, 14, 65, 51, 96, 19, 84, 73, 94, 100,
	}

	ArraySort := [100]int{
		2, 5, 9, 11, 14, 19, 23, 26, 30, 34,
		37, 41, 45, 48, 53, 57, 60, 63, 67, 70,
		74, 78, 81, 84, 88, 91, 94, 97, 101, 105,
		108, 112, 115, 119, 123, 126, 130, 134, 137, 140,
		144, 148, 151, 154, 158, 161, 165, 168, 172, 176,
		179, 183, 187, 190, 193, 197, 200, 204, 208, 211,
		214, 218, 221, 225, 229, 232, 236, 239, 242, 246,
		250, 253, 257, 261, 264, 268, 271, 275, 279, 282,
		286, 289, 293, 296, 300, 304, 307, 311, 315, 318,
		322, 325, 329, 332, 336, 339, 343, 347, 350, 354,
	}

	Binary := BinarySearch(ArraySort[:], 26)
	if Binary != -1 {
		fmt.Println("right in BinarySort")
	}
	Line := LinearSearch(ArrayUnsort[:], 5)
	if Line != -1 {
		fmt.Println("right in LinearSearch")
	}
	Jump := JumpSearch(ArraySort[:], 158)
	if Jump != -1 {
		fmt.Println("right in JumpSearch")
	}
	Interpolation := InterpolationSearch(ArraySort[:], 311)
	if Interpolation != -1 {
		fmt.Println("right in Interpolation")
	}
	Exponential := ExponentialSearch(ArraySort[:], 26)
	if Exponential != -1 {
		fmt.Println("right in Exponential")
	}
}
