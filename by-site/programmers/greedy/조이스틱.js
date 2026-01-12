/* function solution(name) {
  let answer = 0;
  const len = name.length;

  // 1. 상하 이동 횟수 계산
  for (let i = 0; i < len; i++) {
    const charCode = name.charCodeAt(i);
    // 'A'(65)에서 위로 가는 경우 vs 'Z'(90)로 돌아가는 경우 중 최소값
    answer += Math.min(charCode - 65, 91 - charCode);
  }

  // 2. 좌우 이동 횟수 계산 (최적 경로 찾기)
  let minMove = len - 1; // 기본적으로는 끝까지 쭉 오른쪽으로 가는 경우

  for (let i = 0; i < len; i++) {
    let next = i + 1;

    // 연속된 'A'가 끝나는 지점 찾기
    while (next < len && name[next] === "A") {
      // 'A' 뭉텅이가 문자열 뒷부분에 또 있거나, 더 긴 'A' 뭉텅이가 나중에 나타날 수 있기 때문에 끝까지 확인해야 합니다.
      next++;
    }

    // 경로 1: 오른쪽으로 갔다가 다시 왼쪽으로 돌아가는 경우
    // (0 -> i -> 0 -> len-1 쪽의 next)
    minMove = Math.min(minMove, i * 2 + (len - next));

    // 경로 2: 처음부터 왼쪽으로 먼저 갔다가 다시 오른쪽으로 돌아가는 경우
    // (0 -> len-1 쪽의 next -> 0 -> i)
    minMove = Math.min(minMove, (len - next) * 2 + i);
  }

  return answer + minMove;
}

 */

function solution(name) {
  let answer = 0;
  const len = name.length;

  for (let i = 0; i < len; i++) {
    const charCode = name.charCodeAt(i);
    answer += Math.min(charCode - 65, 91 - charCode)
  }

  let minMove = len - 1

  for (let i = 0; i < len; i++) {
    let next = i + 1;

    while (next < len && name[next] === 'A') {
      next++
    }

    minMove = Math.min(minMove, i * 2 + (len - next))

    minMove = Math.min(minMove, (len - next) * 2 + i)
  }

  return answer + minMove
}
