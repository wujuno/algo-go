package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	score := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &score[i])
	}

	if N == 1 {
		fmt.Fprintln(out, score[1])
		return
	}
	if N == 2 {
		fmt.Fprintln(out, score[1]+score[2])
		return
	}

	dp := make([]int, N+1)
	dp[1] = score[1]
	dp[2] = score[1] + score[2]
	dp[3] = max(score[1]+score[3], score[2]+score[3])

	for i := 4; i <= N; i++ {
		dp[i] = max(dp[i-2]+score[i], dp[i-3]+score[i-1]+score[i])
	}

	fmt.Fprintln(out, dp[N])
}
