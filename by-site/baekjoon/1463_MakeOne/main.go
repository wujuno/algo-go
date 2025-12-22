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



	dp := make([]int, X+1)

	for i := 2; i <= X; i++ {
		// i에 도달하기 직전 단계는 아래의 세 경우 중 하나여야 한다.
		best := dp[i-1] + 1 // 기본
		if i%2 == 0 && dp[i/2]+1 < best { // 2로 나눠지는데 이전 2로 곱했던 dp 값에 1 추가한게 위의 best보다 작을 때
			best = dp[i/2] + 1
		}
		if i%3 == 0 && dp[i/3]+1 < best {
			best = dp[i/3] + 1
		}
		dp[i] = best
	}

	fmt.Fprintln(out, dp[X])
}

/* func main() {
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
} */

