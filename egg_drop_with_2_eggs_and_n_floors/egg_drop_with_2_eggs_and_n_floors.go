package eggdropwith2eggsandnfloors

func twoEggDrop(n int) int {
    dp := [1001][2]int{}
    
    for f := 1; f <= n; f++ {
        dp[f][0] = f
        dp[f][1] = f // defaults for 1 egg
    }

    for f := 1; f <= n; f++ {
        for s := 1; s < f; s++ {
            drops := 1 + max(dp[s - 1][0], dp[f - s][1])
            dp[f][1] = min(dp[f][1], drops)
        }
    }
    return dp[n][1]
}