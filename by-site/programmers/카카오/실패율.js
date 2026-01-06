function solution(N, stages) {
  let totalPlayers = stages.length;

  // 1. 각 스테이지별 머물러 있는(실패한) 사람 수 계산
  const stageCounts = new Array(N + 2).fill(0);
  stages.forEach((s) => stageCounts[s]++);

  // 2. 스테이지별 실패율 계산
  const failRates = [];
  for (let i = 1; i <= N; i++) {
    let failureCount = stageCounts[i];
    let ratio = 0;

    if (totalPlayers > 0) {
      ratio = failureCount / totalPlayers;
      totalPlayers -= failureCount; // 다음 스테이지 도전자 수 계산
    } else {
      ratio = 0; // 해당 스테이지에 도달한 유저가 없는 경우
    }

    failRates.push({ stage: i, ratio: ratio });
  }

  // 3. 정렬 로직
  // ratio 기준 내림차순, 같으면 stage 기준 오름차순
  failRates.sort((a, b) => {
    if (b.ratio === a.ratio) {
      return a.stage - b.stage;
    }
    return b.ratio - a.ratio;
  });

  // 4. 스테이지 번호만 추출
  return failRates.map((item) => item.stage);
}

// 1번 스테이지의 도전자 중 그 스테이지를 못 깬 사람을 빼면, 자연스럽게 2번 스테이지의 도전자가 됩니다.
// 이 방식이 filter를 매번 쓰는 것보다 훨씬 빠릅니다($O(N)$).
// 스테이지가 뒤로 갈수록 유저가 다 떨어져서 도전자가 0명이 되는 구간이 생길 수 있습니다. 이때 NaN 방지를 위해 반드시 체크가 필요합니다.
