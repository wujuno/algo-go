function solution(clothes) {
  // 1. 의상 종류별 개수를 담을 객체 생성
  const counts = {};

  for (const [name, type] of clothes) {
    // 해당 종류가 처음이면 1, 이미 있으면 +1
    counts[type] = (counts[type] || 0) + 1;
  }

  // 2. 조합 계산 (각 종류별 개수 + 1)을 모두 곱함
  let answer = 1;
  for (const key in counts) {
    answer *= counts[key] + 1;
  }

  // 3. 아무것도 입지 않은 경우 1을 빼고 반환
  return answer - 1;
}
