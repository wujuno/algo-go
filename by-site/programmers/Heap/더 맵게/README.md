부모/자식 인덱스 공식

parent = (cur - 1) / 2

left = cur _ 2 + 1, right = cur _ 2 + 2

push는 위로 (cur > 0)

heap.push()로 맨 뒤에 넣고, 부모(parent)가 나보다 크면 스왑하며 올라갑니다.

pop은 아래로 (cur \* 2 + 1 < heap.length)

맨 앞(heap[0])을 빼고, 맨 뒤 요소를 앞으로 가져옵니다.

두 자식 중 **더 작은 자식(smaller)**과 나를 비교해서, 내가 더 크면 내려갑니다.
