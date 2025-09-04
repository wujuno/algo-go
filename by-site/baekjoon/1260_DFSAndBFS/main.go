package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, V int
	fmt.Fscan(in, &N, &M, &V)

	// 1) 그래프 만들기: 1-based
	g := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	// 2) 이웃 오름차순 정렬
	for i := 1; i <= N; i++ {
		sort.Ints(g[i])
	}

	// 3) DFS
	visited := make([]bool, N+1)
	dfsOrder := make([]int, 0, N)

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		dfsOrder = append(dfsOrder, u)
		for _, v := range g[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(V)

	// 4) DFS 출력
	for i, x := range dfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)

	// 5) BFS
	for i := range visited {
		visited[i] = false
	}
	bfsOrder := make([]int, 0, N)
	q := make([]int, 0, N)

	visited[V] = true
	q = append(q, V)

	for head := 0; head < len(q); head++ {
		u := q[head]
		bfsOrder = append(bfsOrder, u)
		for _, v := range g[u] {
			if !visited[v] {
				visited[v] = true
				q = append(q, v)
			}
		}
	}

	// 6) BFS 출력
	for i, x := range bfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)
}
