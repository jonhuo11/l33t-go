package main

func PermutationsOfString(s string) []string {
	out := []string{}
	conv := []rune(s)
	
	swap := func(s1 []rune, a, b int) []rune {
		tmp := s1[a]
		s1[a] = s1[b]
		s1[b] = tmp
		return s1
	}
	var h func(str []rune, l int, out []string)
	h = func(str []rune, l int, out []string) {
		out = append(out, string(str))
		for i := l; i < len(str); i++ {
			swap(str, i, l)
			h(str, i + 1, out)
			swap(str, i, l)
		}
	}
	h(conv, 0, out)
	return out
}