/*
CHRIS FELLI, 2019
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays.
The overall run time complexity should be O(log(m+n)).
You may assume nums1 and nums2 cannot be both empty.
*/

// Package gocode ...
package gocode

import (
	"fmt"
	"sort"
)

// FindMedianSortedArrays : Finds median of two sorted arrays
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newArray := AppendAndSort(nums1, nums2)
	var arrLen = len(newArray)

	if arrLen == 0 {
		fmt.Println("Array length is zero")
		return 0
	} else if arrLen <= 1 {
		// Edge case: array length of 1
		fmt.Println("Median is ", float64(newArray[0]))
		return float64(newArray[0])
	} else if arrLen <= 2 {
		// Edge case: array length of 2
		var median = float64(newArray[0]+newArray[1]) / 2
		fmt.Println("Median is ", median)
		return median
	} else if arrLen%2 == 1 {
		// If odd
		var median = float64(newArray[(arrLen-1)/2])
		fmt.Println("Median is ", median)
		return median
	} else if arrLen%2 == 0 {
		// If even
		var medianA = float64(newArray[(arrLen/2)-1])
		var medianB = float64(newArray[(arrLen / 2)])
		var median = float64(medianA+medianB) / 2
		fmt.Println("Median is ", median)
		return median
	} else {
		fmt.Println("I have no idea how you got here.")
		return 0
	}
}

// AppendAndSort : Takes two int arrays, returns one sorted array
func AppendAndSort(nums1 []int, nums2 []int) []int {
	fmt.Println("Array 1 is ", nums1)
	fmt.Println("Array 2 is ", nums2)
	newArr := append(append([]int{}, nums1...), nums2...)
	sort.Ints(newArr)
	fmt.Println("Sorted array is ", newArr)
	return newArr
}
