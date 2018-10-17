package main

import (
	"fmt"
	"github.com/Algorithm/FPTree"
	"github.com/Algorithm/Sequence"
	"github.com/Algorithm/Sort"
)

func FPTreeTest() {
	trans := [][]string{{"a", "b", "c", "d", "e", "f", "g", "h"},
		{"b", "d", "e", "f", "j"},
		{"a", "f", "g"},
		{"a", "b", "d", "i", "k"},
		{"a", "b", "e", "g"},
	}
	trans2 := [][]string{{"A", "B"},
		{"B", "C", "D"},
		{"A", "C", "D", "E"},
		{"A", "D", "E"},
		{"A", "B", "C"},
	}
	itemTable := FPTree.NewItemTable(trans)
	patterns := itemTable.GetFreqPattern(3)
	fmt.Printf("pattern for trans= %v\n", patterns)
	itemTable2 := FPTree.NewItemTable(trans2)
	patterns2 := itemTable2.GetFreqPattern(2)
	fmt.Printf("pattern for trans2 = %v\n", patterns2)

	return
}

func LISTest() {
	input := []int{10, 22, 9, 33, 21, 50, 41, 47, 60, 80}
	nLis, lis := Sequence.GetLIS(input)
	fmt.Printf("The max length of increasing sub-sequence: %d\n", nLis)
	fmt.Printf("The Sequence is %v", lis)
	return
}

func LCCSTest() {
	a := []int{1, 3, 5, 6, 7}
	b := []int{3, 5, 7}
	lccs := Sequence.GetLCCS(a, b)
	fmt.Printf("The longest continuous common subsequence is %v", lccs)
	return
}

func LCSTest() {
	a := []int{1, 3, 5, 6, 7}
	b := []int{3, 4, 5, 7, 9}
	lcs := Sequence.GetLCS(a, b)
	fmt.Printf("The longest common subsequence is %v", lcs)
	return
}

func HeapSortTest(){
	input := []int{10, 22, 9, 33, 21, 50, 41, 47, 60, 80}
	sorted := Sort.HeapSort(input)
	fmt.Printf("the sorted slice is %v\n", sorted)
	return
}

func main() {
	// LISTest()
	//LCCSTest()
	//LCSTest()
	HeapSortTest()
}
