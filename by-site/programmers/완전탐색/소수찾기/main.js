/* function solution(numbers) {
  const numberSet = new Set(); // Set을 만들어 중복 제거

  // 1. 모든 숫자 조합을 만드는 재귀 함수 (순열)
  const getPermutations = (currentStr, remainingStr) => {
    if (currentStr.length > 0) {
      numberSet.add(Number(currentStr));
    }

    for (let i = 0; i < remainingStr.length; i++) {
      getPermutations(
        currentStr + remainingStr[i],
        remainingStr.slice(0, i) + remainingStr.slice(i + 1)
      );
    }
  };

  getPermutations("", numbers);

  // 2. 소수 판별 함수
  const isPrime = (num) => {
    if (num < 2) return false;
    for (let i = 2; i <= Math.sqrt(num); i++) {
      if (num % i === 0) return false;
    }
    return true;
  };

  // 3. Set에 저장된 숫자들 중 소수의 개수를 카운트
  let count = 0;
  numberSet.forEach((num) => {
    if (isPrime(num)) count++;
  });

  return count;
} */

function solution(numbers) {
  const set = new Set();
  function getPermutations(currStr, remainingStr) {
    if (currStr.length >= 1) set.add(Number(currStr));

    for (let i = 0; i < remainingStr.length; i++) {
      getPermutations(
        currStr + remainingStr[i],
        remainingStr.slice(0, i) + remainingStr.slice(i + 1)
      );
    }
  }

  function isPrime(x) {
    if (x < 2) return false;
    for (let i = 2; i <= Math.sqrt(x); i++) {
      if (x % i === 0) return false;
    }
    return true;
  }

  getPermutations("", numbers);

  let count = 0;
  set.forEach((v) => {
    if (isPrime(v)) count++;
  });

  return count;
}

// 완전 탐색(Brute Force)임을 인지하기: "모든 조합을 다 해봐야 한다"는 생각이 들면 일단 재귀나 백트래킹을 후보에 두세요.
// 인형 뽑기 모델링: "하나를 뽑고(선택), 남은 것들 중에서 또 뽑는다(재귀)"는 메커니즘을 기억하세요.
// 소수 판별 함수 작성 시 0,1 예외처리를 위해 x < 2 체크
