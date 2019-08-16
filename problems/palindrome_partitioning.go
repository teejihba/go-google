package problems

import (
	"go-google/utils/goutils"
	"go-google/utils/mathsUtil"
)

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example:
Input: "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/


//Returns minimum number of cuts needed
func MinCutPalindromePartitioning(str string) int {
	size := len(str)
	if size==0||size==1{
		return 0
	}
	table := *goutils.Init2DIntSlice(size,size)
	for l:= 0; l< size ; l++ {
		for i:= 0; i< size -l; i++{
			j:= i+l
			if j==i{
				table[i][j] = 0
			}else if j==i+1{
				if str[i] == str[j]{
					table[i][j] = 0
				}else {
					table[i][j] = 1
				}
			} else {
				if isPal(str, i,j) {
					table[i][j] = 0
				}else{
					min := 1000000
					for k:= i; k <j ; k++{
						temp := table[i][k] + table[k+1][j] + 1
						min =  mathsUtil.Min(temp,min)
					}
					table[i][j] = min
				}
			}


		}
	}
	return table[0][size-1]
}

func isPal(str string , i,j int) bool {
	for str[i] == str[j]  {
		i++
		j--
		if i >= j {
			break
		}
	}
	return i>= j
}