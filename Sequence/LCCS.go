package Sequence

func GetLCCS(a []int, b []int) (ret []int) {
	maxLCCSLen := 0
	maxLCCSXY := []int{-1, -1}
	lenA := len(a)
	lenB := len(b)
	lccsLen := make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		lccsLen[i] = make([]int, lenB)
	}
	for i := 0; i < lenA; i++ {
		if a[i] == b[0] {
			lccsLen[i][0] = 1
			if lccsLen[i][0] > maxLCCSLen {
				maxLCCSLen = lccsLen[i][0]
				maxLCCSXY = []int{i, 0}
			}
		} else {
			lccsLen[i][0] = 0
		}
	}
	for i := 0; i < lenB; i++ {
		if a[0] == b[i] {
			lccsLen[0][i] = 1
			if lccsLen[0][i] > maxLCCSLen {
				maxLCCSLen = lccsLen[0][i]
				maxLCCSXY = []int{0, i}
			}
		} else {
			lccsLen[0][i] = 0
		}
	}
	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			if a[i] == b[j] {
				lccsLen[i][j] = lccsLen[i-1][j-1] + 1
				if lccsLen[i][j] > maxLCCSLen {
					maxLCCSLen = lccsLen[i][j]
					maxLCCSXY = []int{i, j}
				}
			} else {
				lccsLen[i][j] = 0
			}
		}
	}
	ret1 := make(IntSlice, 0)
	x, y := maxLCCSXY[0], maxLCCSXY[1]
	for x >= 0 && y >= 0 {
		ret1 = append(ret1, a[x])
		x--
		y--
	}
	ret = ret1.Reverse()
	return ret
}
