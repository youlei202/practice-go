package main

import (
	"fmt"
	"math/rand"
)

func Swap(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func InsertionSort(nums []int, len int) {
	for i := 0; i < len; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			Swap(nums, j, j-1)
		}
	}
}

func main() {
	nums := rand.Perm(10)
	fmt.Println(nums)
	InsertionSort(nums, len(nums))
	fmt.Println(nums)
}
