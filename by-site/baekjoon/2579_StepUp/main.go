package main

import (
	"bufio"
	"fmt"
	"os"
)

/* func max(a, b int) int {
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
} */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	score := make([]int, 301) // ok. N+1 로 지정하면 1이나 2일때 아래 dp[2]나 dp[3] 대입시 에러남.

	for i:=1; i <= N; i++ {
		fmt.Fscan(in, &score[i])
	} 

	
	dp := make([]int, 301)
	dp[1] = score[1] // ok. 마지막 계단은 무조건 밟아야 하므로. 마지막계단일 때를 가정.
	dp[2] = score[1] + score[2]
	dp[3] = max(score[1]+score[3], score[2]+score[3])

	for i := 4; i <= N; i++ {
		dp[i] = max(dp[i-2]+score[i], dp[i-3]+score[i-1]+score[i]) // ok. 연속 세개 밟으면 안됨 규칙을 어기지 않는 상황들만
	}

	fmt.Fprintln(out, dp[N])
}