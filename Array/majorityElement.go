# using hashTable to find mazority element
func majorityElement(nums []int) []int {
    countFrequency := make(map[int]int)
    for _, v := range nums{
        countFrequency[v]++
    }   
    neededCount := len(nums)/3
    countArr := make([]int, 0)
    for k, v := range countFrequency {
        if v > neededCount{
            countArr = append(countArr, k)
        }
    }
    return countArr
}
