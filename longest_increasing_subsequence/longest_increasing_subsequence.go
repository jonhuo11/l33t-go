package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	// this is a 1D dp problem
    // longest common subsequence ending at index i
    // this way, we can extract both the length and the value which the next must be greater than
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }

    mx := 1
    for i := 1; i < len(nums); i++ {
        for j := i - 1; j >= 0; j-- {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
                mx = max(mx, dp[i])
            }
        }
    }
    return mx
}