package leetcode

/*遍历数组记录 行列下标  再遍历一次归零对应行列*/
func setZeroes(matrix [][]int) {
	rowIndex := make([]bool, len(matrix))
	columnIndex := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				rowIndex[i] = true
				columnIndex[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if rowIndex[i] || columnIndex[j] {
				r[j] = 0
			}
		}
	}

}

/*使用两个标识符 修改第一行和第一列的元素作为标识符*/
func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

/*修改第二个方法 使用1个标识符 修改第一列的元素作为标识符*/
func setZeroes2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}

}
