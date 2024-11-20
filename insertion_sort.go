package main

import "fmt"

// Time Complexity: O(N2)
// Auxiliary Space: O(1)

// [0],9,8,7,6,5,4,3,2,1
// [0],[9],8,7,6,5,4,3,2,1
// 0,[9],[8],7,6,5,4,3,2,1 => insert
// 0,[8],9,[7],6,5,4,3,2,1 => insert
// 0,[7],[8],9,[6],5,4,3,2,1 => insert

// func insert(arr []int, from int, to int) []int {
// 	val := arr[from]

// 	var withoutElement []int = append(arr[:from], arr[from+1:]...)

// 	var tail []int = append([]int{val}, withoutElement[to:]...)

// 	return append(arr[:to], tail...)
// }

func insert(slice []int, fromIndex int, toIndex int) []int {
	if fromIndex == toIndex {
		return slice
	}

	element := slice[fromIndex]

	if fromIndex < toIndex {
		copy(slice[fromIndex:], slice[fromIndex+1:toIndex+1])
	} else {
		copy(slice[toIndex+1:], slice[toIndex:fromIndex])
	}

	slice[toIndex] = element

	return slice
}

func InsertionSort(arr []int) {
	var length int = len(arr)

	// i => index of insertion value
	// j => index of insertion place
	for i := 0; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] && j == 0 {
				// insert to j
				arr = insert(arr, i, j)
			} else if arr[i] < arr[j] && arr[i] > arr[j-1] {
				// insert to j
				arr = insert(arr, i, j)
			}
			fmt.Printf("Array: %d\n", arr)
		}
	}
}

func BinaryInsertionSort(arr []int) {
	// TODO
}

func runInsertionSort() {
	array := []int{0, 9, 8, 7, 6, 1, 2, 3, 4, 5}

	fmt.Printf("Initial array: %d\n", array)

	InsertionSort(array)

	fmt.Printf("Sorted array: %d\n", array)
}
