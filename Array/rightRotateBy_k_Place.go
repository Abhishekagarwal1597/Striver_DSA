#move the last k element at starting..and remaining at the end
#create a new array
#at last copy it to the orignal array

func rotate(nums []int, k int)  {
    newArr := make([]int, len(nums))
    l := len(nums)
    if l == 1{
        return 
    }

    var rotation int

    rotation = k % l

    copy(newArr[:rotation], nums[l-rotation:])
    
    copy(newArr[rotation:], nums[:l-rotation])
    
    copy(nums, newArr)
    
}
