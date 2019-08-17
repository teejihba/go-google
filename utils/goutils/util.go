package goutils

func Init2DIntSlice(row, col int) *[][]int{
	a := make([][]int,row )
	for i := range a {
		a[i] = make([]int, col)
	}
	return &a
}


func Init2DStringSlice(row, col int) *[][]string{
	a := make([][]string,row )
	for i := range a {
		a[i] = make([]string, col)
	}
	return &a
}
