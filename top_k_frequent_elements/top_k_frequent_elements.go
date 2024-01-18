package top_k_frequent_elements

import "slices"

func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
    max := 0
    for _, v := range nums {
        if _, ok := m[v]; ok {
            m[v]++
        } else {
            m[v] = 1
        }
        if m[v] > max {
            max = m[v]
        }
    }
    buckets := make([][]int, max + 1)
    for k, v := range m {
        buckets[v] = append(buckets[v], k)
    }
    out := make([]int, 0)
    slices.Reverse(buckets)
    for _, v := range buckets {
        if (k == 0) {
            break
        }
        if len(v) == 0 {
            continue
        }
        out = append(out, v...)
        k -= len(v)
    }
    return out
}
