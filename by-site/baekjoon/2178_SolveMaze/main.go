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

	var n, m int
	fmt.Fscan(in, &n, &m)

	grid := make([][]byte, n)
	for i := 0; i <n; i++ {
		var s string
		fmt.Fscan(in, &s)
		row := make([]byte, m)
		for j :=0; j <m; j++ {
			if s[j] == '1'{
				row[j] = 1
			} else {
				row[j] = 0
			}
		}
		grid[i] = row
	}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
	}

	type pos struct {
		x,y int
	}
	q := make([]pos, n*m)
	head, tail := 0, 0

	if grid[0][0] == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	dist[0][0] = 1
	q[tail] = pos{0,0}
	tail++

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}

	for head < tail {
		cur := q[head]
		head++

		if cur.x == n-1 && cur.y == m-1 {
			fmt.Fprintln(out, dist[cur.x][cur.y])
			return
		}

		for dir := 0; dir < 4; dir++ {
			nx := cur.x + dx[dir]
			ny := cur.y + dy[dir]
			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}
			if grid[nx][ny] == 1 && dist[nx][ny] == 0 {
				dist[nx][ny] = dist[cur.x][cur.y] + 1
				q[tail] = pos{nx, ny}
				tail++
			}
		}
	}
}

