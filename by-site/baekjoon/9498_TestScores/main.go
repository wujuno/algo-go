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

	if (score >= 90 && score <= 100 ) {
		fmt.Fprintln(out, "A")
	} else if (score >= 80 && score < 90 ) {
		fmt.Fprintln(out, "B")
	} else if (score >= 70 && score < 80 ) {
		fmt.Fprintln(out, "C")
	} else if (score >= 60 && score < 70 ) {
		fmt.Fprintln(out, "D")
	} else {
		fmt.Fprintln(out, "F")
	}
}