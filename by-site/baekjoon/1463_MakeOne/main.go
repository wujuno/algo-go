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

	var X int
	fmt.Fscan(in, &X)

	if X <= 1 {
		fmt.Fprintln(out, 0)
		return
	}

	dp := make([]int, X+1) // dp[0] 미사용, dp[1] = 0
	for i := 2; i <= X; i++ {
		// 1 빼기
		best := dp[i-1] + 1
		// 2로 나누기
		if i%2 == 0 && dp[i/2]+1 < best {
			best = dp[i/2] + 1
		}
		// 3으로 나누기
		if i%3 == 0 && dp[i/3]+1 < best {
			best = dp[i/3] + 1
		}
		dp[i] = best
	}

	fmt.Fprintln(out, dp[X])
}


//큰 문제의 최적해(최소 연산 수)가 더 작은 문제들의 최적해에서 만들어진다는 성질을 이용해, 작은 수부터 차근차근 “최솟값”을 저장해 가는 방식이에요.