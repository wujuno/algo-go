function solution(scoville, K) {
  var answer = 0;
  const heap = []; // 일단 heap을 생성

  function push(val) {
    heap.push(val); // 일단 오른쪽 끝에 추가
    let cur = heap.length - 1; // 추가된 값의 현재 인덱스

    while (cur > 0) {
      let parent = Math.floor((cur - 1) / 2); // 식은 인데스 트리로 그려보면 나옴
      if (heap[parent] <= heap[cur]) break; // 꼭대기가 제일 작은 값이므로 부모가 더 작으면 stop
      else {
        [heap[cur], heap[parent]] = [heap[parent], heap[cur]]; // 부모값과 자식 값 교환
        cur = parent; // index 업데이트
      }
    }
  }

  function pop() {
    if (heap.length == 0) return null; // heap의 길이가 0과 1일때 우선 리턴
    if (heap.length == 1) return heap.pop()

    const min = heap[0] // 리턴할 최소값 먼저 대입
    heap[0] = heap.pop() // 첫번째 인덱스에 마지막 값 넣기

    let cur = 0

    while (cur * 2 + 1 < heap.length) {
      let left = cur * 2 + 1
      let right = cur * 2 + 2

      let smaller = right < heap.length && heap[right] < heap[left] ? right : left

      if (heap[cur] <= heap[smaller]) break
      else {
        [heap[cur], heap[smaller]] = [heap[smaller], heap[cur]];
        cur = smaller
      }
    }

    return min
  }

  for (let s of scoville) push(s)

  while (heap[0] < K) {
    if (heap.length < 2) return -1

    const first = pop();
    const second = pop();
    const mixed = first + second * 2;

    push(mixed);
    answer++;
  }
  return answer;
}
