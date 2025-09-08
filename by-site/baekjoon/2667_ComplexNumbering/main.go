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

	var N int
	fmt.Fscan(in, &N)

}

/* func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	g := make([][]int, N)
	for i :=0; i< N; i++ {
		var s string
		fmt.Fscan(in, &s)
		g[i] = make([]int, N)
		for j, ch := range s {    
			if ch == '1' {
				g[i][j] = 1
			}
		}
	}

	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		vis[i] = make([]bool, N)
	}

	type P struct{ r, c int }
	dr := [4]int{-1, 1, 0, 0}
	dc := [4]int{0, 0, -1, 1}



	var sizes []int

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if g[r][c] == 1 && !vis[r][c] {
				q := make([]P, 0, N*N)
				head := 0
				q = append(q, P{r, c})
				vis[r][c] = true

				cnt := 0
				for head < len(q) {
					cur := q[head]
					head++
					cnt++

					for k := 0; k < 4; k++ {
						nr := cur.r + dr[k]
						nc := cur.c + dc[k]
						if nr < 0 || nr >= N || nc < 0 || nc >= N {
							continue
						}
						if g[nr][nc] == 1 && !vis[nr][nc] {
							vis[nr][nc] = true
							q = append(q, P{nr, nc})
						}
					}
				}
				sizes = append(sizes, cnt)
			}
		}
	}
	sort.Ints(sizes)
	fmt.Fprintln(out, len(sizes))
	for _, v := range sizes {
		fmt.Fprintln(out, v)
	}
} */