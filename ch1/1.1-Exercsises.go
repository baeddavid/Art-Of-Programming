package ch1

import "fmt"

// 1.1
// Given (a, b, c, d) make (b, c, d, a) using the minimum number of swaps
// Swapped in three moves
func SwapABCD() {
	counter := 0
	a := []byte{'a', 'b', 'c', 'd'}
	a[0], a[1] = a[1], a[0]
	counter++
	fmt.Println("Counter", counter)
	fmt.Println(a)
	a[1], a[2] = a[2], a[1]
	counter++
	fmt.Println("Counter", counter)
	fmt.Println(a)
	a[2], a[3] = a[3], a[2]
	counter++
	fmt.Println("Counter", counter)
	fmt.Println(a)
}

// 1.2
// Prove that m is always greater than n in E1

// If at any point n < m, than the algorithm breaks because we will always be returned 0. At E0 we fix this by swapping m and n if m > n. Therefore m is always greater that n at E1.

// 1.3
// Change algorithm E such that all triviial assignments m <- n are avoided. Call this function F.
// We avoid all trivial assignments by converting the iterative algorithm to a recursive algorithm. The only neccessary swap is at E0 where we check if x > y and swap them if it is true.
func F(x int, y int) int {
	if x > y {
		x, y = y, x
	}

	if x == 0 {
		return y
	}

	return E(y%x, x)
}

// 1.4
// What is the greatest common divisor of 2166 and 6099

// Plugging in 2166 and 6099 into our algorithm we get 57
