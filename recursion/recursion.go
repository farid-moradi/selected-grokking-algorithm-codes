package main

import "fmt"

func main() {
	// countdown(12)
	// countdownSimple(12)
	fmt.Println(fact(20))
	fmt.Println(factSimple(20))
}

func countdown(i int) {
	// fmt.Println(i)
	if i == 0 {
		// fmt.Println("done")
		return
	}
	countdown(i - 1)
}

func countdownSimple(i int) {
	for i >= 0 {
		// fmt.Println(i)
		i--
	}
}

func fact(i int) (r int) {
	if i == 0 || i == 1 {
		return 1
	}
	r = i * fact(i-1)
	return
}

func factSimple(i int) (r int) {
	r = 1
	for i > 0 {
		r *= i
		i--
	}
	return
}
