package Sequence

type IntSlice []int

func (a IntSlice) Reverse() (ret []int) {
	for i := len(a) - 1; i >= 0; i-- {
		ret = append(ret, a[i])
	}
	return
}
