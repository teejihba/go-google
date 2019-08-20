package problems

/*
Problem:
Given a string S and a string T, count the number of distinct subsequences of S which equals T.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters.
(ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

Example:
Input: S = "babgbag", T = "bag"
Output: 5
Explanation:

As shown below, there are 5 ways you can generate "bag" from S.
(The caret symbol ^ means the chosen letters)

babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^
*/


func CountDistinctSubsequences(s string, t string) int {
	slen := len(s)
	tlen := len(t)
	if tlen==0{
		return 1
	}
	if slen==0{
		return 0
	}
	dp := make([][]int, tlen+1)
	for i := 0; i< tlen+1;i++{
		dp[i] = make([]int, slen+1)
	}
	for i:=0; i<tlen+1; i++ {
		for j:=0 ; j< slen+1; j++{
			if i==0 {
				dp[i][j] = 1
				continue
			}
			if j==0{
				dp[i][j] = 0
				continue
			}
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1]+dp[i-1][j-1]
			}else{
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[tlen][slen]
}
