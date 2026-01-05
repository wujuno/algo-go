function solution(n, edge) {
  const adj = Array.from({ length: n + 1 }, () => []);

  for (const [src, dest] of edge) {
    adj[src].push(dest);
    adj[dest].push(src);
  }

  const distances = new Array(n + 1).fill(0);
  distances[1] = 1;
  const queue = [1];

  let head = 0;
  while (head < queue.length) {
    const cur = queue[head++];

    for (const neighbor of adj[cur]) {
      if (distances[neighbor] === 0) {
        distances[neighbor] = distances[cur] + 1;
        queue.push(neighbor);
      }
    }
  }

  const maxDistance = Math.max(...distances);
  const answer = [];

  distances.forEach((x, i) => {
    if (x == maxDistance) answer.push(i);
  });

  return answer.length;
}

// Array.from({length:}, callback) 으로 리스트 초기화 방법
