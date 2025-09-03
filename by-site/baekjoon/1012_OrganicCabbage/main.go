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

	var T int
	fmt.Fscan(in, &T)

	type pos struct {
			x,y int
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	for ;T >0; T-- {
		var M, N, K int
		fmt.Fscan(in, &M, &N, &K)

		grid := make([][]bool, N)
		vis := make([][]bool, N)

		for i :=0; i < N; i++ {
			grid[i] = make([]bool, M)
			vis[i] = make([]bool, M)
		}

		for i := 0; i <K; i++ {
			var x,y int
			fmt.Fscan(in, &x,&y)
			grid[y][x] = true
		}

		q := make([]pos,N*M)
		count := 0
		

		
		for y:= 0; y < N; y++ {
			for x:=0; x < M; x++ {
				if(!grid[y][x] || vis[y][x]) {
					continue
				}
				count++
				head, tail := 0,0
				q[tail] = pos{x,y}
				vis[y][x] = true
				tail++
				
				for head < tail {
					cur := q[head]
					head++

					for d:=0; d<4; d++ {
					nx := cur.x + dx[d]
					ny := cur.y + dy[d]

					if ny >= N || nx >= M || ny < 0 || nx < 0 {
						continue
					}
					if !grid[ny][nx] || vis[ny][nx] {
						continue
					}

					vis[ny][nx] = true
					q[tail] = pos{nx, ny}
					tail++
					}

				}

				
			}
		}
		fmt.Fprintln(out, count)
	}
}