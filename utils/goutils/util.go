package goutils

func Init2DIntSlice(row, col int) *[][]int{
	a := make([][]int,row )
	for i := range a {
		a[i] = make([]int, col)
	}
	return &a
}
