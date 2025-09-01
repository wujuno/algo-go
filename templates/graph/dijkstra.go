
package graph

import "container/heap"

type Edge struct{ To, W int }
type Item struct{ Node, Dist int }
type PQ []Item
func (p PQ) Len() int { return len(p) }
func (p PQ) Less(i, j int) bool { return p[i].Dist < p[j].Dist }
func (p PQ) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *PQ) Push(x interface{}) { *p = append(*p, x.(Item)) }
func (p *PQ) Pop() interface{} { old := *p; v := old[len(old)-1]; *p = old[:len(old)-1]; return v }

// Dijkstra returns distance array from source s on adjacency list g (0-indexed).
func Dijkstra(g [][]Edge, s int) []int {
	const INF = int(1e18) // use big number; for tighter control switch to int64 if needed
	n := len(g)
	dist := make([]int, n)
	for i := range dist { dist[i] = INF }
	dist[s] = 0
	pq := &PQ{{s, 0}}; heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.Dist != dist[cur.Node] { continue }
		for _, e := range g[cur.Node] {
			nd := cur.Dist + e.W
			if nd < dist[e.To] {
				dist[e.To] = nd
				heap.Push(pq, Item{e.To, nd})
			}
		}
	}
	return dist
}
