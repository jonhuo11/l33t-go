package jugsofwater

import "testing"

func TestJugsOfWater(t *testing.T) {
	if !jugsOfWater() {
		t.Fatalf("not exist")
	}
}