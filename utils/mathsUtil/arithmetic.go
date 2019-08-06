package mathsUtil

import "math"

func Max(i1 int, i2 int) int {
	if i1>=i2 {
		return i1
	}
	return i2
}
func Min(i1 int, i2 int) int {
	if i1>=i2 {
		return i2
	}
	return i1
}

func VariadicMin(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		min = Min(min, num)
	}
	return min
}

func VariadicMax(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		max = Max(max, num)
	}
	return max
}

func AbsoluteVal(x int) int{
	if x >= 0{
		return x
	}
	return -x
}
