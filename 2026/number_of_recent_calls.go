package _026

type RecentCounter struct {
	Elements []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.Elements = append(r.Elements, t)

	for r.Elements[0] < t-3000 {
		r.Elements = r.Elements[1:]
	}

	return len(r.Elements)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
