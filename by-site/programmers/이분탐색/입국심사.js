function solution(n, times) {
  const bnN = BigInt(n);
  const sortedTimes = times.sort((a, b) => a - b);

  let left = 1n;
  let right = BigInt(sortedTimes[sortedTimes.length - 1]) * bnN;
  let answer = right;

  while (left <= right) {
    let mid = (left + right) / 2n;
    let count = 0n;

    // 1. mid(총 시간)를 각 심사관의 소요 시간(time)으로 나눕니다.
    // 2. 그 결과(해당 심사관이 처리 가능한 인원)를 count에 더합니다.
    for (let time of sortedTimes) {
      count += mid / BigInt(time); // (BigInt 나눗셈은 자동으로 소수점을 버리고 정수만 남깁니다.)
    }

    if (count >= bnN) {
      // n명 이상 처리 가능하므로, 최소 시간을 찾기 위해 범위를 좁힘
      answer = mid; // 일단 mid 값을 answer에 대입.
      right = mid - 1n;
    } else {
      // n명 처리가 불가능하므로, 시간을 늘림
      left = mid + 1n;
    }
  }

  return Number(answer); // 결과는 다시 Number로 변환하여 반환
}

// 문제 Clue
// 1. 압도적인 입력값의 크기 (가장 큰 힌트)
// 일반적으로 알고리즘 문제에서 입력값이 1억을 넘어가면, 시간 복잡도가 $O(n)$인 방법(단순 반복문)으로는 절대 풀 수 없습니다.
// $O(\log n)$의 효율을 가진 알고리즘이 필요한데, 이때 가장 먼저 후보에 오르는 것이 바로 이분 탐색입니다.
// 2. '최솟값'을 구하라는 문제 유형 (매개 변수 탐색)
// 문제 마지막에 **"모든 사람이 심사를 받는 데 걸리는 시간의 '최솟값'을 return 하라"**고 되어 있습니다.
// 이처럼 어떤 값의 최솟값이나 최댓값을 찾는 문제는 종종 '결정 문제'로 바꾸어 풀 수 있습니다.
// 원래 문제 (최적화): "모든 사람을 심사하는 데 걸리는 최소 시간은 몇 분인가?" (어려움)
// 바꾼 문제 (결정): "$T$분이라는 시간을 주면, 그 안에 모든 사람을 심사할 수 있는가? (Yes/No)" (쉬움)
// 3. 정답의 '경계선'이 확실함 (단조성)
// 이분 탐색을 쓰려면 **'정답의 범위가 정렬된 상태'**여야 합니다.
// 이 문제에서의 시간($T$)은 다음과 같은 성질을 갖습니다.28분이 정답이라면, 29분, 30분, 100분에도 당연히 심사를 다 끝낼 수 있습니다. (Yes)
// 28분이 정답이라면, 27분, 26분, 1분에는 절대 심사를 다 끝낼 수 없습니다. (No)
// 즉, 어떤 시점을 기준으로 [No, No, No, Yes, Yes, Yes] 처럼 결과가 명확하게 갈리는 경계선이 존재합니다.
//  우리는 이분 탐색을 통해 그 No가 Yes로 바뀌는 딱 그 시점(최솟값)을 찾아내는 것입니다.
