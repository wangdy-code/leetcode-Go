package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	ans := make([][]int, r)
	for k, _ := range ans {
		ans[k] = make([]int, c)
	}
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = mat[i/m][i%m]
	}
	return ans
}

func matrixReshape1(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	temp := make([]int, m*n)
	index := 0
	ans := make([][]int, r)
	for k := range ans {
		ans[k] = make([]int, c)
	}
	for i := 0; i < n; i++ {
		for j := 0; i < m; j++ {
			temp[index] = mat[i][j]
			index = index + 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; i < m; j++ {
			ans[i][j] = temp[i*c+j]
		}
	}
	return ans
}
