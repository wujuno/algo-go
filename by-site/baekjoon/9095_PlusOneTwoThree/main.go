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

	const MAXN = 11
	dp := make([]int, MAXN+1)
	dp[0] = 1
	for i := 1; i <= MAXN; i++ {
		if i-1 >= 0 {
			dp[i] += dp[i-1]
		}
		if i-2 >= 0 {
			dp[i] += dp[i-2]
		}
		if i-3 >= 0 {
			dp[i] += dp[i-3]
		}
	}

	for ; T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		fmt.Fprintln(out, dp[n])
	}
}