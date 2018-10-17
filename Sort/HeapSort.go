package Sort

import "math"

func IntPow(x int, y int) int {
	s := 1
	for i:=0;i<y;i++{
		s *= x
	}
	return s
}

type HeapItem struct{
	priority int
	elem interface{}
}

type Heap struct{
	node []*HeapItem
	size int
}

func NewHeap() *Heap{
	return &Heap{
		node:make([]*HeapItem, 0),
		size:0,
	}
}

func (p *Heap) swap(i,j int){
	temp := p.node[i]
	p.node[i] = p.node[j]
	p.node[j] = temp
}

func (p *Heap) up(i int){
	idx := i
	for idx > 0{
		father := (idx-1)/2
		if p.node[idx].priority > p.node[father].priority{
			p.swap(idx, father)
		}
		idx = father
	}
}

func (p *Heap) Pop() *HeapItem{
	retItem := p.node[0]
	p.node[0] = p.node[p.size-1]
	p.size--
	p.down(0)
	return retItem
}

func (p *Heap) Push(item *HeapItem){
	p.node = append(p.node, item)
	p.size++
	p.up(p.size-1)
}

func (p *Heap) down(i int){
	left := i*2+1
	right := i*2+2
	maxIdx := -1
	if left >= p.size && right >= p.size{
		return
	}
	if left >= p.size || right >= p.size{
		var toTest int
		if left >= p.size{
			toTest = right
		} else {
			toTest = left
		}
		if p.node[toTest].priority > p.node[i].priority{
			maxIdx = toTest
		} else {
			maxIdx = i
		}
	} else {
		if p.node[left].priority > p.node[right].priority{
			maxIdx = left
		} else {
			maxIdx = right
		}
		if !(p.node[maxIdx].priority > p.node[i].priority){
			maxIdx = i
		}
	}
	if maxIdx == i{
		return
	} else {
		temp := p.node[maxIdx]
		p.node[maxIdx] = p.node[i]
		p.node[i] = temp
	}
	p.down(maxIdx)
	return
}

// O(n) 复杂度，自下而上down操作
func (p *Heap) Heapify(){
	maxLevel := int(math.Log2(float64(p.size)))
	for level:=maxLevel-1;level>=0;level--{
		cnt := 0
		for i:=IntPow(2, level)-1; cnt<IntPow(2, level);i++{
			if i >= p.size {
				break
			}
			p.down(i)
			cnt++
		}
	}
	return
}

func HeapSort(input []int) (ret []int) {
	heap := NewHeap()
	for _, item := range input{
		heap.Push(&HeapItem{priority:item, elem:item})
	}
	for i:=0;i<len(input);i++{
		popItem := heap.Pop()
		ret = append(ret, popItem.elem.(int))
	}
	return ret
}