package main

import "fmt"

func countSum(arr []int) map[int]int {
	result := make(map[int]int)
	for _, num := range arr {
		result[num]++
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3, 5, 1, 2, 3, 3, 3}
	count := countSum(arr1)

	fmt.Println(arr1)
	fmt.Println(count)
}
