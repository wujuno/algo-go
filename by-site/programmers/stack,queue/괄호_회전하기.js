function solution(s) {
    let answer = 0;
    const n = s.length;

    // 1. 문자열을 n번 회전시킵니다.
    for (let i = 0; i < n; i++) {
        // i만큼 회전시킨 문자열 생성
        const rotated = s.slice(i) + s.slice(0, i);

        // 2. 올바른 괄호인지 확인
        if (isValid(rotated)) {
            answer++;
        }
    }

    return answer;
}

// 올바른 괄호인지 판별하는 함수 (Stack 활용)
function isValid(str) {
    const stack = [];
    const pairs = {
        ')': '(',
        ']': '[',
        '}': '{'
    };

    for (let char of str) {
        // 여는 괄호인 경우 스택에 추가
        if (char === '(' || char === '[' || char === '{') {
            stack.push(char);
        } else {
            // 닫는 괄호인 경우
            // 스택이 비어있거나 짝이 맞지 않으면 false
            const top = stack.pop();
            if (top !== pairs[char]) {
                return false;
            }
        }
    }

    // 모든 검사가 끝난 후 스택이 비어있어야 올바른 괄호
    return stack.length === 0;
}