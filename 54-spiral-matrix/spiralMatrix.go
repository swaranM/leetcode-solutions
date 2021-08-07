package spiralmatrix

func spiralOrder(matrix [][]int) []int {

	rowBegin := 0
	rowEnd := len(matrix) - 1
	colBegin := 0
	colEnd := len(matrix[0]) - 1

	res := make([]int, 0)

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				res = append(res, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				res = append(res, matrix[i][colBegin])
			}

		}
		colBegin++

	}
	return res
}
