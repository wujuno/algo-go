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

	var N, L int
	fmt.Fscan(in, &N, &L)

	found := false
	var start int
	var ansK int 
	
	for i := L; i<=100; i++ {
		a := 2*N - (i*(i-1))
		b := 2*i

		if a <0 {
			continue
		}
		if a % b != 0 {
			continue
		}

		start = a/b
		ansK = i
		found = true
		break
	}

	if !found {
		fmt.Fprintln(out, -1)
	}

	for i :=0; i<ansK; i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, start+i)
	}
}

/* func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int64 
	var L int
	fmt.Fscan(in, &N, &L)

	found := false // 최소 값으로 조건문을 탈출하기 위한 flag
	var start int64
	var ansK int

	for k := L; k <= 100; k++ { // 주어진 최소길이 부터 100까지 순회
		num := 2*N - int64(k*(k-1))
		den := int64(2*k)

		if num < 0 { // 분가값이 0보다 작아지는 순간 -1 출력
			break
		}

		if num % den != 0 { // 자연수가 아니라면 지속
			continue
		}

		x := num / den

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
} */
