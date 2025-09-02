# BOJ 10828 스택

- 티어: Silver IV (solved.ac)
- 체감 난이도/유형: 기초 구현(스택)
- 아이디어: slice로 스택 구현 후 명령어별 O(1) 처리 (cap을 N으로 선할당)
- 복잡도: 시간 O(N), 공간 O(N)
- 풀이 일자: 2025-09-02

## 회고

- fmt.Fscan(in, &n)은 공백/개행 기준 토큰 단위로 읽으며, 읽은 지점 이후부터 이어서 파싱됨.

- Go에서는 slice를 스택처럼 사용(append, 슬라이싱으로 pop/top).

- make([]int, 0, n)으로 cap 선할당

- 빈 스택 처리가 핵심: pop/top에서 비었으면 -1, empty는 비었으면 1, 아니면 0.

- bufio.NewReader/Writer + defer out.Flush()로 빠르고 안정적인 I/O.

## 키워드

- ds/stack
- pattern/implementation
- pattern/command-parse
- container/slice
- io/bufio
- io/fscan
- edge/empty-pop
- edge/empty-top
- perf/fast-io
- complexity/time-O(N)-per-op
- complexity/space-O(N)-maxstack
- N/num-commands
