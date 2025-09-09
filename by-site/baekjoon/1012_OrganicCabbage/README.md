# BOJ 1012 유기농 배추

- 티어: 실버 II
- 난이도(체감): 연결요소 + BFS/DFS 기본 (쉬움~보통)
- 아이디어: grid[y][x] = true로 배추 위치를 표시하고, visited[y][x] 배열을 둔다. 모든 칸을 순회하며 아직 방문하지 않은 배추 칸을 만나면 거기서 BFS/DFS로 flood fill → 연결된 모든 칸을 방문 처리 → 카운트 +1.
- 복잡도: 시간: O(N\*M) (최대 50×50, 충분히 여유) 공간: O(N\*M)
- 풀이 일자: 2025-09-03
- 재풀이 일자: 2025-09-09

## 회고

- 큐 시드 삽입 후 tail++ 누락 → head < tail가 처음부터 거짓이 되어 BFS가 돌지 않음.

- grid[i] = make([]bool, 0, M)처럼 len=0으로 만들고 인덱싱 시도 → 패닉.

- pos{ny, nx}처럼 x/y 순서 뒤바뀜.

## 키워드

- graph/grid
- flood fill
- connected components
- BFS/DFS
- visited marking
- head/tail queue
- retry
