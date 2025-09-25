package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Edge struct {
	to int
	w  int
}

type Item struct {
	node int
	dist int
}
type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	it := old[n-1]
	*h = old[:n-1]
	return it
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var V, E int
	fmt.Fscan(in, &V, &E)
	var K int
	fmt.Fscan(in, &K)

	g := make([][]Edge, V+1)
	for i := 0; i < E; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		g[u] = append(g[u], Edge{to: v, w: w}) // 방향 그래프
	}

	const INF = int(1<<60) // 충분히 큰 값
	dist := make([]int, V+1)
	for i := 1; i <= V; i++ {
		dist[i] = INF
	}
	dist[K] = 0

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Item{node: K, dist: 0})

	for h.Len() > 0 {
		cur := heap.Pop(h).(Item)
		if cur.dist != dist[cur.node] {
			continue // 이미 더 짧은 경로로 갱신된 경우 스킵
		}
		for _, e := range g[cur.node] {
			nd := cur.dist + e.w
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(h, Item{node: e.to, dist: nd})
			}
		}
	}

	for i := 1; i <= V; i++ {
		if dist[i] == INF {
			fmt.Fprintln(out, "INF")
		} else {
			fmt.Fprintln(out, dist[i])
		}
	}
}
