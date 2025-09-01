
package dsu

type DSU struct {
	Parent []int
	Rank   []int
}

func New(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := range p { p[i] = i }
	return &DSU{p, r}
}

func (d *DSU) Find(x int) int {
	if d.Parent[x] == x { return x }
	d.Parent[x] = d.Find(d.Parent[x])
	return d.Parent[x]
}

func (d *DSU) Union(a, b int) {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb { return }
	if d.Rank[ra] < d.Rank[rb] {
		d.Parent[ra] = rb
	} else if d.Rank[rb] < d.Rank[ra] {
		d.Parent[rb] = ra
	} else {
		d.Parent[rb] = ra
		d.Rank[ra]++
	}
}
