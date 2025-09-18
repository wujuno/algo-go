package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// N 나무 수, M 나무 길이

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var M int64
	fmt.Fscan(in, &N, &M)

	h := make([]int64, N)
	var maxH int64
	for i:=0; i < N; i++ {
		fmt.Fscan(in, &h[i])
		if h[i] > maxH {
			maxH = h[i]
		}
	}

	var lo int64= 0
	hi := maxH
	var ans int64 = 0

	for lo <= hi {
		mid := (lo + hi) / 2
		var got int64 =0

		for _,x := range h {
			if x > mid {
				got += x - mid
				if got >= M {
					break
				}
			}
		}

		if got >= M {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	fmt.Fprintln(out, ans)
}