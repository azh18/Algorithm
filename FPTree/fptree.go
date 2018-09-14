package FPTree

import (
	"container/list"
	"sort"
)

type ItemRecord []string

type ItemTable struct {
	allItems         []ItemRecord
	allDistinctItems ItemRecord
	itemExist        map[string]bool
}

type FreqTableEntry struct {
	Item    string
	Freq    int
	AllElem *list.List
}

type FreqTable map[string]*FreqTableEntry

type FPNode struct {
	Father    *FPNode
	Daughters []*FPNode
	Item      string
	Freq      int
}

type FPTree struct {
	Root  *FPNode
	Table *FreqTable
}

func NewItemTable(allTrans [][]string) *ItemTable {
	itemTable := ItemTable{}
	for _, r := range allTrans {
		itemRecord := ItemRecord{}
		for _, i := range r {
			itemRecord = append(itemRecord, i)
		}
		itemTable.allItems = append(itemTable.allItems, itemRecord)
	}
	return &itemTable
}

func (it *ItemTable) GetFreqPattern(support int) (patterns [][]string) {
	it.Sort()
	it.GetDistinctItems()
	fpTree := NewFPTree()
	fpTree.BuildFPTree(it)
	fpTree.GetPatternFromFPTree(support, it.allDistinctItems, []string{}, &patterns)
	return
}

func (it *ItemTable) GetDistinctItems() {
	it.itemExist = map[string]bool{}
	for _, ir := range it.allItems {
		for _, item := range ir {
			it.AddItemsToDistinctItem(item)
		}
	}
	sort.Sort(&it.allDistinctItems)
}

func (it *ItemTable) AddItemsToDistinctItem(item string) {
	if !it.itemExist[item] {
		it.allDistinctItems = append(it.allDistinctItems, item)
		it.itemExist[item] = true
	}
}

func (it *ItemTable) Sort() {
	for _, ir := range it.allItems {
		sort.Sort(&ir)
	}
}

// clean item table, to make all items with freq >= support
// only clean itemTable.allItem, need to sort and build other structures manually
func (it *ItemTable) Clean(support int, freqTable *FreqTable) *ItemTable {
	newItemTable := ItemTable{}
	for _, r := range it.allItems {
		recordAfterClean := ItemRecord{}
		for _, i := range r {
			if (*freqTable)[i].Freq >= support {
				recordAfterClean = append(recordAfterClean, i)
			}
		}
		newItemTable.allItems = append(newItemTable.allItems, recordAfterClean)
	}
	return &newItemTable
}

func (ir *ItemRecord) Swap(i, j int) {
	temp := (*ir)[i]
	(*ir)[i] = (*ir)[j]
	(*ir)[j] = temp
}

func (ir *ItemRecord) Less(i, j int) bool {
	return (*ir)[i] < (*ir)[j]
}

func (ir *ItemRecord) Len() int {
	return len(*ir)
}

func NewFPTree() *FPTree {
	return &FPTree{
		Root: &FPNode{
			Item:      "",
			Father:    nil,
			Daughters: []*FPNode{},
			Freq:      0,
		},
		Table: &FreqTable{},
	}
}

func (n *FPNode) FindItemInDaughters(s string) *FPNode {
	for _, i := range n.Daughters {
		if i.Item == s {
			return i
		}
	}
	return nil
}

// hasExist: whether node has been established in fp-tree before adding 1 freq to it
func (t *FreqTable) AddNode(node *FPNode, hasExist bool) {
	item := node.Item
	if entry, ok := (*t)[item]; !ok {
		(*t)[item] = &FreqTableEntry{
			Item:    item,
			Freq:    1,
			AllElem: list.New(),
		}
		(*t)[item].AllElem.PushBack(node)
	} else {
		entry.Freq++
		if !hasExist {
			entry.AllElem.PushBack(node)
		}
	}
}

func (t *FPTree) BuildFPTree(table *ItemTable) {
	for _, ir := range table.allItems {
		pNode := t.Root
		for _, item := range ir {
			nextNode := pNode.FindItemInDaughters(item)
			if nextNode != nil {
				pNode = nextNode
				pNode.Freq++
				t.Table.AddNode(pNode, true)
			} else {
				nextNode = &FPNode{
					Father: pNode,
					Item:   item,
					Freq:   1,
				}
				pNode.Daughters = append(pNode.Daughters, nextNode)
				pNode = nextNode
				t.Table.AddNode(pNode, false)
			}
		}
	}
}

func (t *FPTree) GetPatternFromFPTree(support int, distinctItems ItemRecord, rootItem []string, patterns *[][]string) {
	if len(t.Root.Daughters) == 0 {
		return
	} else {
		for i := len(distinctItems) - 1; i >= 0; i-- {
			// build conditional tree for each item
			entry := (*t.Table)[distinctItems[i]]
			if entry.Freq < support {
				continue
			} else {
				// itself is a frequent pattern
				*patterns = append(*patterns, append(rootItem, entry.Item))
				// generate conditional tree for this entry
				newItemTable := &ItemTable{}
				newFreqTable := FreqTable{}
				for pNode := entry.AllElem.Front(); pNode != nil; pNode = pNode.Next() {
					leaf, _ := pNode.Value.(*FPNode)
					loopNum := leaf.Freq
					for i := 0; i < loopNum; i++ {
						newItemRecord := ItemRecord{}
						for p := leaf.Father; p != t.Root; p = p.Father {
							newItemRecord = append(newItemRecord, p.Item)
							newFreqTable.AddNode(&FPNode{Item: p.Item}, true)
						}
						newItemTable.allItems = append(newItemTable.allItems, newItemRecord)
					}
				}
				newItemTable = newItemTable.Clean(support, &newFreqTable)
				newItemTable.Sort()
				newItemTable.GetDistinctItems()
				newFPTree := NewFPTree()
				newFPTree.BuildFPTree(newItemTable)
				newFPTree.GetPatternFromFPTree(support, newItemTable.allDistinctItems, append(rootItem, entry.Item), patterns)
			}
		}
	}
	return
}
