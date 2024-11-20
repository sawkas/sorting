package main

import "fmt"

// Time Complexity: O(n log n)
// Auxiliary Space: O(n)

func merge(arr []int, left int, mid int, right int) {
	var leftSize int = mid - left + 1
	var rightSize int = right - mid

	var leftArray = make([]int, leftSize)
	var rightArray = make([]int, rightSize)

	copy(leftArray, arr[left:mid+1])
	copy(rightArray, arr[mid+1:right+1])

	i := 0    // index of first array
	j := 0    // index of second array
	k := left // index of original array

	for i < leftSize && j < rightSize {

		if leftArray[i] < rightArray[j] {
			arr[k] = leftArray[i]
			i += 1
		} else {
			arr[k] = rightArray[j]
			j += 1
		}

		k += 1
	}

	// copy the rest of leftArray
	for i < leftSize {
		arr[k] = leftArray[i]
		i += 1
		k += 1
	}

	// copy the rest of rightArray
	for j < rightSize {
		arr[k] = rightArray[j]
		j += 1
		k += 1
	}

	fmt.Printf("Array: %d\n", arr)
}

func MergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	var mid int = left + (right-left)/2

	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func runMergeSort() {
	array := []int{0, 9, 8, 7, 6, 1, 2, 3, 4, 5}

	fmt.Printf("Initial array: %d\n", array)

	MergeSort(array, 0, len(array)-1)

	fmt.Printf("Sorted array: %d\n", array)
}
