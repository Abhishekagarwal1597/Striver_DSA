// remove duplicates element in an array
// using 2 pointer approach
// only update the element if they are differ and update the one pointer as well on same
func removeDuplicates(nums []int) int { 
    i := 0
    for j:=1; j<len(nums); j++{
        if nums[i] != nums[j] {
            nums[i+1] = nums[j]
            i++
        } 
    }
    return i+1
}

// Using hashMap, create a hashMap store all unique element in it and append it in a slice(as map key are not ordered when list)
// then copy them to the orignal array

    // hashMap := make(map[int]int)
    // var newArr []int
    // for _, v := range nums{
    //     if _, ok := hashMap[v]; !ok{
    //         newArr = append(newArr, v)
    //         hashMap[v] = 1
    //     }
    // }

    // index := 0
    // for _, v := range newArr{
    //     nums[index] = v
    //     index++
    // }
    // return len(hashMap)
// }
