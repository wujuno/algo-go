# BOJ 1927 최소 힙

- 티어: 실버 II
- 난이도(체감): 자료구조(우선순위 큐 / 힙)
- 아이디어: 최소 힙을 사용하면 삽입/삭제가 모두 O(log N).Go에서는 표준 라이브러리 **container/heap**을 쓰는 것이 가장 간단/안전. 사용자 정의 타입에 Len/Less/Swap/Push/Pop 5~6개 메서드를 구현하면 heap.Init/Push/Pop 사용 가능.
- 복잡도: O(log N)
- 풀이 일자: 2025-09-10

## 회고

- 최소 힙 문제는 패턴화하기 좋다. x==0 → pop/print or 0, x>0 → push

## 키워드

- heap
- priority-queue
- container/heap
- min-heap
- O(log N)
