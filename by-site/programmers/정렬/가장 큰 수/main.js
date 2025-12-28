function solution(numbers) {
  // 1. 숫자 배열을 문자열 배열로 변환합니다.
  const answer = numbers
    .map(String)
    // 2. 두 문자열을 합쳤을 때 더 큰 숫자가 되는 순서로 정렬합니다.
    // (b + a)와 (a + b)를 비교하여 내림차순 정렬
    .sort((a, b) => b + a - (a + b))
    // 3. 정렬된 배열을 하나로 합칩니다.
    .join("");

  // 4. 모든 숫자가 0일 경우(예: [0, 0, 0]) "000"이 아닌 "0"을 반환하도록 예외 처리합니다.
  return answer[0] === "0" ? "0" : answer;
}
