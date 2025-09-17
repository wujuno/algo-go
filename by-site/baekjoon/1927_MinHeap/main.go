package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type IntMinHead []int

func (h IntMinHead) Len() int {return len(h)}
func (h IntMinHead) Less(i, j int) bool {return h[i] < h[j]}
func (h IntMinHead) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntMinHead) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *IntMinHead) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	h := &IntMinHead{}
	heap.Init(h)

	for i:=0 ; i <N; i++ {
		var num int 
		fmt.Fscan(in, &num)

		if num != 0 {
			heap.Push(h ,num)
		} else {
			if h.Len() == 0 {
				fmt.Fprintln(out, 0)
			} else {
				minValue := heap.Pop(h)
				fmt.Fprintln(out, minValue)
			}
		}
	}
}
