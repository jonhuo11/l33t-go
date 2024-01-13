package koko_eating_bananas

import (
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
    sort.Ints(piles)
    return helper(1, piles[len(piles) - 1], piles, h)
}

func helper(l, r int, piles []int, h int) int {
    hrs := 0
    mid := (l + r) / 2
    //fmt.Println(l, r, mid)
    for i, _ := range piles {
        hrs += int(math.Ceil(float64(piles[i]) / float64(mid)))
    }
    if hrs > h {
        if r - l == 1 {
            return r
        }
        if r - l == 0 {
            return mid + 1
        } else {
            return helper(mid, r, piles, h)
        }
    } else {
        if r - l == 1 {
            return min(math.Min(mid, helper(r, r, piles, h)))
        }
        if r == l {
            return mid
        }
        return min(mid, helper(l, mid, piles, h))
    } 
}