package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(nums []int, len int) {
	quickSort(nums, 0, len-1)
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	pivot := nums[i]
	for {
		if i >= j {
			break
		}
		for {
			if i < j && nums[j] >= pivot {
				j -= 1
			} else {
				break
			}
		}
		nums[i] = nums[j]
		for {
			if i < j && nums[i] <= pivot {
				i += 1
			} else {
				break
			}
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

func main() {
	nums := rand.Perm(10)
	fmt.Println(nums)
	QuickSort(nums, len(nums))
	fmt.Println(nums)
}
