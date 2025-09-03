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
			grid[i] = make([]bool, 0, M)
			vis[i] = make([]bool, 0, M)
		}

		for i := 0; i <K; i++ {
			var x,y int
			fmt.Fscan(in, &x,&y)
			grid[y][x] = true
		}

		q := make([]pos,0,N*M)
		count := 0

		for y:= 0; y < N; y++ {
			for x:=0; x < M; x++ {
				if(!grid[y][x] || vis[y][x]) {
					continue
				}

				count++

				for d:=0; d<4; d++ {
					
				}
			}
		}

	}
}