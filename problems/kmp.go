package problems

import "fmt"

// return first occurence of pattern in text with its start index in text
// return true, startIndex if present, else false, -1
func KmpSubstrSearch(text, pattern string) (bool , int) {
	lps := *lengthOfLongestPrefixWhichIsAlsoSuffix(pattern)
	fmt.Println(lps)
	ti,pi := 0, 0
	for {
		if ti == len(text){
			break
		}
		if text[ti] == pattern[pi]{
			ti++
			pi++
			if pi == len(pattern){
				return true, ti - len(pattern)
			}
		}else{
			if pi == 0{
				ti++
			}else {
					pi = lps[pi-1]
			}
		}
	}
	return false, -1
}

func lengthOfLongestPrefixWhichIsAlsoSuffix(s string) *[] int {
	//lps[i] represents length of longest prefix in s[0:i] which is also a suffix in s[0:i]
	lps := make([]int, len(s))
	//start can be seen as length of already matched prefix with suffix
	start, end := 0,1
	//base case
	lps[0] = 0
	for {
		//loop break condition
		if end >= len(s){
			break
		}
		//if char at start and end are equal, length of matched prefix with suffix in s[0:i] can be incremented by 1
		//start and end both are incremented to match next character
		//check definition of start for clarification
		if s[start] == s[end]{
			lps[end] = start + 1
			start++
			end++
		} else {
				if start != 0{
					//Tricky part :
					//lps[start-1] represents length of matching prefix with suffix in s[0:start)
					//setting start = lps[start-1], means that now start is at location just after already matched prefix
					//we dont need to compare the already matched prefixes.
					//consider string "acabacacd"
					// start = 3, end = 7 and s[3] != s[7],
					//now, start will be set to lps[2]. lps[2] is 1, now, s[start] is equal to s[end] , so we set s[7] = 2,
					// we saved comparing the first a again.
					start = lps[start-1]
				}else {
					// if chars dont match and start = 0 , i.e length of already matched prefix is zero.
					// lps[end] has to be 0, nothing was matching previously and current chars also dont match
					// only end is incremented in hope that it matches with start
					lps[end] = 0
					end++
				}
		}

	}

	return &lps
}