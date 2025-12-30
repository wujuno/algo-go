/* function solution(scoville, K) {
  let answer = 0;
  const heap = [];

  // 1. 힙 삽입 함수 (가장 끝에 넣고 부모와 비교하며 위로)
  function push(val) {
    heap.push(val);
    let cur = heap.length - 1;
    while (cur > 0) {
      let parent = Math.floor((cur - 1) / 2);
      if (heap[parent] <= heap[cur]) break; // 부모가 더 작으면 중단
      [heap[cur], heap[parent]] = [heap[parent], heap[cur]];
      cur = parent;
    }
  }

  // 2. 힙 삭제 함수 (루트를 빼고, 마지막 요소를 루트로 옮긴 뒤 자식과 비교하며 아래로)
  function pop() {
    if (heap.length === 0) return null;
    if (heap.length === 1) return heap.pop();

    const min = heap[0];
    heap[0] = heap.pop(); // 마지막 요소를 루트로 이동
    let cur = 0;

    while (cur * 2 + 1 < heap.length) {
      // 왼쪽 자식이 있는 동안 반복
      let left = cur * 2 + 1;
      let right = cur * 2 + 2;
      // 왼쪽과 오른쪽 중 더 작은 자식 찾기
      let smaller =
        right < heap.length && heap[right] < heap[left] ? right : left;

      if (heap[cur] <= heap[smaller]) break; // 내가 자식보다 작으면 중단
      [heap[cur], heap[smaller]] = [heap[smaller], heap[cur]];
      cur = smaller;
    }
    return min;
  }

  // 초기 스코빌 배열을 힙에 모두 넣기
  for (let s of scoville) push(s);

  // 3. 메인 로직
  while (heap[0] < K) {
    // 음식이 하나 남았는데도 K보다 작으면 불가능
    if (heap.length < 2) return -1;

    const first = pop();
    const second = pop();
    const mixed = first + second * 2;

    push(mixed);
    answer++;
  }

  return answer;
} */

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
    if(heap.length == 0) return null; // heap의 길이가 0과 1일때 우선 리턴
    if(heap.length == 1) return heap.pop()

    const min = heap[0] // 리턴할 최소값 먼저 대입
    heap[0] = heap.pop() // 첫번째 인덱스에 마지막 값 넣기

    let cur = 0 

    while(cur *2 + 1 < heap.length) {
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
  
  while(heap[0] < K) {
    if(heap.length < 2) return -1

     const first = pop();
    const second = pop();
    const mixed = first + second * 2;

    push(mixed);
    answer++;
  }
  return answer;
}
