function solution(prices) {
    const n = prices.length;
    const answer = new Array(n).fill(0);
    const stack = []; // 가격이 떨어지지 않은 시점의 인덱스를 저장

    for (let i = 0; i < n; i++) {
        // 스택이 비어있지 않고, 현재 가격이 스택 최상단(이전 시점) 가격보다 떨어졌을 때
        while (stack.length > 0 && prices[stack[stack.length - 1]] > prices[i]) {
            const index = stack.pop();
            answer[index] = i - index; // 기간 계산
        }
        stack.push(i);
    }

    // 끝까지 가격이 떨어지지 않은 인덱스들 처리
    while (stack.length > 0) {
        const index = stack.pop();
        answer[index] = n - 1 - index;
    }

    return answer;
}