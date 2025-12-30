/* function solution(progresses, speeds) {
  var answer = [];

  // 1. 각 작업이 완료되기까지 걸리는 일수 계산
  // 예: [93, 30, 55], [1, 30, 5] -> [7, 3, 9]
  let days = progresses.map((progress, index) =>
    Math.ceil((100 - progress) / speeds[index])
  );

  // 2. 첫 번째 작업의 완료 일수를 기준으로 설정
  let maxDay = days[0];
  let count = 0;

  for (let i = 0; i < days.length; i++) {
    if (days[i] <= maxDay) {
      // 현재 작업이 기준일보다 빨리 끝나면 같이 배포
      count++;
    } else {
      // 기준일보다 오래 걸리는 작업을 만나면 이전까지의 그룹 배포
      answer.push(count);
      count = 1; // 새로운 그룹의 첫 번째 작업
      maxDay = days[i]; // 기준일 갱신
    }
  }

  // 마지막으로 쌓인 그룹 배포
  answer.push(count);

  return answer;
}
 */

function solution(progresses, speeds) {}
