
package graph

type P struct{ R, C int }

// GridBFS skeleton using 4-direction movement.
// caller provides: rows, cols, start list, visit predicate, enqueue callback.
func GridBFS(rows, cols int, starts []P, canVisit func(r, c int) bool, visit func(r, c int, dist int)) {
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}

	q := make([]P, 0, rows*cols)
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	push := func(p P) { q = append(q, p) }
	pop := func() P { v := q[0]; q = q[1:]; return v }

	for _, s := range starts {
		if canVisit(s.R, s.C) {
			dist[s.R][s.C] = 0
			push(s)
			visit(s.R, s.C, 0)
		}
	}

	for len(q) > 0 {
		cur := pop()
		for k := 0; k < 4; k++ {
			r := cur.R + dr[k]
			c := cur.C + dc[k]
			if r < 0 || r >= rows || c < 0 || c >= cols { continue }
			if !canVisit(r, c) || dist[r][c] != -1 { continue }
			dist[r][c] = dist[cur.R][cur.C] + 1
			visit(r, c, dist[r][c])
			push(P{r, c})
		}
	}
}
