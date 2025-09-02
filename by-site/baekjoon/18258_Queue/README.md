# BOJ 18258 큐2

- 티어: Silver IV (solved.ac)
- 난이도(체감): 기초 자료구조(큐), I/O 최적화
- 아이디어: 슬라이스 하나로 큐를 구현하되, head 인덱스를 따로 두어 pop 시 앞에서 꺼낸 것처럼 처리
- 복잡도: 시간: O(N), 공간: O(N)
- 풀이 일자: 2025-09-02

## 회고

- 처음에 스택처럼 뒤에서 꺼내고 front/back을 뒤집어 처리해서 오답이 났음.

- queue = queue[1:]로 매번 자르면 최악 O(N)이 누적되므로, head 인덱스 방식이 안전하고 빠름.

## 키워드

- ds/queue
- slice/head-index
- io/bufio
- fmt/fscan
