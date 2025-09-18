# BOJ 2805 나무 자르기

- 티어: Silver II (solved.ac)
- 난이도(체감): 이분 탐색 / 파라메트릭 서치 기초
- 아이디어: “절단기 높이 H로 잘랐을 때 얻는 합 f(H)”가 단조 감소하므로 f(H) ≥ M을 만족하는 최댓값 H(upper bound) 를 이분 탐색으로 찾는다.
- 복잡도: 시간 O(N log max(h)), 공간 O(N)
- 풀이 일자: 2025-09-18

## 회고

- upper bound 방향(“조건 만족 시 높이를 올린다”)을 혼동하기 쉬움.

- int로 합계를 더하면 오버플로우 위험 → int64 필수.

- hi = max(h)로 시작해야 불필요한 반복이 줄고 코너 케이스(모두 낮은 나무)에도 안전.

- M = 0이면 어떤 높이로 잘라도 되므로 자연스럽게 max(h)가 정답으로 귀결.

## 키워드

- ds/array

- pattern/binary-search|parametric-search
