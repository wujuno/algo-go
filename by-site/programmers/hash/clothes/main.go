package clothes

func solution(clothes [][]string) int {
	// 1. 의상 종류별로 개수를 담을 맵 생성
	// key: 의상 종류(string), value: 개수(int)
	counts := make(map[string]int)

	for _, c := range clothes {
		// c[0]은 의상 이름, c[1]은 의상 종류
		clotheType := c[1]
		counts[clotheType]++
	}

	// 2. 조합 계산 (각 종류별 개수 + 1)을 모두 곱함
	answer := 1
	for _, count := range counts {
		// 해당 종류를 입지 않는 경우를 포함하기 위해 +1
		answer *= (count + 1)
	}

	// 3. 아무것도 입지 않은 경우 1을 빼고 반환
	return answer - 1
}