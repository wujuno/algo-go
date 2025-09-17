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

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	dp := make([]int, n) // 한 줄 DP

	for i := 0; i < n; i++ {
		// 현재 행 입력
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			fmt.Fscan(in, &row[j])
		}

		// 오른쪽 -> 왼쪽으로 갱신 (덮어쓰기 방지)
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[0] = dp[0] + row[0] // 왼쪽 가장자리: 위에서만 옴
			} else if j == i {
				dp[j] = dp[j-1] + row[j] // 오른쪽 가장자리: 왼위에서만 옴
			} else {
				// 가운데: 위(dp[j])와 왼위(dp[j-1]) 중 최댓값 선택
				if dp[j-1] > dp[j] {
					dp[j] = dp[j-1] + row[j]
				} else {
					dp[j] = dp[j] + row[j]
				}
			}
		}
	}

	// 마지막 줄의 최대값이 정답
	ans := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	fmt.Fprintln(out, ans)
}
