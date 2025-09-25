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

	var N, K int
	fmt.Fscan(in, &N, &K)
	
	coins := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &coins[i])
	}

	count := 0

	for i := N-1; i >= 0; i-- {
		if coins[i] > K {
			continue
		}

		use := K / coins[i]
		count += use
		K -= use * coins[i]
		if K == 0 {
			break
		}

	}

	fmt.Fprintln(out, count)
}

/* func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K int
	fmt.Fscan(in, &N, &K)
	
	coins := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &coins[i])
	}

	count := 0

	for i := N - 1; i >= 0; i-- {
		if coins[i] > K {
			continue
		}
		use := K / coins[i]
		count += use
		K -= use * coins[i]
		if K == 0 {
			break
		}
	}

	fmt.Fprintln(out, count)
} */