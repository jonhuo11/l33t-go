package permutationsofstring

import (
	"testing"
)

func TestPermutationsOfString(t *testing.T) {
	c1 := permutationsOfString("abc")
	//fmt.Println(c1)
	c1_expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	if len(c1) != len(c1_expected) {
		t.Errorf("Test case 1 failed")
	}
	for _, v := range c1 {
		found := false
		for _, ve := range c1_expected {
			if v == ve {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Test case 1 failed")
		}
	}
}