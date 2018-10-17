package Sequence

type Coordinate struct {
	x, y     int
	increase bool
}

func GetLCS(a []int, b []int) (ret []int) {
	lenA := len(a)
	lenB := len(b)
	lcsLen := make([][]int, lenA+1)
	ancestor := make([][]Coordinate, lenA+1)
	for i := 0; i <= lenA; i++ {
		lcsLen[i] = make([]int, lenB+1)
		ancestor[i] = make([]Coordinate, lenB+1)
	}
	lcsLen[0][0] = 0
	ancestor[0][0] = Coordinate{-1, -1, false}
	for i := 0; i < lenA; i++ {
		lcsLen[i+1][0] = 0
		ancestor[i+1][0] = Coordinate{-1, -1, false}
	}
	for i := 0; i < lenB; i++ {
		lcsLen[0][i+1] = 0
		ancestor[0][i+1] = Coordinate{-1, -1, false}
	}
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			var increase int
			if a[i] == b[j] {
				increase = 1
			} else {
				increase = 0
			}
			if lcsLen[i][j]+increase > lcsLen[i+1][j] {
				if lcsLen[i][j]+increase > lcsLen[i][j+1] {
					lcsLen[i+1][j+1] = lcsLen[i][j] + increase
					ancestor[i+1][j+1] = Coordinate{i, j, increase == 1}
				} else {
					lcsLen[i+1][j+1] = lcsLen[i][j+1]
					ancestor[i+1][j+1] = Coordinate{i, j + 1, false}
				}
			} else {
				if lcsLen[i+1][j] > lcsLen[i][j+1] {
					lcsLen[i+1][j+1] = lcsLen[i+1][j]
					ancestor[i+1][j+1] = Coordinate{i + 1, j, false}
				} else {
					lcsLen[i+1][j+1] = lcsLen[i][j+1]
					ancestor[i+1][j+1] = Coordinate{i, j + 1, false}
				}
			}
		}
	}
	retRevert := make(IntSlice, 0)
	x, y := lenA, lenB
	for x >= 0 && y >= 0 {
		if ancestor[x][y].increase {
			retRevert = append(retRevert, a[x-1])
		}
		newCorr := ancestor[x][y]
		x = newCorr.x
		y = newCorr.y
	}
	ret = retRevert.Reverse()
	return ret
}
