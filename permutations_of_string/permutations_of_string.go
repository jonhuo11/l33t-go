package permutationsofstring

func permutationsOfString(s string) []string {
	o := []string{}
	conv := []rune(s)
	
	swap := func(s1 []rune, a, b int) {
		tmp := s1[a]
		s1[a] = s1[b]
		s1[b] = tmp
	}
	var h func(str []rune, l int, out *[]string)
	h = func(str []rune, l int, out *[]string) {
		if l == len(str) - 1 {
			*out = append(*out, string(str))
		}
		for i := l; i < len(str); i++ {
			swap(str, i, l)
			h(str, l + 1, out)
			swap(str, i, l)
		}
	}
	h(conv, 0, &o)
	return o
}