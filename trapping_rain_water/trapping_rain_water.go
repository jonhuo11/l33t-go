package trappingrainwater

import "fmt"

func trap(height []int) int {
    l, r := 0, 1
    total := 0
    last_tallest := 0
    for {
        fmt.Println(l, r, last_tallest, total)
        if l >= len(height) {
            break
        }
        if r >= len(height) {
            l++
            r = l + 1
            last_tallest = l
            continue
        }
        if height[r] >= height[l] || height[r] >= height[last_tallest] {
            if r == l + 1 {
                fmt.Println("skipping")
                r += 1
                l++
                last_tallest = l
                continue
            }
            if height[r] >= height[last_tallest] && !(height[r] >= height[l]) {
                fmt.Println("checking for taller")
                // lookahead to see if there is any taller
                flag := l
                for i := r + 1; i < len(height); i++ {
                    if height[i] > height[r] {
                        r = i
                        flag = r
                        break
                    }
                }
                if flag != l { // found a taller one
                    fmt.Println("found a taller at ", flag)
                    continue
                }
            }
            h := min(height[l], height[r])
            fmt.Println("scan start", h)
            var i int
            if height[l] < height[r] {
                i = l
            } else {
                i = l + 1
            }
            for i < r {
                fmt.Println(i, h - height[i])
                total += (h - height[i])
                i++
            }
            l = r
            r += 1
            last_tallest = l
            fmt.Println("end", l, r, last_tallest, total)
            continue
        } else {
            if last_tallest == l {
                last_tallest = r
            } else {
                if height[r] >= height[last_tallest] {
                    last_tallest = r
                }
            }
        }
        r++
    }
    return total
}