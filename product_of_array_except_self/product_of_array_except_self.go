package product_of_array_except_self

func productExceptSelf(nums []int) []int {
    // [ 1  2  6 24]
    // [24 24 12  4]

    forward := []int{nums[0]}
    backward := make([]int, len(nums))
    backward[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1; i < len(nums); i++ {
        forward = append(forward, nums[i] * forward[i - 1])
    }
    for i := len(nums) - 2; i >= 0; i-- {
        backward[i] = backward[i + 1] * nums[i]
    }
    out := make([]int, len(nums))
    out[0] = backward[1]
    out[len(nums) - 1] = forward[len(forward) - 2]
    for i := 1; i < len(nums) - 1; i++ {
        out[i] = forward[i - 1] * backward[i + 1]
    }
    return out
}