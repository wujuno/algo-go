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

	var C, N int
	fmt.Fscan(in, &C,&N)

	g := make([][]int, C+1)

	for i := 0; i < N; i++ {
		var a,b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 1; i <= C; i++ {
		sort.Ints(g[i])
	}

	visited := make([]bool, C+1)
	q := make([]int, 0, C)

	visited[1] = true
	q = append(q, 1)

	head := 0
	infected := 0

	for head < len(q) {
		cur := q[head]
		head++

		for _, nxt := range g[cur] {
			if !visited[nxt] {
				visited[nxt] = true
				q = append(q, nxt)
				infected++
			}
		}
	}
	fmt.Fprintln(out, infected)
}