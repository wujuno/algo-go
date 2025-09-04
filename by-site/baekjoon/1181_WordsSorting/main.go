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
	arr := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) != len(arr[j]) {
			return len(arr[i]) < len(arr[j])
		}
		return arr[i] < arr[j]
	})

	prev := ""
	for i, s := range arr {
		if i == 0 || s != prev {
			fmt.Fprintln(out, s)
			prev = s
		}
	}
}