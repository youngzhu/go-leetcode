package v1

// FindNumberIn2DArray 官方的，从右上角开始
// matrix[i][j] > target，值大了，往左边找，j--
// matrix[i][j] < target，值小了，往下边找，i++
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	i, j := 0, m-1
	for i < n && j >= 0 {
		x := matrix[i][j]
		switch {
		case x > target:
			j--
		case x < target:
			i++
		case x == target:
			return true
		}
	}
	return false
}
