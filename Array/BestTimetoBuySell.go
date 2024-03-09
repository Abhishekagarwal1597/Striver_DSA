# Intution approach
# We will linearly travel the array. We can maintain a minimum from the start of the array and compare it with every element of the array, 
# if it is greater than the minimum then take the difference and maintain it in max, otherwise update the minimum.
func maxProfit(prices []int) int {
    min, maxDiff := prices[0], 0
    for i:=1;i<len(prices);i++{
        diff := prices[i] - min
        if diff > maxDiff {
            maxDiff = diff
        }
        if prices[i] < min{
            min = prices[i]
        }
    }
    return maxDiff
}
