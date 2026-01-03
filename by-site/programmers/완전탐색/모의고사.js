function solution(answers) {
  var answer = [];

  // 1. 수포자들의 반복 패턴 정의
  const humanA = [1, 2, 3, 4, 5];
  const humanB = [2, 1, 2, 3, 2, 4, 2, 5];
  const humanC = [3, 3, 1, 1, 2, 2, 4, 4, 5, 5];

  // 2. 점수 계산 (배열을 사용하면 관리하기 편합니다)
  let counts = [0, 0, 0];

  for (let i = 0; i < answers.length; i++) {
    if (answers[i] === humanA[i % humanA.length]) counts[0]++;
    if (answers[i] === humanB[i % humanB.length]) counts[1]++;
    if (answers[i] === humanC[i % humanC.length]) counts[2]++;
  }

  // 3. 가장 높은 점수 찾기
  const maxScore = Math.max(...counts);

  // 4. 최댓값을 가진 수포자 번호를 answer에 추가 (1번부터 시작)
  for (let j = 0; j < counts.length; j++) {
    if (counts[j] === maxScore) {
      answer.push(j + 1);
    }
  }

  return answer;
}

// 가장 높은 점수를 찾은 뒤 높은 점수와 같은 인덱스를 answer에 push하는 방식이 효율이 좋음.
