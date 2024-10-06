package main

type ladder struct {
	positions map[int]int
}

func NewLadder(n int, positions map[int]int) *ladder {
	mp := make(map[int]int, n)
	for k, v := range positions {
		mp[k] = v

	}
	return &ladder{
		positions: mp,
	}
}

func (l *ladder) IsLadder(pos int) int {
	val, ok := l.positions[pos]
	if !ok {
		return -1
	}
	return val
}
