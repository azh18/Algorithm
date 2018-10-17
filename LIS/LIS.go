package LIS

type Elem struct{
	Idx int
	Val int
}

// find the smallest idx in end slice where end[idx] > val
func binaryFindEnd(end []Elem, val int, nEnd int) int{
	left := 0
	right := nEnd-1
	for left < right {
		middle := (left+right)/2
		if end[middle].Val <= val{
			left = middle + 1
		} else {
			right = middle
		}
	}
	if end[left].Val <= val{
		return nEnd // no elem in end that end[idx] > val
	} else {
		return left // return the idx normally
	}
}

func GetLIS(input []int) (nLIS int, LIS []int){
	nIn := len(input)
	if nIn == 0{
		return 0, []int{}
	}
	ancestor := make([]Elem, nIn)
	end := make([]Elem, 1)
	end[0] = Elem{0,input[0]}
	ancestor[0] = Elem{-1, -1}
	nEnd := 1
	for i:=1;i<nIn;i++{
		last_idx := binaryFindEnd(end, input[i], nEnd)
		if last_idx == nEnd{
			end = append(end, Elem{i, input[i]})
			ancestor[i] = Elem{end[nEnd-1].Idx, end[nEnd-1].Val}
			nEnd++
		} else {
			end[last_idx] = Elem{i, input[i]}
			if last_idx == 0{
				ancestor[i] = Elem{-1, -1}
			} else {
				ancestor[i] = Elem{end[last_idx-1].Idx, end[last_idx-1].Val}
			}
		}
	}
	nLIS = nEnd
	// get LIS by rescan
	lisRevert := make([]int, 0)
	lisRevert = append(lisRevert, end[nEnd-1].Val)
	ancestorIdx := ancestor[end[nEnd-1].Idx].Idx
	//if ancestorIdx == -1 {
	//	return nEnd, lisRevert
	//}
	for ancestorIdx != -1{
		val := input[ancestorIdx]
		lisRevert = append(lisRevert, val)
		ancestorIdx = ancestor[ancestorIdx].Idx
	}
	for i:=len(lisRevert)-1;i>=0;i--{
		LIS = append(LIS, lisRevert[i])
	}
	return
}
