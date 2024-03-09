// using kadane Algorithm
// find max subarray surm in an array, if number is negative discard it.
package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	sum, max := 0, math.MinInt64
	// start, end := 0, 0
	for i := 0; i < len(nums); i++ {

		sum += nums[i]
		if sum > max {
			max = sum
		}

		// If sum < 0: discard the sum calculated
		if sum < 0 {
			sum = 0
		}

	}

	return max
}

func main() {
	var nums []int = []int{1, 2, -1, 1, 2}
	res := maxSubArray(nums)
	fmt.Print(res)
}
