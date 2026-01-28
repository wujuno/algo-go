function solution(N, stages) {
  let result = [];
  let totalPlayers = stages.length; // 처음엔 모든 인원이 1스테이지 도달자

  for (let i = 1; i <= N; i++) {
    // 1. 현재 스테이지에 도달했으나 클리어하지 못한 인원
    const challengers = stages.filter(v => v === i).length;

    // 2. 실패율 계산 (분모가 0인 경우 예외처리 필수)
    const failRatio = totalPlayers === 0 ? 0 : challengers / totalPlayers;

    result.push({ stage: i, ratio: failRatio });

    // 3. 다음 스테이지 도달 인원은 현재 스테이지 인원을 뺀 만큼임
    totalPlayers -= challengers;
  }

  // 4. 정렬: 실패율 내림차순 -> 실패율 같으면 스테이지 오름차순
  result.sort((a, b) => {
    if (b.ratio === a.ratio) return a.stage - b.stage;
    return b.ratio - a.ratio;
  });

  return result.map(v => v.stage);
}

// 1번 스테이지의 도전자 중 그 스테이지를 못 깬 사람을 빼면, 자연스럽게 2번 스테이지의 도전자가 됩니다.
// 이 방식이 filter를 매번 쓰는 것보다 훨씬 빠릅니다($O(N)$).
// 스테이지가 뒤로 갈수록 유저가 다 떨어져서 도전자가 0명이 되는 구간이 생길 수 있습니다. 이때 NaN 방지를 위해 반드시 체크가 필요합니다.
