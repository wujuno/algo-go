package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	g := make([][]int, N+1)

	for i :=0; i < M; i++ {
		var a,b int
		fmt.Fscan(in, &a, &b)

		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	vis := make([]bool, N+1)
	q:= make([]int, N)
	head:=0
	q[0] = 1
	count := 0
	vis[1] = true

	for head < len(q) {
		cur:= q[head]
		head++
		for _, v := range g[cur] {
			if !vis[v] {
				vis[v] = true
				q = append(q, v)
				count++
			}

		}

	}
	fmt.Fprintln(out, count)
}

/* func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	g := make([][]int, N+1)

	for i :=0; i < M; i++ {
		var a,b int
		fmt.Fscan(in, &a, &b)

		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	vis := make([]bool, N+1)
	vis[1]= true
	q := make([]int, N)
	head := 0
	q = append(q, 1)
	count := 0

	for head < len(q) {
		cur := q[head]
		head++

		for _, v := range g[cur] {
			if !vis[v] {
				vis[v] = true
				q = append(q, v)
				count++
			}

		}
	}

	fmt.Fprintln(out, count)
}
 */