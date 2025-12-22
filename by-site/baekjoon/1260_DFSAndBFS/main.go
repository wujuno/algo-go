package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/* func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, V int
	fmt.Fscan(in, &N, &M, &V)

	g := make([][]int, N+1)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 1; i <= N; i++ {
		sort.Ints(g[i])
	}

	visited := make([]bool, N+1)
	dfsOrder := make([]int, 0, N)

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		dfsOrder = append(dfsOrder, u)
		for _, v := range g[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(V)

	for i, x := range dfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)

	for i := range visited {
		visited[i] = false
	}
	bfsOrder := make([]int, 0, N)
	q := make([]int, 0, N)

	visited[V] = true
	q = append(q, V)

	for head := 0; head < len(q); head++ {
		u := q[head]
		bfsOrder = append(bfsOrder, u)
		for _, v := range g[u] {
			if !visited[v] {
				visited[v] = true
				q = append(q, v)
			}
		}
	}

	// 6) BFS 출력
	for i, x := range bfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)
} */

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M, V int
	fmt.Fscan(in, &N, &M, &V) // ok. 일단 변수 생성 및 입력을 받아야하니까.

	g := make([][]int, N+1) //  ok. N,M을 넣을 중첩 슬라이스 생성
	for i := 0; i < M ; i++ { // ok. 간선(M)의 개수 만큼 입력 줄이 있으니까.
		var a, b int
		fmt.Fscan(in, &a, &b)

		g[a] = append(g[a], b) 
		g[b] = append(g[b], a) // ok. 문제에서 "간선은 양방향이다" 라고 명시.
	}

	for i := 1; i <= N; i++ {
		sort.Ints(g[i])  // ok. "방문할 수 있는 정점이 여러 개인 경우에는 정점 번호가 작은 것을 먼저 방문하고" 문제 명시
	}

	visited := make([]bool, N+1) // ok. 지나간걸 표시해야하니까. n번 점을 가봤는가라는 상태만 저장
	dfsOrder := make([]int, 0, N) // dfs 수행 순서

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true // 입력 받은 숫자는 가본 곳으로 체크
		dfsOrder = append(dfsOrder, u) // 입력 받은 숫자를 dfs 순서에 삽입
		for _, v := range g[u] { // 입력 받은 숫자에서 연결된 다른 값들 확인
			if !visited[v] { // 가보지 않은 데가 있으면 다시 dfs 재귀실행
				dfs(v)
			}
		}
	}
	dfs(V)

	for i, x := range dfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)

	for i := range visited { 
		visited[i] = false // visited 함수 재활용. 전부 fasle로 변경
	}

	bfsOrder := make([]int, 0, N) 
	q := make([]int, 0, N)

	visited[V] = true
	q = append(q, V) // 큐에 첫번재 값 부터 넣기

	for head := 0; head < len(q); head++ { // head값이 len(q)와 같아질때 까지 큐를 순회
		u := q[head]
		bfsOrder = append(bfsOrder, u)
		for _, v := range g[u] {
			if !visited[v] {
				visited[v] = true
				q = append(q, v)
			}
		}
	}

	// 6) BFS 출력
	for i, x := range bfsOrder {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, x)
	}
	fmt.Fprintln(out)

}
