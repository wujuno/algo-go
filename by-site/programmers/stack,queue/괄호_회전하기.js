function solution(s) {
    const obj = {
        "}": "{",
        "]": "[",
        ")": "("
    }
    let count = 0
    const sArr = s.split('')
    for (let i = 0; i < s.length; i++) {
        // 최적화를 위해 
        let start = i

        sArr.push(s[start])
        const mArr = sArr.slice(start + 1)

        const queue = []
        // sArr을 queue에 하나씩 넣어 남는값이 있는 지확인
        mArr.forEach(v => {
            if (queue.length > 0 && obj[v] == queue[queue.length - 1]) {
                queue.pop()
            } else {
                queue.push(v)
            }
        })
        if (queue.length == 0) count++
        // 남는 값 없으면 count++
    }

    return count
}