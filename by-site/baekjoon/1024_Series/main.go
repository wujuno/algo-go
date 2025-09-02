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

	var N int64 
	var L int
	fmt.Fscan(in, &N, &L)

	found := false
	var start int64
	var ansK int

	for k := L; k <= 100; k++ {
		num := 2*N - int64(k*(k-1))
		den := int64(2*k)

		if num < 0 {
			break
		}

		if num % den != 0 {
			continue
		}

		x := num / den
		if x < 0 {
			continue
		}

		start = x
		ansK = k
		found = true
		break
	}

	if !found {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 0; i < ansK; i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, start+int64(i))
	}
	fmt.Fprintln(out)
}