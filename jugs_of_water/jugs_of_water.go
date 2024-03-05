package jugsofwater

func jugsOfWater() bool {
	m := map[[3]int]bool{}

	maxes := [3]int{4, 7, 10}
	s:= [3]int{4, 7, 0}

	queue := [][3]int{s}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		for jug := range state {
			for otherjug := range state {
				if jug == otherjug || state[otherjug] == maxes[otherjug] || state[jug] == 0 {
					continue
				}
				
				otherneed := maxes[otherjug] - state[otherjug]

				cp := [3]int{state[0], state[1], state[2]}
				
				if otherneed > cp[jug] {
					cp[otherjug] += cp[jug]
					cp[jug] = 0
				} else {
					cp[otherjug] = maxes[otherjug]
					cp[jug] -= otherneed
				}
				
				if _, in := m[cp]; in {
					continue
				}
				m[cp] = true
				queue = append(queue, cp)
			}
		}
	}

	for k := range m {
		//fmt.Println(k)
		if k[0] == 2 {
			return true
		}
	}
	return false
}
