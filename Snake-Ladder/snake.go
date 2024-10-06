package main

type snake struct {
	positions map[int]int
}

func NewSnake(n int,positions map[int]int) *snake {
	mp := make(map[int]int, n)
	for k, v := range positions {
		mp[k] = v

	}
	return &snake{
		positions: mp,
	}
}


func (sn *snake) IsSnake(pos int) int {
	val, ok := sn.positions[pos]
	if !ok {
		return -1
	}
	return val
}
