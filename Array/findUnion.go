// find the union of 2 array

package main

import "fmt"

func main() {
	//fmt.Printf("hello world")
	arr1 := []int{1, 2, 3, 3, 4}
	arr2 := []int{2, 3, 4, 5, 5}
	resArr := findUnion(arr1, arr2)
	fmt.Print(resArr)
}


func findUnion(arr1 []int, arr2 []int)[]int{
	// we can also use 2 pointer approach here
  // USING SET
	newMap := make(map[int]int)
	//outSlice := make([]int, len(arr1))
	var outSlice []int
	for _, val := range arr1 {
		if _, ok := newMap[val]; !ok {
			newMap[val] = 1
		}
	}
	for _, val := range arr2 {
		if _, ok := newMap[val]; !ok {
			newMap[val] = 1
		}
	}
	fmt.Println(newMap)
	for k := range newMap{
		outSlice = append(outSlice, k)
	}
	return outSlice


	//return unionArr
}
