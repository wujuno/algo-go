package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N,M int
	fmt.Fscan(in, &N)
	
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	sort.Ints(A)

	
	fmt.Fscan(in, &M)
	for i := 0; i < M; i++ {
		var x int
		fmt.Fscan(in, &x)

		idx := sort.SearchInts(A, x)
		if idx < len(A) && A[idx] == x {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}
