package searchina2dmatrix

func searchInA2DMatrix(matrix [][]int, target int) bool {
	l, r := 0, (len(matrix[0]) * len(matrix)) - 1
	for l < r {
		if r - l == 1 {
			return matrixValueFromArrIndex(matrix, l) == target || matrixValueFromArrIndex(matrix, r) == target
		}
		mid := (l + r) / 2
		if matrixValueFromArrIndex(matrix, mid) < target {
			l = mid
		} else if matrixValueFromArrIndex(matrix, mid) > target {
			r = mid
		} else {
			return true
		}
	}
	return matrixValueFromArrIndex(matrix, l) == target
}

func matrixIndexFromArrIndex(matrix [][]int, i int) [2]int {
	w := len(matrix[0])
	return [2]int{i / w, i % w}
}

func matrixValueFromArrIndex(matrix [][]int, i int) int {
	w := len(matrix[0])
	return matrix[i / w][i % w]
}