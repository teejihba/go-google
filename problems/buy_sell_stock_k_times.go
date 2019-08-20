package problems

/*
Problem:
Say you have an array for which the i-th element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

Example:
Input: [3,2,6,5,0,3], k = 2
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
             Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
*/

func MaxProfitOnBuyingSellingStockAtmostKTimes(k int, prices []int) int {
	n := len(prices)
	if n==0 || n == 1 || k == 0{
		return 0
	}
	if k >= len(prices) / 2 {
		return quickSolve(k, prices)
	}
	prev := make([]int, n)
	cur := make([]int, n)

	for i:=1; i<k+1;i++{
		for j:=0;j<n;j++{
			if j==0{
				cur[j] = 0
				continue
			}
			cur[j]=cur[j-1]
			for x:=0; x<j; x++{
				p := prices[j] -prices[x]
				if prev[x] + p > cur[j]{
					cur[j] = prev[x] + p
				}
			}


		}
		copy(prev, cur)
	}
	return cur[n-1]

}
func quickSolve(k int, prices []int) int{
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			profit += prices[i] - prices[i - 1]
		}
	}

	return profit
}