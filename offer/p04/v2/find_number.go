package v2

// 网友的，思路是一样的，不过从从左下角开始
// matrix[i][j] > target，值大了，往上边找，i--
// matrix[i][j] < target，值小了，往右边找，j++

// 跟v1效果差不多。空间一样的，时间上，相对更稳定

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		x := matrix[i][j]
		switch {
		case x > target:
			i--
		case x < target:
			j++
		case x == target:
			return true
		}
	}
	return false
}
