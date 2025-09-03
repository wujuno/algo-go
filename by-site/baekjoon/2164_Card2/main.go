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

	var N int
	fmt.Fscan(in, &N)
	set := make([]int, 0, N)

	for i := 1; i <= N; i++ {
		set = append(set, i)
	}

	for len(set) > 1 {
		set = set[1:]

		top := set[0]
		set = set[1:]
		set = append(set, top)
	}
	fmt.Fprintln(out, set[0])
}