function solution(participant, completion) {
  const runnerMap = new Map();

  // 1. 참여자 명단을 돌며 Map에 이름과 인원수 기록
  for (const name of participant) {
    runnerMap.set(name, (runnerMap.get(name) || 0) + 1);
  }

  // 2. 완주자 명단을 돌며 Map에서 인원수 차감
  for (const name of completion) {
    runnerMap.set(name, runnerMap.get(name) - 1);
  }

  // 3. 인원수가 남아있는(0보다 큰) 사람이 완주하지 못한 선수
  for (const [name, count] of runnerMap) {
    if (count > 0) return name;
  }
}
