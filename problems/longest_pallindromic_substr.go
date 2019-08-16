package problems

/*
Problem description :

Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.
Example 1 :
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
*/


func LongestPalindrome(str string) string{
	if len(str) == 0 || len(str) == 1{
		return str
	}
	max := -1
	start := -1
	for i := range str{
		tempOdd, startO := lpsUtilOddLength(i, str)
		if tempOdd > max{
			max = tempOdd
			start = startO
		}
		tempEven, startE := lpsutilEvenLength(i, str)
		if tempEven > max{
			max = tempEven
			start = startE
		}
	}
	return str[start:start+max]
}

func lpsUtilOddLength(index int , str string) (int, int){
	left, right, count := index -1, index +1, 1
	for left >= 0&& right<len(str) && str[left]==str[right] {
		count += 2
		left --
		right ++
	}
	return count, left+1
}

func lpsutilEvenLength(index int, str string) (int, int) {
	left , right, count := index, index+1, 0
	for left >= 0&& right<len(str) && str[left]==str[right] {
		count += 2
		left --
		right ++
	}

	return count, left+1
}
