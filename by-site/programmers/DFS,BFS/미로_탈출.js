function solution(maps) {
    const rows = maps.length;
    const cols = maps[0].length;

    // 1. 시작점(S), 레버(L), 출구(E)의 좌표 찾기
    let start, lever, exit;
    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (maps[r][c] === 'S') start = [r, c];
            if (maps[r][c] === 'L') lever = [r, c];
            if (maps[r][c] === 'E') exit = [r, c];
        }
    }

    // 2. BFS 함수 정의
    const bfs = (startPos, targetPos) => {
        const [sr, sc] = startPos;
        const [tr, tc] = targetPos;
        const queue = [[sr, sc, 0]]; // [row, col, distance]
        const visited = Array.from({ length: rows }, () => Array(cols).fill(false));
        const dr = [-1, 1, 0, 0];
        const dc = [0, 0, -1, 1];

        visited[sr][sc] = true;

        while (queue.length > 0) {
            const [r, c, dist] = queue.shift();

            if (r === tr && c === tc) return dist;

            for (let i = 0; i < 4; i++) {
                const nr = r + dr[i];
                const nc = c + dc[i];

                // 맵 범위 내에 있고, 벽('X')이 아니며, 방문하지 않은 경우
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                    maps[nr][nc] !== 'X' && !visited[nr][nc]) {
                    visited[nr][nc] = true;
                    queue.push([nr, nc, dist + 1]);
                }
            }
        }
        return -1; // 목표 지점에 도달할 수 없는 경우
    };

    // 3. (S -> L) 거리와 (L -> E) 거리 합산
    const toLever = bfs(start, lever);
    if (toLever === -1) return -1;

    const toExit = bfs(lever, exit);
    if (toExit === -1) return -1;

    return toLever + toExit;
}