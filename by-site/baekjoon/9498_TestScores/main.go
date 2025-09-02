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

	var score int
	fmt.Fscan(in, &score)

	switch score/10 {
	case 10, 9:
		fmt.Fprintln(out, "A")
	case 8:
		fmt.Fprintln(out, "B")
	case 7:
		fmt.Fprintln(out, "C")
	case 6:
		fmt.Fprintln(out, "D")
	default:
		fmt.Fprintln(out, "F")
	}
}