package ch1

func E(x int, y int) int {
	if x > y {
		x, y = y, x
	}

	if x == 0 {
		return y
	}

	return E(y%x, x)
}
