# 🏁 Algo Go Starter

코딩테스트(백준/프로그래머스/LeetCode)용 **Go 스타터 레포**입니다.  
`by-site/baekjoon/<문제번호_식별자>` 폴더에 문제별 `main.go`를 두고 풀이/오답을 기록하세요.

## 1) 문제 선택 규칙(백준 + solved.ac 기준)

- **난이도**: 실버5 → 실버1 → 골드5~4 → 골드3 순서
- **유형**: 배열/문자열/해시 → 스택/큐/덱 → 투포인터/슬윈 → 그래프(BFS/DFS) → 우선순위큐/그리디 → 이진탐색 → DP
- **세트 구성**: 쉬움1 + 보통2(속도), 보통2 + 어려움1(심화) 중 택1

## 2) 풀이 기록 방법

각 문제 폴더에 `README.md`(간단 회고)를 남깁니다.

- 실패 원인 1줄
- 핵심 인사이트 1줄
- 복잡도(O-표기), 재풀이 날짜

## 3) 실행 & 제출 팁

```bash
# 로컬 실행(예시: A+B )
go run ./by-site/baekjoon/1000_AplusB/main.go < by-site/baekjoon/1000_AplusB/input_example.txt
```

- 실제 제출 시에는 `main.go`만 복사하여 백준 제출 창에 붙여넣으면 됩니다.
- 입출력은 `bufio.NewReader/Writer`를 **항상** 사용하세요(시간 초과 방지).

## 4) 폴더 구조

```
by-site/
  baekjoon/
    1000_AplusB/
      main.go
      README.md
      input_example.txt
templates/
  fastio/        # 빠른 I/O 템플릿
  graph/         # BFS, Dijkstra 등
  dsu/           # Union-Find
notes/
  odab_note_template.md
.gitignore
go.mod
README.md
```

## 5) 커밋 규칙(예시)

- `feat(boj): 1000 A+B 풀이 추가`
- `refactor(boj): 2178 미로탐색 큐 → 덱 최적화`
- `docs(notes): 2주차 회고 업데이트`
