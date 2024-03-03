func twoSum(nums []int, target int) []int {
    // using hashing and prefix sum
    hashMap := make(map[int]int)
    for i:=0;i<len(nums);i++{   
        if index, ok := hashMap[target - nums[i]]; ok{
            return []int{i, index}
        }
        hashMap[nums[i]] = i
    }
    return []int{-1, -1}
