package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] } 
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	h := &maxHeap{}
	heap.Init(h)

	for i:=0; i < N; i++ {
		var x int 
		fmt.Fscan(in, &x)

		if x == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(out, 0)
			} else {
				maxVal := heap.Pop(h).(int)
				fmt.Fprintln(out, maxVal)
			}
		} else {
			heap.Push(h,x)
		}
	}
}