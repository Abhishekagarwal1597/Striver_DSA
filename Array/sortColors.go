func sortColors(nums []int)  {
    // sorting using counting sort
    //  for i:=0;i<len(nums);i++{
    //     for j:=0;j<len(nums)-1;j++{
    //         if nums[j] > nums[j+1]{
    //             nums[j], nums[j+1] = nums[j+1], nums[j]
    //         }
    //     }
    // }

    // use dutch national flag algorithm
    mid, low, high := 0, 0, len(nums)-1
    for mid <= high {
        if nums[mid] == 0{
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        }else if nums[mid] == 1{
            mid++
        } else{
            nums[mid], nums[high]  = nums[high], nums[mid]
            high--
        }
    }
}
