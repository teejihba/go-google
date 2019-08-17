package problems
/*
Problem :
Given an array A of positive integers,
call a (contiguous, not necessarily distinct) subarray of A good,
if the number of different integers in that subarray is exactly K.

(For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.)
Return the number of good subarrays of A.

Example :
Input: A = [1,2,1,3,4], K = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].


*/

func SubarrayWithKDistinctKDistinct(A []int , k int) int {
	left, right := 0, 0
	table := make(map[int]int)
	//table[A[0]] = 1
	count := 0
	for  {
		for len(table) <= k && right < len(A){
			if val, ok := table[A[right]] ; ok{
				table[A[right]] = val+1
			} else{
				table[A[right]] = 1
			}
			if len(table) == k{
				count++
			}
			right++
			if right == len(A){
				break
			}
		}
		for len(table) >= k || right == len(A) {
			if val, ok := table[A[left]] ; ok{
				if val > 1 {
					table[A[left]] = val-1
				}else{
					delete(table, A[left])
				}
			}else {
				panic("chaata re")
			}
			if len(table) == k{
				count++
			}
			left++
			if left == len(A){
				break
			}
		}
		if left ==right{
			break
		}


	}
	return count
}
