package main

import "fmt"

// Time Complexity: O(n log n) <=> O(N2)
// Auxiliary Space: O(1)

func partition(arr []int, low int, high int) int {
	pivot := arr[high] // pivot

	i := low - 1 // pointer for greater element

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i] // swap
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1] // move pivot to it's own place

	fmt.Printf("Array: %d\n", arr)

	return i + 1
}

func QuickSort(arr []int, low int, high int) {
	if low < high {

		pi := partition(arr, low, high)

		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func runQuickSort() {
	array := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Printf("Initial array: %d\n", array)

	QuickSort(array, 0, len(array)-1)

	fmt.Printf("Sorted array: %d\n", array)
}
