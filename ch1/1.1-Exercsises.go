package ch1

import "fmt"

// 1.1
// Given (a, b, c, d) make (b, c, d, a) using the minimum number of swaps
// Swapped in five assignments. Sort of similar to insertion sort bubbling.
func SwapABCD() {
	counter := 0
	a := []byte{'a', 'b', 'c', 'd'}

	temp := a[0]
	counter++
	fmt.Println(a)
	fmt.Println("Counter", counter)
	a[0] = a[1]
	counter++
	fmt.Println(a)
	fmt.Println("Counter", counter)
	a[1] = a[2]
	counter++
	fmt.Println(a)
	fmt.Println("Counter", counter)
	a[2] = a[3]
	counter++
	fmt.Println(a)
	fmt.Println("Counter", counter)
	a[3] = temp
	counter++
	fmt.Println(a)
	fmt.Println("Counter", counter)
}

// 1.2
// Prove that m is always greater than n in E1

// m and n are the previous values of n and r. n > r meaning that m is always greater than n.

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

// 1.5
// Show that the "Procedure for Reading this Set of Books" fails to be a genuine algorithm

// The preface fails in 3 places specifically.
// 1. It fails to be precisely defined. The preface is very ambiguous about what certain things mean. For example N is not clearly defined. Is it an integer? A float? It is never explicitly stated
// 2. It does not have any output. An algorithm can take 0 inputs, but it must always output something.
// 3. It can be argued that the preface fails the efficiency qualifier. Not to say that Knuth work is poor studying material, but the preface's guidelines technically have duplicated work, where the
// reader can potentially read the same part multiple times if going strictly by the chart.

// 1.6
// What is T5, the average number of times step E1 is performed when n = 5?

// Modifying E so that it prints a counter to the console each time E1 is called. We see that plugging in the values of m 1 - 5 we get a value of 2.6.
