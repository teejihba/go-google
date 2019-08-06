package disjointset

type DisjointSet struct {
	Size   int
	Parent []int
	Rank   []int
}

func (set *DisjointSet) MakeSet(n int) {
	set.Size = n
	set.Parent = make([]int, n)
	set.Rank = make([]int, n)
	for i := range set.Parent {
		set.Parent[i] = i
		set.Rank[i] = 0
	}
}

// Find by Path Compression
func (set *DisjointSet) Find(x int) int {
	if set.Parent[x] != x {
		set.Parent[x] = set.Find(set.Parent[x])
	}
	return set.Parent[x]
}

//Union by Rank
func (set *DisjointSet) Union(x int, y int) {
	xSet := set.Find(x)
	ySet := set.Find(y)

	if set.Rank[xSet] < set.Rank[ySet] {
		set.Parent[xSet] = ySet
	} else if set.Rank[xSet] > set.Rank[ySet] {
		set.Parent[ySet] = xSet
	} else {
		set.Parent[ySet] = xSet
		set.Rank[xSet]++
	}
}
