package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	c1 := twoSum([]int{1,2,3,4,10}, 7)
	if c1 != [2]int{2,3} && c1 != [2]int{3,2} {
		t.Errorf("Test case 1 wrong")
	}
}

func TestTwoSumQuick(t *testing.T) {
	c1 := twoSumQuick([]int{1,2,3,4,5}, 9)
	if c1 != [2]int{3,4} && c1 != [2]int{4,3} {
		t.Errorf("Test case 1 wrong %v", c1)
	}
}
