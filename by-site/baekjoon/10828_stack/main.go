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

	for i := 0; i <n ; i++ {
		var s string
		fmt.Fscan(in, &s)

		switch s {
		case "push":
			var num int
			fmt.Fscan(in, &num)
			stack = append(stack, num)
		case "pop": 
			if len(stack) == 0 {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, stack[len(stack)-1])
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
