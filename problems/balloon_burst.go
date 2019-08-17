package problems

import (
	"fmt"
	"go-google/utils/goutils"
	"go-google/utils/mathsUtil"
)

func MaxCoinsForBurstingBalloons(nums[] int) {
	count := len(nums)
	dp := *goutils.Init2DIntSlice(count, count)
	last := *goutils.Init2DIntSlice(count, count)

	for l := 0; l <= count; l++ {
		for i := 0; i < count-l; i++ {
			j := i + l
			//fmt.Println(i , j)
			if i == j {
				dp[i][j] = nums[i]
				last[i][j] = nums[i]
				continue
			}
			if j == i+1 {
				dp[i][j] = nums[i]*nums[j] + mathsUtil.Max(nums[i], nums[j])
				last[i][j] =  mathsUtil.Max(nums[i], nums[j])
				continue
			}
			max := -1
			for k := i+1 ; k < j ; k++{
				temp := last[i][k]*nums[k]*last[k][j]  + dp[i][k-1] +dp[k+1][j]
				if temp > max{
					max = temp
					last[i][j] = nums[k]
				}
			}
			dp[i][j] = max
		}

	}

	for x:= 0 ; x < count; x++{
		fmt.Println(dp[x])
	}
}