package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, target int) int {
	var s func(arr *[]int, l, r, t int) int
	sort.Ints(nums)
	s = func (arr *[]int, l, r, t int) int {
		fmt.Println(*arr, l, r, t)
		if r - l == 1  {
			if (*arr)[l] == t {
				return l
			}
			if (*arr)[r] == t {
				return r
			}
			return -1
		}
		if r == l && (*arr)[l] != t {
			return -1
		}
		mid := (l + r) / 2
		if t < (*arr)[mid] {
			return s(arr, l, mid, t)
		} else if t > (*arr)[mid] {
			return s(arr, mid, r, t)
		} else {
			return mid
		}
	}
	return s(&nums, 0, len(nums) - 1, target)
}