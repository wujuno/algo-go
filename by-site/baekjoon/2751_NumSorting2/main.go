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

	var N int
	fmt.Fscan(in, &N)
	q := make([]int, N)

	for i :=0; i < N; i++ {
		var num int
		fmt.Fscan(in, &num)

		q[i] = num
	}

	sort.Ints(q)

	for i := 0; i < N; i++ {
		fmt.Fprintln(out, q[i])
	}
}