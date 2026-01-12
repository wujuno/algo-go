function solution(board, moves) {
    const stack = [];
    let count = 0;

    for (const v of moves) {
        const area = v - 1;

        for (let i = 0; i < board.length; i++) {
            if (board[i][area] !== 0) {
                const doll = board[i][area]; // 현재 집은 인형
                board[i][area] = 0; // 인형을 집었으므로 빈 공간으로 만듦

                // 바구니에 인형이 있고, 맨 위 인형과 현재 인형이 같다면
                if (stack.length > 0 && stack[stack.length - 1] === doll) {
                    stack.pop();
                    count += 2;
                } else {
                    // 일치하지 않을 때만 바구니에 넣음
                    stack.push(doll);
                }

                // 중요: 인형을 하나 집었으므로 이 move(열 탐색)는 종료
                break;
            }
        }
    }

    return count;
}