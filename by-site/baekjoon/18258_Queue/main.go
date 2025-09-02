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

	q := make([]int, 0, n)
	head := 0 

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "push":
			var x int
			fmt.Fscan(in, &x)
			q = append(q, x)

		case "pop":
			if head >= len(q) {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, q[head]) 
				head++
			}

		case "size":
			fmt.Fprintln(out, len(q)-head)

		case "empty":
			if head >= len(q) {
				fmt.Fprintln(out, 1)
			} else {
				fmt.Fprintln(out, 0)
			}

		case "front":
			if head >= len(q) {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, q[head]) 
			}

		case "back":
			if head >= len(q) {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, q[len(q)-1]) 
			}
		}
	}
}
