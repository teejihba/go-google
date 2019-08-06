package problems

import "fmt"

/*
Problem Statement :
Given a mountain array of N elements [increasing then decreasing] e.g [1,5,9,20,12,8,6] ,
Find if a given number exists or not.
*/

type Input struct {
	Arr      []int
	Len      int
	GetCount int
}

//var Arr = [] int {1,5,9,20,24,12,8,6,4,3,2}

func (input *Input) GetValAtIndex(index int) int {
	input.GetCount++
	return input.Arr[index]
}

func (input *Input) FindIfExists(num int) bool {
	peakIndex := input.FindPeak(0, input.Len-1)
	fmt.Println("Peak index = ", peakIndex)
	left, _ := input.BinarySearch(0, peakIndex, num, func(i1 int, i2 int) bool {
		return i1 > i2
	})
	right, _ := input.BinarySearch(peakIndex+1, input.Len-1, num, func(i1 int, i2 int) bool {
		return i2 > i1
	})
	return left || right

}

func (input *Input) BinarySearch(start, end, num int, compare func(int, int) bool) (bool, int) {
	if start >= end {
		return input.GetValAtIndex(start) == num, start
	}
	mid := (start + end) / 2
	midVal := input.GetValAtIndex(mid)
	if midVal == num {
		return true, mid
	} else if compare(num, midVal) {
		return input.BinarySearch(mid+1, end, num, compare)
	} else {
		return input.BinarySearch(start, mid-1, num, compare)
	}

}

func (input *Input) FindPeak(start, end int) int {
	if start == end {
		return start
	}
	mid := (start + end) / 2
	midVal := input.GetValAtIndex(mid)
	midNextVal := input.GetValAtIndex(mid + 1)
	if midVal < midNextVal {
		return input.FindPeak(mid+1, end)
	}
	midPrevVal := input.GetValAtIndex(mid - 1)
	if midVal > midNextVal && midPrevVal > midVal {
		return input.FindPeak(start, mid-1)
	}
	if midVal > midNextVal && midVal > midPrevVal {
		return mid
	}
	return -1
}
