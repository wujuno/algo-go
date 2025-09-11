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

	g := make([][]int, N)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(in, &s)
		g[i] = make([]int, M)
		for j, ch := range s {
			if ch == '1' {
				g[i][j] = 1
			}
		}
	}

	type P struct{ x, y int }
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}

	dist := make([][]int, N)
	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, M)
		vis[i] = make([]bool, M)
	}

	// BFS 시작: (0,0)
	q := make([]P, 0, N*M)
	head := 0
	q = append(q, P{0, 0})
	vis[0][0] = true
	dist[0][0] = 1

	for head < len(q) {
		cur := q[head]
		head++

		for d := 0; d < 4; d++ {
			nx := cur.x + dx[d]
			ny := cur.y + dy[d]

			if nx < 0 || nx >= N || ny < 0 || ny >= M {
				continue
			}
			if g[nx][ny] == 0 || vis[nx][ny] {
				continue
			}

			vis[nx][ny] = true
			dist[nx][ny] = dist[cur.x][cur.y] + 1
			q = append(q, P{nx, ny})
		}
	}

	fmt.Fprintln(out, dist[N-1][M-1])
}

