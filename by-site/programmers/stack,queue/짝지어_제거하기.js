function solution(s) {
    const stack = [];

    for (let i = 0; i < s.length; i++) {
        const char = s[i];

        // 스택이 비어있지 않고, 맨 위 요소가 현재 문자열과 같다면?
        if (stack.length > 0 && stack[stack.length - 1] === char) {
            stack.pop(); // 짝이 맞으므로 꺼내기만 함 (방금 들어올 예정이었던 char는 넣지도 않음)
        } else {
            stack.push(char); // 짝이 안 맞으면 넣기
        }
    }

    // 최종적으로 스택이 비어있으면 1, 남은게 있으면 0
    return stack.length === 0 ? 1 : 0;
}