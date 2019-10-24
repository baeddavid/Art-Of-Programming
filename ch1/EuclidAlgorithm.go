package ch1

func E(m int, n int) int {
	// E0
	if n > m {
		m, n = n, m
	}

	for {
		if n == 0 {
			return m
		}
		// E1 - E3
		t := n
		n = m % n
		m = t
	}
}
