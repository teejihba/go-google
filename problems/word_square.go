package problems

import (
	"strings"
)

var size int
func GetWordSquare(words []string, ) [][]string{
	var ans [][]string
	size = len(words[0])
	for i := range words{
		recUtil(words, []string{words[i]}, &ans)
	}
	return ans
}

func recUtil(words []string, curr []string, ans *[][]string) {
	if len(curr) == size{
		*ans = append(*ans, curr)
		return
	}
	pref := ""
	currSize := len(curr)
	for i := range curr{
		pref = pref+ string(curr[i][currSize])
	}
	poss := SearchWithPrefix(words, pref)
	if len(poss) == 0{
		return
	}
	for i := range poss{
		c := make([]string, len(curr))
		copy(c, curr)
		newArr := append(c, poss[i])
		recUtil(words, newArr, ans)
	}
}

func SearchWithPrefix(words []string, prefix string) []string{
	var res []string
	for i := range words{
		if strings.HasPrefix(words[i], prefix){
			res = append(res, words[i])
		}
	}
	return res
}
