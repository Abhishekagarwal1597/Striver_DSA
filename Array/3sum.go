# hashmap solution
# Always insert data in hashmap in key not in value part

/*
No sort algo
Time - O(n2)
Space - O(n)
*/
func threeSum(nums []int) [][]int {
    res:=make(map[[3]int]bool)
    for i:=0;i<len(nums);i++ {
        seen:=make(map[int]bool)
        for j:=i+1;j<len(nums);j++ {
            // Check if the compliment exists in the hash map
            // If the compliment exists, put the result into a map to avoid duplicates in the result
            if seen[-(nums[i]+nums[j])] {
                temp:=[]int{nums[i],nums[j],-(nums[i]+nums[j])}
                // This sorting was required since the leetcode solution was expected to be sorted. This sorting takes O(1) since it sorts only 3 elements at a given point of time
                sort.Ints(temp)
                res[[3]int{temp[0],temp[1],temp[2]}]=true
            }
            // Mark the just visited elemented and seen
            seen[nums[j]]=true
        } 
    }
    
    // Convert the map answer constructed in the previous steps to a slice which is expected by the leetcode solution
    ans:=make([][]int, 0)
    for k,_:=range res {
        ans=append(ans, []int{k[0],k[1],k[2]})
    }
    return ans
}

# using 2 pointer approach
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    newMap := make([][]int, 0)
    for i:=0;i<len(nums);i++{
        if i!=0 && nums[i] == nums[i-1] {
            continue
        }
        right := len(nums) - 1
        left := i+1
        for right > left {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                newMap = append(newMap, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                
                //skip the duplicates
                for left < right && nums[left] == nums[left - 1]{
                    left++
                }
                for left < right && nums[right] == nums[right+ 1] {
                    right--
                }
                
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return newMap
}
