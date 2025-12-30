function solution(numbers, target) {
  let answer = 0;

  // DFS 함수 정의 (현재 인덱스, 현재까지의 합계)
  function dfs(index, sum) {
    // 1. 탈출 조건: 모든 숫자를 다 사용했을 때
    if (index === numbers.length) {
      // 합계가 타겟 넘버와 같다면 정답 카운트 증가
      if (sum === target) {
        answer++;
      }
      return;
    }

    // 2. 재귀 호출: 현재 숫자를 더하는 경우와 빼는 경우로 분기
    dfs(index + 1, sum + numbers[index]); // 현재 숫자를 더함
    dfs(index + 1, sum - numbers[index]); // 현재 숫자를 뺌
  }

  // DFS 시작 (0번 인덱스부터, 초기 합계 0)
  dfs(0, 0);

  return answer;
}
