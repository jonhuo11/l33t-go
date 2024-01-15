package container_with_most_water

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    mx := 0
    for l < r {
        area := (r - l) * min(height[l], height[r])
        if area > mx {
            mx = area
        }
        if height[l] < height[r] {
            c := height[l]
            for height[l] <= c && l < r {
                l++
            }
        } else {
            c := height[r]
            for height[r] <= c && l < r {
                r--
            }
        }
    }
    return mx
}