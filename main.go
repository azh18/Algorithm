package main

import (
	"github.com/Algorithm/FPTree"
	"fmt"
	"github.com/Algorithm/LIS"
)

func FPTreeTest(){
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

func LISTest(){
	input := []int{10,22,9,33,21,50,41,47,60,80}
	nLis, lis := LIS.GetLIS(input)
	fmt.Printf("The max length of increasing sub-sequence: %d\n", nLis)
	fmt.Printf("The LIS is %v", lis)
	return
}

func main() {
	LISTest()
}
