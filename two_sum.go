package main

func twoSum(nums []int, target int) [2]int {
	for i, a := range nums {
		for j, b := range nums {
			if a+b == target && i != j {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0,0}
}

func twoSumQuick(nums []int, target int) [2]int {
	m := map[int]int{}
	for i, _ := range nums {
		//fmt.Println(m)
		diff := target - nums[i]
		if _, ok := m[diff]; !ok {
			m[nums[i]] = i
		} else {
			return [2]int{m[diff], i}
		}
	}
	return [2]int{0,0}
}
