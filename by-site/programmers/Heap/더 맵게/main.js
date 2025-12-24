/* class MinHeap {
  constructor() {
    this.heap = [];
  }

  push(val) {
    this.heap.push(val);
    this.bubbleUp();
  }

  pop() {
    if (this.size() === 0) return null;
    if (this.size() === 1) return this.heap.pop();

    const min = this.heap[0];
    this.heap[0] = this.heap.pop();
    this.bubbleDown();
    return min;
  }

  peek() {
    return this.heap[0];
  }

  size() {
    return this.heap.length;
  }

  bubbleUp() {
    let index = this.heap.length - 1;
    while (index > 0) {
      let parentIndex = Math.floor((index - 1) / 2);
      if (this.heap[parentIndex] <= this.heap[index]) break;
      [this.heap[parentIndex], this.heap[index]] = [
        this.heap[index],
        this.heap[parentIndex],
      ];
      index = parentIndex;
    }
  }

  bubbleDown() {
    let index = 0;
    const length = this.heap.length;
    while (true) {
      let leftChildIndex = 2 * index + 1;
      let rightChildIndex = 2 * index + 2;
      let swapIndex = null;

      if (leftChildIndex < length) {
        if (this.heap[leftChildIndex] < this.heap[index]) {
          swapIndex = leftChildIndex;
        }
      }

      if (rightChildIndex < length) {
        if (
          (swapIndex === null &&
            this.heap[rightChildIndex] < this.heap[index]) ||
          (swapIndex !== null &&
            this.heap[rightChildIndex] < this.heap[leftChildIndex])
        ) {
          swapIndex = rightChildIndex;
        }
      }

      if (swapIndex === null) break;
      [this.heap[index], this.heap[swapIndex]] = [
        this.heap[swapIndex],
        this.heap[index],
      ];
      index = swapIndex;
    }
  }
}

function solution(scoville, K) {
  let answer = 0;
  const heap = new MinHeap();

  // 1. 모든 요소를 힙에 삽입
  for (const s of scoville) {
    heap.push(s);
  }

  // 2. 가장 작은 값이 K보다 작을 동안 반복
  while (heap.peek() < K) {
    // 음식이 하나 남았는데도 K 미만이면 불가능한 경우
    if (heap.size() < 2) {
      return -1;
    }

    const first = heap.pop();
    const second = heap.pop();
    const mixed = first + second * 2;

    heap.push(mixed);
    answer++;
  }

  return answer;
}
 */

function solution(scoville, K) {
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
}
