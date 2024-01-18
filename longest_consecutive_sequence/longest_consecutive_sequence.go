package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    m := map[int]bool{}
    for _, v := range nums {
        m[v] = true
    }
    mx := 1

    for _, v := range nums {
        if _, in := m[v - 1]; in {
            continue
        } else {
            sum := 1
            t := v + 1
            for m[t] {
                sum++
                t++
            }
            if sum > mx {
                mx = sum
            }
            if sum > len(nums) / 2 {
                break
            }
        }
    }
    return mx
}