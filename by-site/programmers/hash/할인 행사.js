function solution(want, number, discount) {
    let answer = 0;

    // 1. 내가 원하는 제품과 수량을 객체 형태로 정리 (Hash Map 역할)
    const wantMap = {};
    want.forEach((item, idx) => {
        wantMap[item] = number[idx];
    });

    // 2. discount 배열을 순회하며 10일간의 윈도우를 확인
    // i는 회원등록을 시작하는 날짜 (인덱스)
    for (let i = 0; i <= discount.length - 10; i++) {
        // 10일간의 할인 품목을 담을 객체
        const currentDiscountMap = {};

        // 3. 시작일부터 10일간의 할인 품목 개수를 세기
        for (let j = i; j < i + 10; j++) {
            const item = discount[j];
            currentDiscountMap[item] = (currentDiscountMap[item] || 0) + 1;
        }

        // 4. 원하는 품목과 수량이 모두 일치하는지 확인
        let isMatch = true;
        for (const item of want) {
            if (currentDiscountMap[item] !== wantMap[item]) {
                isMatch = false;
                break;
            }
        }

        // 5. 모두 일치한다면 정답 카운트 증가
        if (isMatch) answer++;
    }

    return answer;
}