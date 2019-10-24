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
