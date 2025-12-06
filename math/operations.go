package math

func Add(x int, y int) int {
	return x + y
}

func AddAll(x ...int) (result int) {
	for _, v := range x {
		result += v
	}
	return
}