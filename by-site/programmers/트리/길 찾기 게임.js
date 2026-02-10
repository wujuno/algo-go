function solution(nodeinfo) {
    // 1. 노드에 원래 번호(1번부터)를 포함시켜 객체화
    const nodes = nodeinfo.map((info, idx) => ({
        x: info[0],
        y: info[1],
        id: idx + 1,
        left: null,
        right: null
    }));

    // 2. y값 내림차순, x값 오름차순으로 정렬
    nodes.sort((a, b) => {
        if (b.y !== a.y) return b.y - a.y;
        return a.x - b.x;
    });

    // 3. 트리 구성 함수
    const insertNode = (parent, child) => {
        if (child.x < parent.x) {
            if (parent.left === null) parent.left = child;
            else insertNode(parent.left, child);
        } else {
            if (parent.right === null) parent.right = child;
            else insertNode(parent.right, child);
        }
    };

    const root = nodes[0];
    for (let i = 1; i < nodes.length; i++) {
        insertNode(root, nodes[i]);
    }

    // 4. 순회 알고리즘
    const preorderList = [];
    const postorderList = [];

    const preorder = (node) => {
        if (!node) return;
        preorderList.push(node.id);
        preorder(node.left);
        preorder(node.right);
    };

    const postorder = (node) => {
        if (!node) return;
        postorder(node.left);
        postorder(node.right);
        postorderList.push(node.id);
    };

    preorder(root);
    postorder(root);

    return [preorderList, postorderList];
}