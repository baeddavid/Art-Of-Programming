package ch1

func E(x int, y int) int {
	if x > y {
		x, y = y, x
	}

	remainder := x % y
	x /= y
	for {
		if remainder == 0 {
			return y
		}
		x = y
		y = remainder
		remainder = x % y
		x /= y
	}
}
