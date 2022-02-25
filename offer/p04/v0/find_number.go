package v0

// 对角平移+二分查找法
/*
该方案不行，总有特例。
大概从思路上就是错的。
*/
type xyRange struct {
	x, y   int
	length int
}

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	// 先大范围排除
	n := len(matrix)
	if n == 0 {
		return false
	}

	if n == 1 {
		return findNumberIn1DArray(matrix[0], target)
	}

	m := len(matrix[0])
	if m == 0 {
		return false
	}
	if m == 1 {
		return findNumberIn1DArray2(matrix, target)
	}

	min := matrix[0][0]
	max := matrix[n-1][m-1]
	if target < min || target > max {
		return false
	}
	if target == min || target == max {
		return true
	}

	length := minValue(n, m)
	ans := binarySearch(matrix, target, xyRange{0, 0, length})
	if ans {
		return true
	}
	if n > 1 {
		ans = binarySearch(matrix, target, xyRange{1, 0, length})
		if ans {
			return true
		}
	}
	if m > 1 {
		ans = binarySearch(matrix, target, xyRange{0, 1, length})
		if ans {
			return true
		}
	}
	len1 := length - 1
	for j := 1; j < m; j++ {
		ans = binarySearch(matrix, target, xyRange{0, j, len1})
		if ans {
			return true
		}
		len1--
	}
	len2 := length - 1
	for i := 1; i < n; i++ {
		ans = binarySearch(matrix, target, xyRange{i, 0, len2})
		if ans {
			return true
		}
		len2--
	}

	return false
}

func binarySearch(matrix [][]int, target int, searchRange xyRange) bool {
	if searchRange.length == 1 {
		return target == matrix[searchRange.x][searchRange.y]
	} else {
		xlo, ylo := searchRange.x, searchRange.y
		xhi, yhi := xlo+searchRange.length-1, ylo+searchRange.length-1
		xMax, yMax := len(matrix), len(matrix[0])
		if xhi < xMax && yhi < yMax {
			for xlo <= xhi && ylo <= yhi {
				xmid := xlo + (xhi-xlo)/2
				ymid := ylo + (yhi-ylo)/2
				midVal := matrix[xmid][ymid]
				if midVal > target {
					xhi, yhi = xmid-1, ymid-1
				} else if midVal < target {
					xlo, ylo = xmid+1, ymid+1
				} else {
					return true
				}
			}
		}

		return false
	}
}

func findNumberIn1DArray(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midVal := nums[mid]
		if midVal > target {
			hi = mid - 1
		} else if midVal < target {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func findNumberIn1DArray2(matrix [][]int, target int) bool {
	lo, hi := 0, len(matrix)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midVal := matrix[mid][0]
		if midVal > target {
			hi = mid - 1
		} else if midVal < target {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func minValue(i, j int) int {
	if j > i {
		return j
	}
	return i
}
