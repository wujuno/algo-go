function solution(numbers) {
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
}
