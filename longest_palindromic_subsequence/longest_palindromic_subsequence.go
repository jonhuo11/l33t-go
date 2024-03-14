package longestpalindromicsubsequence

func longestPalindromeSubseq(s string) int {
    size := len(s)
    dp := [1001][1001]int{}
    for isize := 0; isize < size; isize++ {
        l := 0
        r := isize
        for r < size {
            if l == r { // len 1 palin
                dp[l][r] = 1
            } else {
                if s[l] == s[r] { // can expand palin
                    dp[l][r] = 2 + dp[l + 1][r - 1]
                } else {
                    dp[l][r] = max(dp[l + 1][r], dp[l][r-1])
                }
            }

            l++
            r++
        }
    }
    //fmt.Println(dp)
    return dp[0][size-1]
}
