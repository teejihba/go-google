package problems

import (
	"go-google/utils/goutils"
	"go-google/utils/mathsUtil"
)

func LongestPalindromicSubsequence(str string) int {
	size := len(str)
	if size==0||size==1{
		return size
	}
	table := *goutils.Init2DIntSlice(size,size)
	for l:= 0; l< size ; l++ {
		for i:= 0; i< size -l; i++{
			j:= i+l
			if j==i{
				table[i][j] = 1
			}else{
				if str[i] == str[j]{
					table[i][j] = table[i+1][j-1]+2
				}else{
					table[i][j] = mathsUtil.Max(table[i+1][j], table[i][j-1])
				}
			}

		}
	}
	return table[0][size-1]
}
