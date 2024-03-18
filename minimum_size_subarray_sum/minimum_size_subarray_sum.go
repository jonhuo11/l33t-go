package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
    mn := 0
    s := 0
    for _, v := range nums {
        mn += 1
        s += v
    }
    if s < target {
        return 0
    }
    l := 0
    sum := 0
    for r := 0; r < len(nums); r++ {
        //fmt.Println(mn)
        if sum + nums[r] >= target {
            mn = min(r - l + 1, mn)
            for l < r  {
                sum -= nums[l]
                l++
                //fmt.Println(l, r, r - l, sum + nums[r], mn)
                if sum + nums[r] < target {
                    break
                }
                mn = min(r - l + 1, mn)
            }
        }
        sum += nums[r]
    }
    return mn
}