/*
CHRIS FELLI, 2019
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays.
The overall run time complexity should be O(log(m+n)).
You may assume nums1 and nums2 cannot be both empty.
*/
package findmediansortedarrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	got := FindMedianSortedArrays([]int{}, []int{1})
	want := float64(1)

	if got != want {
		t.Errorf("Got %.2f, wanted %.2f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %.2f, wanted %.2f\nPASSED!\n\n", got, want)
	}
}
func TestFindMedianSortedArrays2(t *testing.T) {
	got := FindMedianSortedArrays([]int{}, []int{1, 2})
	want := float64(1.5)

	if got != want {
		t.Errorf("Got %.2f, wanted %.2f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %.2f, wanted %.2f\nPASSED!\n\n", got, want)
	}
}
func TestFindMedianSortedArrays3(t *testing.T) {
	got := FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	want := float64(2.5)

	if got != want {
		t.Errorf("Got %.2f, wanted %.2f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %.2f, wanted %.2f\nPASSED!\n\n", got, want)
	}
}
func TestFindMedianSortedArrays4(t *testing.T) {
	got := FindMedianSortedArrays([]int{1, 2}, []int{3, 4, 5})
	want := float64(3)

	if got != want {
		t.Errorf("Got %.2f, wanted %.2f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %.2f, wanted %.2f\nPASSED!\n\n", got, want)
	}
}

func TestFindMedianSortedArrays5(t *testing.T) {
	got := FindMedianSortedArrays([]int{}, []int{})
	want := float64(0)

	if got != want {
		t.Errorf("Got %.2f, wanted %.2f\nFAILED!\n", got, want)
	} else {
		fmt.Printf("Got %.2f, wanted %.2f\nPASSED!\n\n", got, want)
	}
}
