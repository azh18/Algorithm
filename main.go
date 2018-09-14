package main

import (
	"github.com/Algorithm/FPTree"
	"fmt"
)

func main() {
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
