package main

type ladder struct {
	positions map[int]int
}

func NewLadder(n int) *ladder {
	mp := make(map[int]int, n)
	return &ladder{
		positions: mp,
	}
}

func (l *ladder) FillladderPosition(mp map[int]int) {
	for k, v := range mp {
		l.positions[k] = v

	}
}

func (l *ladder) IsLadder(pos int) int {
	val, ok := l.positions[pos]
	if !ok {
		return -1
	}
	return val
}
