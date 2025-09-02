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
	fmt.Fscan(in, &n)

	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "push":
			var x int
			fmt.Fscan(in, &x)
			stack = append(stack, x)

		case "pop":
			if len(stack) == 0 {
				fmt.Fprintln(out, -1)
			} else {
				top := stack[len(stack)-1]
				fmt.Fprintln(out, top)
				stack = stack[:len(stack)-1]
			}

		case "size":
			fmt.Fprintln(out, len(stack))

		case "empty":
			if len(stack) == 0 {
				fmt.Fprintln(out, 1)
			} else {
				fmt.Fprintln(out, 0)
			}

		case "top":
			if len(stack) == 0 {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, stack[len(stack)-1])
			}
		}
	}

}