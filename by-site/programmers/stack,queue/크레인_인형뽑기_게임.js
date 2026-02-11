function solution(board, moves) {
    const queue = []
    let count = 0
    moves.forEach(v => {
        const idx = v - 1

        for (let i = 0; i < board.length; i++) {
            if (board[i][idx] == 0) continue
            else {
                if (queue.length != 0 && queue[queue.length - 1] == board[i][idx]) {
                    queue.pop()
                    count += 2
                } else {
                    queue.push(board[i][idx])
                }

                board[i][idx] = 0
                break
            }
        }
    })

    return count
}