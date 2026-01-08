function solution(dirs) {
  // 1. 방향 이동을 위한 매핑
  const move = {
    U: [0, 1],
    D: [0, -1],
    L: [-1, 0],
    R: [1, 0],
  };

  let curr = [0, 0]; // 현재 위치 [x, y]
  const visited = new Set(); // 방문한 경로를 저장할 Set

  for (let dir of dirs) {
    const [dx, dy] = move[dir];
    const nx = curr[0] + dx;
    const ny = curr[1] + dy;

    // 2. 좌표평면 경계 확인 (정사각형 -5 ~ 5)
    if (nx < -5 || nx > 5 || ny < -5 || ny > 5) continue;

    // 3. 경로 식별자 생성
    // (A -> B)와 (B -> A)는 같은 경로이므로 정렬하여 저장
    const path1 = `${curr[0]}${curr[1]}${nx}${ny}`;
    const path2 = `${nx}${ny}${curr[0]}${curr[1]}`;

    // 경로의 방향성을 없애기 위해 더 작은 좌표를 앞에 두는 방식도 가능하지만,
    // 여기서는 두 가지 경우 중 하나만 Set에 저장하도록 처리합니다.
    const route = [path1, path2].sort().join("");

    visited.add(route);

    // 4. 현재 위치 갱신
    curr = [nx, ny];
  }

  // 5. 중복 제거된 경로의 개수가 곧 방문 길이
  return visited.size;
}
