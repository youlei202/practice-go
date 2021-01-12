package main

import (
	"fmt"
	"math/rand"
)

func MergeSort(nums []int, len int) {
	temp := make([]int, len)
	mergeSort(nums, 0, len-1, temp)
}

func mergeSort(nums []int, l int, r int, temp []int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	mergeSort(nums, l, m, temp)
	mergeSort(nums, m+1, r, temp)

	i, j, k := l, m+1, l
	for {
		if i <= m || j <= r {
			if i > m || (j <= r && nums[j] <= nums[i]) {
				temp[k] = nums[j]
				j += 1
			} else {
				temp[k] = nums[i]
				i += 1
			}
			k += 1
		} else {
			break
		}
	}
	for k = l; k <= r; k++ {
		nums[k] = temp[k]
	}
}

func main() {
	nums := rand.Perm(20000)
	fmt.Println(nums)
	MergeSort(nums, len(nums))
	fmt.Println(nums)
}
