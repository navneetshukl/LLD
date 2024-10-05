package main

type snake struct {
	positions map[int]int
}

func NewSnake(n int) *snake {
	mp := make(map[int]int, n)
	return &snake{
		positions: mp,
	}
}

func (sn *snake) FillSnakePosition(mp map[int]int) {
	for k, v := range mp {
		sn.positions[k] = v

	}
}

func (sn *snake) IsSnake(pos int) int {
	val, ok := sn.positions[pos]
	if !ok {
		return -1
	}
	return val
}

// type LadderSvc interface{
// 	FillladderPosition(mp map[int]int)
// 	IsLadder(pos int) int
// }