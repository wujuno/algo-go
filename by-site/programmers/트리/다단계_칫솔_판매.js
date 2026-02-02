function solution(enroll, referral, seller, amount) {
    // 1. 빠른 조회를 위해 이름: {부모, 인덱스} 형태의 Map 생성
    const parentMap = new Map();
    const result = new Array(enroll.length).fill(0);

    enroll.forEach((name, i) => {
        parentMap.set(name, { parent: referral[i], index: i });
    });

    // 2. 판매 기록을 하나씩 순회
    for (let i = 0; i < seller.length; i++) {
        let currentName = seller[i];
        let money = amount[i] * 100;

        // 3. 수수료를 줄 수 없을 때까지(0원) 혹은 센터(-)에 도달할 때까지 반복
        while (currentName !== "-" && money > 0) {
            const { parent, index } = parentMap.get(currentName);

            // 내 몫 계산: 전체 금액의 10%를 뺀 나머지
            const tax = Math.floor(money * 0.1);
            const mine = money - tax;

            // 내 수익금 누적
            result[index] += mine;

            // 다음 사람에게 줄 돈과 이름 업데이트
            money = tax;
            currentName = parent;
        }
    }

    return result;
}