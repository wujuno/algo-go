package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct{ x, y, d int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var M, N int
	fmt.Fscan(in, &M, &N) // M=cols, N=rows

	grid := make([][]int, N)
	for i := range grid {
		grid[i] = make([]int, M)
	}

	q := make([]Node, 0, N*M)
	head := 0
	rawCount := 0 // 안 익은(0) 개수

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fscan(in, &grid[i][j])
			if grid[i][j] == 1 {
				q = append(q, Node{j, i, 0}) // 모든 1을 시작점으로
			} else if grid[i][j] == 0 {
				rawCount++
			}
		}
	}

	// 처음부터 모두 익어있다면 0일
	if rawCount == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	maxDay := 0

	for head < len(q) {
		cur := q[head]
		head++

		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dx[k], cur.y+dy[k]
			if nx < 0 || ny < 0 || nx >= M || ny >= N {
				continue
			}
			if grid[ny][nx] != 0 {
				continue
			}
			// 익게 만들기
			grid[ny][nx] = 1
			rawCount--
			nd := cur.d + 1
			if nd > maxDay {
				maxDay = nd
			}
			q = append(q, Node{nx, ny, nd})
		}
	}

	if rawCount > 0 {
		fmt.Fprintln(out, -1) // 막혀서 못 익은 토마토 존재
	} else {
		fmt.Fprintln(out, maxDay)
	}
}
