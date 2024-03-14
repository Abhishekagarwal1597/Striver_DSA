# Using Recursion and Backtracking
func permute(nums []int) [][]int {
    var newArray [][]int
    findAllPermutation(&newArray, nums, 0, len(nums)-1)
    return newArray
}

func findAllPermutation(res *[][]int, nums []int, l, h int){
    if l == h {
         // Append a copy of nums to res, as slice is refrence type
         // the next swap will revert our changes
        perm := make([]int, len(nums))
        copy(perm, nums)
        *res = append(*res, perm)
        return 
    }

    for i:=l; i<=h; i++{
        //swaping
        nums[l], nums[i] = nums[i], nums[l]
        findAllPermutation(res, nums, l+1, h)
        //backtrack
        nums[l], nums[i] = nums[i], nums[l]
    }
}
