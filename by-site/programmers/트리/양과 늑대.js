function solution(info, edges) {
    let maxSheep = 0;
    const tree = Array.from({ length: info.length }, () => []);
    edges.forEach(([parent, child]) => tree[parent].push(child));

    function dfs(currentNodes, sheep, wolf, possibleNodes) {
        // 1. 현재 노드들의 정보를 바탕으로 양/늑대 카운트 업데이트는
        // 호출 시점에서 계산해서 넘겨주는 것이 편합니다.
        maxSheep = Math.max(maxSheep, sheep);

        // 2. 다음에 갈 수 있는 후보들을 하나씩 방문
        possibleNodes.forEach((node, index) => {
            const isWolf = info[node];
            const nextSheep = sheep + (isWolf ? 0 : 1);
            const nextWolf = wolf + (isWolf ? 1 : 0);

            // 3. 늑대가 양보다 같거나 많아지면 못 감
            if (nextSheep <= nextWolf) return;

            // 4. 새로운 후보 목록 생성
            // (기존 후보 중 현재 선택한 노드 제외 + 현재 선택한 노드의 자식들 추가)
            const nextPossible = [...possibleNodes];
            nextPossible.splice(index, 1); // 현재 노드 제외
            nextPossible.push(...tree[node]); // 자식 추가

            dfs(node, nextSheep, nextWolf, nextPossible);
        });
    }

    // 루트 노드(0)부터 시작
    // 0번은 항상 양이므로 (0, 1, 0, [0번의 자식들])로 시작
    dfs(0, 1, 0, tree[0]);

    return maxSheep;
}

// 핵심: possibleNodes 배열을 매 단계마다 새로 생성하여 넘겨주는 것이 포인트입니다.

// 주의: 자바스크립트에서 배열을 복사할 때 [...array](Spread Operator)를 사용하면 깊은 복사(1레벨)가 되어 원본 배열을 해치지 않고 탐색할 수 있습니다.