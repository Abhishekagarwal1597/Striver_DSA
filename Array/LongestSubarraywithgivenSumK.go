// using hashing and prefix sum
// work for -ve number as well
func longestSubstring(inputArr []int, requiredSum int){
	hashMap := make(map[int]int)
	var maxSlice []int
	var prefixSum int
	for i:=0;i<len(inputArr);i++{
		// calculate the prefix sum till index i
		prefixSum += inputArr[i]
		// if sum = k, update the maxLen
		if prefixSum == requiredSum {
			maxSlice = append(maxSlice, i+1)
		}

		// calculate the sum of remaining part i.e. x-k:
		//Calculate the length and update maxLen:
		if filterIndex, ok := hashMap[prefixSum-requiredSum]; ok{
			maxSlice = append(maxSlice, i-filterIndex)

		}
		//for edge case, for longest e.x. [2, 0, 0, 3]
		//if sum already exist don't update, because we need the longest subarray
		if _, ok := hashMap[prefixSum]; !ok{
			hashMap[prefixSum] = i
		}
	}
	fmt.Print(findMax(maxSlice))
}

func findMax(inputSlice []int)int{
	max := inputSlice[0]
	for _, v := range inputSlice{
		if v > max {
			max = v
		}
	}
	return max
}
