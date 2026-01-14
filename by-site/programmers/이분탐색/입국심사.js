function solution(n, times) {
  // 1. 오름차순 정렬 (효율적인 탐색을 위해 권장)
  times.sort((a, b) => a - b);

  let minTime = 1;
  let maxTime = n * times[times.length - 1]; // 가장 오래 걸리는 경우
  let answer = maxTime;

  while (minTime <= maxTime) {
    // 2. 중간값(심사 시간) 설정
    let midTime = Math.floor((minTime + maxTime) / 2);

    // 3. midTime 동안 총 몇 명을 심사할 수 있는지 계산
    // 각 심사관이 midTime 동안 처리할 수 있는 인원: Math.floor(midTime / 심사시간)
    let totalProcessed = times.reduce((acc, time) => {
      return acc + Math.floor(midTime / time);
    }, 0);

    // 4. 이분 탐색 조건 체크
    if (totalProcessed >= n) {
      // n명 이상 처리 가능하므로, 시간을 더 줄여볼 수 있음
      answer = midTime;
      maxTime = midTime - 1;
    } else {
      // n명 처리가 불가능하므로, 시간을 더 늘려야 함
      minTime = midTime + 1;
    }
  }

  return answer;
}