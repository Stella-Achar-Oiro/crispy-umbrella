package main

func Abort(a, b,c,d,e int) int {
	n := []int{a, b, c, d, e}
	swap := func(i, j int) {
		n[i], n[j] = n[j], n[i]
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 5; j++ {
			if n[j] < n[i] {
				swap(i, j)
			}
		}
	}
	return n[2]
}